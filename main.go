package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"
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

	"golang.org/x/sync/errgroup"
)

func main() {
	ctx, stop := signal.NotifyContext(
		context.Background(), syscall.SIGINT, syscall.SIGTERM)

	if err := run(ctx); err != nil {
		log.Printf("root err: %+v", err)
		stop()
		os.Exit(1)
	}

	stop()
}

//nolint:funlen
func run(ctx context.Context) error {
	cfg, err := ReadConfig()
	if err != nil {
		return err
	}

	logger := logger.New(cfg.LogLevel)

	videoRepo, err := youtube.New(ctx, logger, cfg.YtAPIKey)
	if err != nil {
		return fmt.Errorf("failed to init youtube repo adapter: %w", err)
	}

	videoSrv := video.New(videoRepo)

	imgRepo := unsplash.New(cfg.UnsplashAPIKey)
	imgSrv := images.New(imgRepo)

	postgres, err := postgres.New(ctx, cfg.DatabaseURL)
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

	telegram, err := telegram.New(
		ctx,
		logger,
		cfg.Token,
		cfg.BotOwnUsername,
		tgAdoptedServices,
	)
	if err != nil {
		return fmt.Errorf("failed to init telegram service: %w", err)
	}

	httpSrv := httpSrvAdapter.New(httpSrvAdapter.Services{
		Video: videoSrv, Image: imgSrv, Text: textSrv, Voice: voiceSrv,
	})

	httpServer, err := httpserver.New(
		ctx, logger, cfg.HTTPServerHost, cfg.HTTPServerPort, httpSrv,
	)
	if err != nil {
		return fmt.Errorf("failed to init http server: %w", err)
	}

	gErrGr, gCtx := errgroup.WithContext(ctx)

	gErrGr.Go(func() error {
		logger.Info("starting telegram server")
		telegram.Start()

		return nil
	})

	gErrGr.Go(func() error {
		logger.Info("starting http server")

		return httpServer.Start()
	})

	// shutdown goroutine
	gErrGr.Go(func() error {
		<-gCtx.Done()
		telegram.Stop()

		err := httpServer.Stop(ctx)
		if err != nil {
			return fmt.Errorf("failed to stop http server: %w", err)
		}

		logger.Info("tg and http servers stoped correct")

		return nil
	})

	if err := gErrGr.Wait(); err != nil {
		return fmt.Errorf("error group encountered an error: %w", err)
	}

	return nil
}
