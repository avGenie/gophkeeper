// Package config implements server config
package config

import (
	"flag"
	"fmt"
	"os"

	"github.com/avGenie/gophkeeper/server/internal/validation"
	"gopkg.in/yaml.v3"
)

// Handler config errors
//
// ErrGetConfigPath - returned if config file path could not be gotten
var (
	ErrGetConfigPath = fmt.Errorf("couldn't get config file path from environment")
)

// Config Server configuration
type Config struct {
	Server  ServerConfig  `json:"server"  yaml:"server"`
	Storage StorageConfig `json:"storage" yaml:"storage"`
}

type ServerConfig struct {
	GRPCAddress string `json:"grpc_addr" yaml:"grpc_addr"`
	LogLevel    string `json:"log_level" yaml:"log_level"`
}

type StorageConfig struct {
	DatabaseDSN string `json:"database_dsn" yaml:"database_dsn"`
	MaxPool     int32  `json:"max_pool"     yaml:"max_pool"`
}

// NewConfig Creates new config from yaml file getting from env variable
func NewConfig() (Config, error) {
	configPath, err := parseConfigPathEnv()
	if err != nil {
		return Config{}, ErrGetConfigPath
	}

	file, err := os.Open(configPath)
	if err != nil {
		return Config{}, fmt.Errorf("couldn't open config file: %w", err)
	}
	defer file.Close()

	d := yaml.NewDecoder(file)

	var config Config
	if err := d.Decode(&config); err != nil {
		return Config{}, fmt.Errorf("couldn't decode config file: %w", err)
	}

	return config, nil
}

func parseConfigPathEnv() (string, error) {
	var configPath string

	flag.StringVar(&configPath, "c", "./config.yml", "path to config file")

	flag.Parse()

	if err := validation.ValidateConfigPath(configPath); err != nil {
		return "", err
	}

	return configPath, nil
}
