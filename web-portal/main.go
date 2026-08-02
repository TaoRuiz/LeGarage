package main

import (
	"log"

	"github.com/TaoRuiz/LeGarage/web-portal/internal/config"
	"github.com/TaoRuiz/LeGarage/web-portal/internal/server"
)

func main() {
	cfg := config.Load()
	s := server.New(cfg)

	log.Println("Starting portal on :8080")
	if err := s.Start(":8080"); err != nil {
		log.Fatal(err)
	}
}
