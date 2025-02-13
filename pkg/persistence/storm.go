package persistence

import (
	"time"

	"github.com/asdine/storm"
	"github.com/roka-crew/pkg/config"
	"go.etcd.io/bbolt"
)

const (
	timeout = 1 * time.Second
)

type Storm struct {
	db *storm.DB
}

func NewStorm(cfg *config.Config) (*Storm, error) {
	options := &bbolt.Options{
		Timeout: timeout,
	}

	openedDB, err := storm.Open(cfg.Persistence.Path, storm.BoltOptions(0600, options))
	if err != nil {
		return nil, err
	}

	return &Storm{db: openedDB}, nil
}
