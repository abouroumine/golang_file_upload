package api

import (
	"log"
	"net/http"
)

const (
	maxUploadSize = 2 * 1024 * 1024 // 2 Mb
	uploadPath    = "files"
)

// Here we define The Server Struct that will contain our ServeMux.
// Responsible for all our http routes available.

// Server ...
type Server struct {
	S *http.ServeMux
}

// This Method is Associated to the Server Struct
// It is used to Start The Server on port defined by 'add'

func (s *Server) Run(add *string) {
	log.Fatal(http.ListenAndServe(*add, s.S))
}

// We Initialize our ServeMux and HandleFunc in this Initialization Function.

func (s *Server) Initialization() {
	s.S = http.NewServeMux()

	fs := http.FileServer(http.Dir(uploadPath))
	s.S.Handle("/", http.StripPrefix("/files/", fs))

	s.S.HandleFunc("/savefile", s.fileUpload)
	s.S.HandleFunc("/savefiles", s.filesUpload)
}
