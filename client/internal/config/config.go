package config

import (
	"flag"
	"fmt"
	"os"

	"github.com/avGenie/gophkeeper/package/validator"
	"gopkg.in/yaml.v3"
)

// Handler config errors
//
// ErrGetConfigPath - returned if config file path could not be gotten
var (
	ErrGetConfigPath = fmt.Errorf("couldn't get config file path from environment")
)

// const (
// 	grpcConnection = ":9090"
// 	logLevel       = "debug"
// 	logPath        = "client.log"
// 	certPath       = "/home/alexander/projects/gophkeeper/certs/"
// )

type Config struct {
	Client ClientConfig `json:"client"  yaml:"client"`
}

// ClientConfig Client configuration
type ClientConfig struct {
	CertsPath   string `json:"certs_path" yaml:"certs_path"`
	GRPCAddress string `json:"grpc_addr"  yaml:"grpc_addr"`
	LogLevel    string `json:"log_level"  yaml:"log_level"`
	LogPath     string `json:"log_path"   yaml:"log_path"`
}

// func NewConfig() Config {
// 	return Config{
// 		Client: ClientConfig{
// 			CertsPath:   certPath,
// 			GRPCAddress: grpcConnection,
// 			LogLevel:    logLevel,
// 			LogPath:     logPath,
// 		},
// 	}
// }

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

	if err := validator.ValidateConfigPath(configPath); err != nil {
		return "", err
	}

	return configPath, nil
}
