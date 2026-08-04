package main

import (
	"log"

	"github.com/TaoRuiz/LeGarage/web-portal/internal/config"
	"github.com/TaoRuiz/LeGarage/web-portal/internal/server"
	"github.com/TaoRuiz/LeGarage/web-portal/internal/store"
)

func main() {
	cfg := config.Load()

	db, err := store.New("data/garage-web-portal-releases.db")
	if err != nil {
		log.Fatalf("failed to open store: %v", err)
	}

	s := server.New(cfg, db)

	log.Println("Starting portal on :8080")
	if err := s.Start(":8080"); err != nil {
		log.Fatal(err)
	}
}
