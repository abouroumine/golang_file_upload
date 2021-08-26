package main

import (
	s "file_upload/main/api"
)

func main() {
	server := s.Server{}
	server.Initialization()
	port := ":8080"
	server.Run(&port)
}
