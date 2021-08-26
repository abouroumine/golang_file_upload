package api

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
)

func (s *Server) filesUpload(w http.ResponseWriter, r *http.Request) {
	_, _, err := r.FormFile("files")
	formData := r.MultipartForm

	if err != nil {
		fmt.Fprintf(w, err.Error())
		return
	}

	filesURLs := make([]string, 0)

	files := formData.File["files"]

	for i, fHeader := range files {
		fileName := fHeader.Filename
		file, _ := files[i].Open()
		fileSize := fHeader.Size
		if fileSize > maxUploadSize {
			fmt.Fprintf(w, "Max Size Reached")
			return
		}
		fileBytes, err := ioutil.ReadAll(file)
		if err != nil {
			fmt.Fprintf(w, err.Error())
			return
		}
		newPath := filepath.Join(uploadPath, fileName)
		newFile, err := os.Create(newPath)
		if err != nil {
			fmt.Fprintf(w, err.Error())
			return
		}
		defer newFile.Close()
		if _, err := newFile.Write(fileBytes); err != nil {
			fmt.Fprintf(w, err.Error())
			return
		}
		filesURLs = append(filesURLs, fileName)
	}
	fmt.Fprintf(w, "Success, The Number of Files Uploaded is: "+strconv.FormatInt(int64(len(filesURLs)), 10))
	return
}
