package main

import (
	"fmt"
	"os"
	"path/filepath"

	"gopkg.in/yaml.v3"

	"github.com/DostonAkhmedov/async-logger/pkg/alog"
)

const ConfigFileEnv = "CONFIG_FILE"

type Config struct {
	ALog *alog.Config `yaml:"a_log"`
}

func NewConfig() (*Config, error) {
	var config Config

	path := filepath.Clean(os.Getenv(ConfigFileEnv))

	data, err := os.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf("read file: %w", err)
	}

	data = []byte(os.ExpandEnv(string(data)))

	err = yaml.Unmarshal(data, &config)
	if err != nil {
		return nil, fmt.Errorf("unmarshal: %w", err)
	}

	return &config, nil
}
