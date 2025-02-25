package sqlite

import (
	"github.com/roka-crew/pkg/config"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type SQLite struct {
	*gorm.DB
}

func New(cfg *config.Config) (*SQLite, error) {
	db, err := gorm.Open(sqlite.Open(cfg.Persistence.Path), &gorm.Config{
		Logger:         logger.Default.LogMode(logger.Silent),
		TranslateError: true,
	})
	if err != nil {
		return nil, err
	}

	return &SQLite{DB: db}, nil
}
