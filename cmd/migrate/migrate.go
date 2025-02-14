package main

import (
	"log"

	"github.com/roka-crew/domain"
	"github.com/roka-crew/pkg/config"
	"github.com/roka-crew/pkg/persistence/sqlite"
)

func main() {
	cfg, err := config.New("./configs/config.yaml")
	if err != nil {
		panic(err)
	}

	db, err := sqlite.New(cfg)
	if err != nil {
		panic(err)
	}

	log.Printf(">>> start migrate")
	if err := db.AutoMigrate(&domain.User{}, &domain.Group{}, &domain.Goal{}, &domain.Topic{}); err != nil {
		panic(err)
	}
	log.Printf(">>> end migrate")
}
