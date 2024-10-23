package main

import "github.com/caarlos0/env/v11"

type Config struct {
	Token          string `env:"TOKEN,notEmpty"`
	YtAPIKey       string `env:"YT_API_KEY,notEmpty"`
	UnsplashAPIKey string `env:"UNSPLASH_API_KEY,notEmpty"`
	LogLevel       string `env:"LOG_LEVEL"`
	DatabaseURL    string `env:"DATABASE_URL,notEmpty"`
	BotOwnUsername string `env:"BOT_OWN_USERNAME,notEmpty"`
	HTTPServerPort string `env:"HTTP_SERVER_PORT,notEmpty"`
	HTTPServerHost string `env:"HTTP_SERVER_HOST,notEmpty"`
}

func ReadConfig() (Config, error) {
	return env.ParseAs[Config]()
}
