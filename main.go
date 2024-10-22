package main

import (
	"fmt"
	"log"
	"os"
	"telegraminput/lib/logger"
	"telegraminput/services/images"
	"telegraminput/services/images/adapters/unsplash"
	"telegraminput/services/telegram"
	tgAdapters "telegraminput/services/telegram/adapters"
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

	tgAdoptedServices := tgAdapters.New(tgAdapters.Services{
		Video: videoSrv, Image: imgSrv,
	})

	telegram, err := telegram.New(logger,
		cfg.Token,
		tgAdoptedServices,
	)
	if err != nil {
		return fmt.Errorf("failed to init telegram service: %w", err)
	}

	telegram.Start()

	return nil
}
