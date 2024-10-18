package telegram

import (
	"context"
	"fmt"
	"telegraminput/internal/logger"
	"time"

	tele "gopkg.in/telebot.v4"
)

type Telegram struct{ Bot *tele.Bot }

type Services struct {
	Video VideoService
	Image ImageService
}

type VideoService interface {
	RandVideo() (url string, caption string, err error)
}

type ImageService interface {
	RandomImg(ctx context.Context) (url string, caption string, err error)
	Taksa(ctx context.Context) (url string, caption string, err error)
}

func New(
	logger logger.Logger, token string, services Services,
) (*Telegram, error) {
	const pollerTimeoutSeconds = 10

	pref := tele.Settings{
		Token: token,
		Poller: &tele.LongPoller{
			Timeout: pollerTimeoutSeconds * time.Second,
		},
	}

	bot, err := tele.NewBot(pref)
	if err != nil {
		return nil, fmt.Errorf("failed to init tg bot: %w", err)
	}

	bot.Handle("/jeka", func(c tele.Context) error {
		return c.Send("Hello!")
	})

	bot.Handle("/video", func(ctx tele.Context) error {
		url, caption, err := services.Video.RandVideo()
		if err != nil {
			logger.Error("/video handler", "err", err)

			return ctx.Send(errCommon.Error())
		}

		return ctx.Send(fmt.Sprintf("%s \n %s", caption, url))
	})

	bot.Handle("/random_img", func(ctxTb tele.Context) error {
		const timeout = 2 * time.Second

		ctx, cancel := context.WithTimeout(context.Background(), timeout)
		defer cancel()

		url, caption, err := services.Image.RandomImg(ctx)
		if err != nil {
			logger.Error("error to call image/randomimg service", "err", err)

			return ctxTb.Send(errCommon.Error())
		}

		return ctxTb.Send(fmt.Sprintf("%s \n\n %s", caption, url))
	})

	bot.Handle("/taksa", func(ctxTb tele.Context) error {
		const timeout = 2 * time.Second

		ctx, cancel := context.WithTimeout(context.Background(), timeout)
		defer cancel()

		url, caption, err := services.Image.Taksa(ctx)
		if err != nil {
			logger.Error("error to call image/taksa service", "err", err)

			return ctxTb.Send(errCommon.Error())
		}

		return ctxTb.Send(fmt.Sprintf("%s \n\n %s", caption, url))
	})

	return &Telegram{Bot: bot}, nil
}

func (t *Telegram) Start() {
	t.Bot.Start()
}
