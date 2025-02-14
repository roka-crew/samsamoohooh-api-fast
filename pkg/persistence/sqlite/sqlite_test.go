package sqlite

import (
	"testing"

	"github.com/roka-crew/pkg/config"
	"github.com/stretchr/testify/assert"
)

func TestNew(t *testing.T) {
	cfg := &config.Config{
		Persistence: config.Persistence{
			Path: "test.db",
		},
	}

	_, err := New(cfg)
	assert.Nil(t, err)
}
