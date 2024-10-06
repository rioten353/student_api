package config

import (
	"flag"
	"log"
	"os"

	"github.com/ilyakaznacheev/cleanenv"
)

type HttpServer struct {
	Address string `yaml:"address"`
}

type Config struct {
	Env         string `yaml:"env" env:"ENV" env-require:"true"`
	StoragePath string `yaml:"storage_path" env-require:"true"`
	HttpServer  `yaml:"http_server"`
}

func MustLoad() *Config {
	var configPath string

	configPath = os.Getenv("CONFIG_PATH")

	if configPath == "" {
		flags := flag.String("config", "", "config file path")
		flag.Parse()
		configPath = *flags

		if configPath == "" {
			log.Fatal("config file path is required")
		}
	}
	if _, err := os.Stat(configPath); os.IsNotExist(err) {
		log.Fatal("config file not found")
	}
	var cfg Config

	if err := cleanenv.ReadConfig(configPath, &cfg); err != nil {
		log.Fatalf("failed to read config file %s", err.Error())
	}

	return &cfg
}
