package telegram

import (
	"context"
	"fmt"
	"time"

	"github.com/timsofteng/jeka/lib/logger"
	"github.com/timsofteng/jeka/services/telegram/ports"

	"golang.org/x/sync/errgroup"
	tele "gopkg.in/telebot.v4"
)

type Telegram struct {
	bot         *tele.Bot
	logger      logger.Logger
	services    ports.Services
	ownUsername string
}

func New(
	ctx context.Context,
	logger logger.Logger,
	token string,
	ownUsername string,
	services ports.Services,
) (*Telegram, error) {
	const pollerTimeout = 10 * time.Second

	pref := tele.Settings{
		OnError: onError(logger),
		Token:   token,
		Poller: &tele.LongPoller{
			Limit:   0,
			Timeout: pollerTimeout,
		},
	}

	gErrGr, ctx := errgroup.WithContext(ctx)

	var bot *tele.Bot

	gErrGr.Go(func() error {
		var err error

		bot, err = tele.NewBot(pref)
		if err != nil {
			return fmt.Errorf("failed to init tg bot: %w", err)
		}

		return nil
	})

	if err := gErrGr.Wait(); err != nil {
		return nil, fmt.Errorf("error group encountered an error: %w", err)
	}

	telegram := &Telegram{
		bot:         bot,
		logger:      logger,
		services:    services,
		ownUsername: ownUsername,
	}

	telegram.handlers(ctx)

	return telegram, nil
}

func onError(logger logger.Logger) func(error, tele.Context) {
	return func(err error, c tele.Context) {
		logger.Error("unhandled telebot error", "error", err)

		err = c.Send(errCommon.Error())
		if err != nil {
			logger.Error("failed to send client error in error handler",
				"err", err,
			)
		}
	}
}

func (t *Telegram) Start() {
	t.bot.Start()
}

func (t *Telegram) Stop() {
	t.bot.Stop()
}
