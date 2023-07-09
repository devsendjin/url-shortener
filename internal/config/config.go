package config

import (
	"log"
	"os"
	"time"

	"github.com/ilyakaznacheev/cleanenv"
)

type Config struct {
	Env         string `yaml:"env" env:"ENV" env-default:"local"`
	StoragePath string `yaml:"storage_path" env-required:"true"`
	HTTPServer  `yaml:"http_server"`
}

type HTTPServer struct {
	Address  string        `yaml:"address" env-default:"localhost:8080"`
	TimeOut  time.Duration `yaml:"timeout" env-default:"4s"`
	IdleTime time.Duration `yaml:"idle_timeout" env-default:"60s"`
}

func MustLoad() *Config {
	// TODO: load env file and set env vars
	// configPath := os.Getenv("CONFIG_PATH")

	// if configPath == "" {
	// 	log.Fatal("CONFIG_PATH is not set")
	// }

	configPath := "./config/local.yaml"

	if _, err := os.Stat(configPath); os.IsNotExist(err) {
		log.Fatalf("Config file not found: %s", configPath)
	}

	var cfg Config

	if err := cleanenv.ReadConfig(configPath, &cfg); err != nil {
		log.Fatalf("Failed to load config: %s", err)
	}

	return &cfg
}
