package api

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
)

func (s *Server) fileUpload(w http.ResponseWriter, r *http.Request) {
	err := r.ParseMultipartForm(maxUploadSize << 10)
	if err != nil {
		fmt.Fprintf(w, err.Error())
		return
	}
	file, handler, err := r.FormFile("file")
	if err != nil {
		fmt.Fprintf(w, err.Error())
		return
	}
	fileSize := handler.Size
	if fileSize > maxUploadSize {
		fmt.Fprintf(w, "Max Size Reached!")
		return
	}
	fileBytes, err := ioutil.ReadAll(file)
	if err != nil {
		fmt.Fprintf(w, err.Error())
		return
	}
	fileName := handler.Filename
	newPath := filepath.Join(uploadPath, fileName)
	newFile, err := os.Create(newPath)
	if err != nil {
		fmt.Fprintf(w, err.Error())
		return
	}
	defer newFile.Close()
	if _, err = newFile.Write(fileBytes); err != nil {
		fmt.Fprintf(w, err.Error())
		return
	}
	fmt.Fprintf(w, "Success, File: "+fileName+" Size is: "+strconv.FormatInt(fileSize, 10))
	return
}
