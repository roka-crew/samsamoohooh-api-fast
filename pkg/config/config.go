package config

import (
	"os"

	"gopkg.in/yaml.v3"
)

type Config struct {
	Listen      string      `yaml:"listen"`
	Token       Token       `yaml:"token"`
	Persistence Persistence `yaml:"persistence"`
}

type Token struct {
	SecretKey []byte `yaml:"secretKey"`
}

func (t *Token) UnmarshalYAML(value *yaml.Node) error {
	var key string
	if err := value.Decode(&key); err != nil {
		return err
	}
	t.SecretKey = []byte(key)
	return nil
}

type Persistence struct {
	Path string `yaml:"path"`
}

func New(path string) (*Config, error) {
	cfg := &Config{}

	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	err = yaml.NewDecoder(file).Decode(cfg)
	if err != nil {
		return nil, err
	}

	return cfg, nil
}
