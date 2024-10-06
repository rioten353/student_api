package config

import (
	"flag"
	"log/slog"
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
			slog.Info("config file path is required")
		}
	}
	if _, err := os.Stat(configPath); os.IsNotExist(err) {
		slog.Info("config file not found: " + configPath)
	}
	var cfg Config

	if err := cleanenv.ReadConfig(configPath, &cfg); err != nil {
		slog.Error("failed to read config file", slog.String("error", err.Error()))
	}

	return &cfg
}
