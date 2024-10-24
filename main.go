package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/timsofteng/jeka/lib/logger"
	"github.com/timsofteng/jeka/lib/postgres"
	"github.com/timsofteng/jeka/services/httpserver"
	httpSrvAdapter "github.com/timsofteng/jeka/services/httpserver/adapters/services"
	"github.com/timsofteng/jeka/services/images"
	"github.com/timsofteng/jeka/services/images/adapters/unsplash"
	"github.com/timsofteng/jeka/services/telegram"
	tgAdapters "github.com/timsofteng/jeka/services/telegram/adapters/services"
	"github.com/timsofteng/jeka/services/text"
	textPGAdapter "github.com/timsofteng/jeka/services/text/adapters/postgres"
	"github.com/timsofteng/jeka/services/video"
	"github.com/timsofteng/jeka/services/video/adapters/youtube"
	"github.com/timsofteng/jeka/services/voice"
	voicePGAdapter "github.com/timsofteng/jeka/services/voice/adapters/postgres"

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

	httpSrv := httpSrvAdapter.New(logger, httpSrvAdapter.Services{
		Video: videoSrv, Image: imgSrv, Text: textSrv, Voice: voiceSrv,
	})

	httpServer, err := httpserver.New(
		ctx, logger, cfg.HTTPServerHost, cfg.HTTPServerPort, httpSrv,
	)
	if err != nil {
		return fmt.Errorf("failed to init http server: %w", err)
	}

	errGr, gCtx := errgroup.WithContext(ctx)

	errGr.Go(func() error {
		logger.Info("starting telegram server")
		telegram.Start()

		return nil
	})

	errGr.Go(func() error {
		logger.Info("starting http server")

		return httpServer.Start()
	})

	// shutdown goroutine
	errGr.Go(func() error {
		<-gCtx.Done()
		telegram.Stop()

		err := httpServer.Stop(ctx)
		if err != nil {
			return fmt.Errorf("failed to stop http server: %w", err)
		}

		logger.Info("tg and http servers stoped correct")

		return nil
	})

	if err := errGr.Wait(); err != nil {
		return fmt.Errorf("error group encountered an error: %w", err)
	}

	return nil
}
