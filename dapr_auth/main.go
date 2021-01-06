package main

import (
	"authentication/router"
	"log"
	"net/http"
)

func main() {
	s := router.Server
	if err := s.Start(); err != nil && err != http.ErrServerClosed {
		log.Fatalf("error listenning: %v", err)
	}
}

