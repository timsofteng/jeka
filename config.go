package main

import "github.com/caarlos0/env/v11"

type Config struct {
	Token          string `env:"TOKEN,notEmpty"`
	YtAPIKey       string `env:"YT_API_KEY,notEmpty"`
	UnsplashAPIKey string `env:"UNSPLASH_API_KEY,notEmpty"`
	LogLevel       string `env:"LOG_LEVEL"`
}

func ReadConfig() (Config, error) {
	return env.ParseAs[Config]()
}
