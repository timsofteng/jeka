package main

import (
	"fmt"
	"log"
	"os"
	"telegraminput/internal/logger"
	"telegraminput/services/images"
	"telegraminput/services/youtube"
	"telegraminput/transport/telegram"
)

func main() {
	err := run()
	if err != nil {
		log.Printf("root err: %+v", err)
		os.Exit(1)
	}
}

func run() error {
	cfg, err := ReadConfig()
	if err != nil {
		return err
	}

	logger := logger.New(cfg.LogLevel)

	ytSrv, err := youtube.New(logger, cfg.YtAPIKey)
	if err != nil {
		return fmt.Errorf("failed to init youtube service: %w", err)
	}

	imgSrv := images.New(cfg.UnsplashAPIKey)

	telegram, err := telegram.New(logger,
		cfg.Token,
		telegram.Services{
			Video: ytSrv,
			Image: imgSrv,
		},
	)
	if err != nil {
		return fmt.Errorf("failed to init telegram service: %w", err)
	}

	telegram.Start()

	return nil
}
