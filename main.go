package main

import (
	"fmt"
	"log"
	"os"
	"telegraminput/lib/logger"
	"telegraminput/services/images"
	"telegraminput/services/images/adapters/unsplash"
	"telegraminput/services/telegram"
	tgAdapters "telegraminput/services/telegram/adapters/services"
	"telegraminput/services/text"
	"telegraminput/services/text/adapters/postgres"
	"telegraminput/services/video"
	"telegraminput/services/video/adapters/youtube"
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

	videoRepo, err := youtube.New(logger, cfg.YtAPIKey)
	if err != nil {
		return fmt.Errorf("failed to init youtube repo adapter: %w", err)
	}

	videoSrv := video.New(videoRepo)

	imgRepo := unsplash.New(cfg.UnsplashAPIKey)
	imgSrv := images.New(imgRepo)

	textRepo, err := postgres.New(cfg.DatabaseURL)
	if err != nil {
		return fmt.Errorf("failed to init postgres text repo adapter: %w", err)
	}

	textSrv := text.New(textRepo)

	tgAdoptedServices := tgAdapters.New(tgAdapters.Services{
		Video: videoSrv, Image: imgSrv, Text: textSrv,
	})

	telegram, err := telegram.New(logger,
		cfg.Token,
		cfg.BotOwnUsername,
		tgAdoptedServices,
	)
	if err != nil {
		return fmt.Errorf("failed to init telegram service: %w", err)
	}

	telegram.Start()

	return nil
}
