package sqlite

import (
	"github.com/roka-crew/pkg/config"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type SQLite struct {
	*gorm.DB
}

func New(cfg *config.Config) (*SQLite, error) {
	db, err := gorm.Open(sqlite.Open(cfg.Persistence.Path), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	return &SQLite{DB: db}, nil
}
