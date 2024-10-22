package telegram

import (
	"fmt"
	"telegraminput/lib/logger"
	"telegraminput/services/telegram/ports"
	"time"

	tele "gopkg.in/telebot.v4"
)

type Telegram struct {
	bot      *tele.Bot
	logger   logger.Logger
	services ports.Services
}

func New(
	logger logger.Logger, token string, services ports.Services,
) (*Telegram, error) {
	const pollerTimeout = 10 * time.Second

	pref := tele.Settings{
		Token: token,
		Poller: &tele.LongPoller{
			Timeout: pollerTimeout,
		},
	}

	bot, err := tele.NewBot(pref)
	if err != nil {
		return nil, fmt.Errorf("failed to init tg bot: %w", err)
	}

	tg := &Telegram{bot: bot, logger: logger, services: services}

	tg.handlers()

	return tg, nil
}

func (t *Telegram) Start() {
	t.bot.Start()
}

func (t *Telegram) Stop() {
	t.bot.Stop()
}
