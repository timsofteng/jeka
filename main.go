package main

import (
	"fmt"
	"log"
	"os"
	"telegraminput/lib/logger"
	"telegraminput/lib/postgres"
	"telegraminput/services/httpserver"
	httpSrvAdapter "telegraminput/services/httpserver/adapters/services"
	"telegraminput/services/images"
	"telegraminput/services/images/adapters/unsplash"
	"telegraminput/services/telegram"
	tgAdapters "telegraminput/services/telegram/adapters/services"
	"telegraminput/services/text"
	textPGAdapter "telegraminput/services/text/adapters/postgres"
	"telegraminput/services/video"
	"telegraminput/services/video/adapters/youtube"
	"telegraminput/services/voice"
	voicePGAdapter "telegraminput/services/voice/adapters/postgres"
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

	postgres, err := postgres.New(cfg.DatabaseURL)
	if err != nil {
		return fmt.Errorf("failed to init postgres db: %w", err)
	}

	textRepo := textPGAdapter.New(postgres.Conn)
	textSrv := text.New(textRepo)

	voiceRepo := voicePGAdapter.New(postgres.Conn)
	voiceSrv := voice.New(voiceRepo)

	tgAdoptedServices := tgAdapters.New(tgAdapters.Services{
		Video: videoSrv, Image: imgSrv, Text: textSrv, Voice: voiceSrv,
	})

	telegram, err := telegram.New(logger,
		cfg.Token,
		cfg.BotOwnUsername,
		tgAdoptedServices,
	)
	if err != nil {
		return fmt.Errorf("failed to init telegram service: %w", err)
	}

	go telegram.Start()

	httpSrv := httpSrvAdapter.New(httpSrvAdapter.Services{
		Video: videoSrv, Image: imgSrv, Text: textSrv, Voice: voiceSrv,
	})

	httpServer, err := httpserver.New(
		logger, cfg.HTTPServerHost, cfg.HTTPServerPort, httpSrv,
	)
	if err != nil {
		return fmt.Errorf("failed to init http server: %w", err)
	}

	httpServer.Start()

	return nil
}
