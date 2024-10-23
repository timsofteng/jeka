package telegram

import (
	"fmt"
	"telegraminput/lib/logger"
	"telegraminput/services/telegram/ports"
	"time"

	tele "gopkg.in/telebot.v4"
)

type Telegram struct {
	bot         *tele.Bot
	logger      logger.Logger
	services    ports.Services
	ownUsername string
}

func New(
	logger logger.Logger,
	token string,
	ownUsername string,
	services ports.Services,
) (*Telegram, error) {
	const pollerTimeout = 10 * time.Second

	pref := tele.Settings{
		Token: token,
		Poller: &tele.LongPoller{
			Limit:   0,
			Timeout: pollerTimeout,
		},
	}

	bot, err := tele.NewBot(pref)
	if err != nil {
		return nil, fmt.Errorf("failed to init tg bot: %w", err)
	}

	telegram := &Telegram{
		bot:         bot,
		logger:      logger,
		services:    services,
		ownUsername: ownUsername,
	}

	telegram.handlers()

	return telegram, nil
}

func (t *Telegram) Start() {
	t.bot.Start()
}

func (t *Telegram) Stop() {
	t.bot.Stop()
}
