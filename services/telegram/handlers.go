package telegram

import (
	"context"
	"time"

	tele "gopkg.in/telebot.v4"
)

func (t *Telegram) handlers() {
	t.bot.Handle("/test", func(c tele.Context) error {
		return c.Send("Hello!")
	})

	t.bot.Handle("/video", func(ctx tele.Context) error {
		resp, err := t.services.RandVideo()
		if err != nil {
			t.logger.Error("/video handler", "err", err)

			return ctx.Send(errCommon.Error())
		}

		return ctx.Send(resp)
	})

	t.bot.Handle("/random-img", func(ctxTb tele.Context) error {
		const timeout = 2 * time.Second

		ctx, cancel := context.WithTimeout(context.Background(), timeout)
		defer cancel()

		resp, err := t.services.RandImg(ctx)
		if err != nil {
			t.logger.Error("error to call image/randomimg service", "err", err)

			return ctxTb.Send(errCommon.Error())
		}

		return ctxTb.Send(resp)
	})

	t.bot.Handle("/taksa", func(ctxTb tele.Context) error {
		const timeout = 2 * time.Second

		ctx, cancel := context.WithTimeout(context.Background(), timeout)
		defer cancel()

		resp, err := t.services.Taksa(ctx)
		if err != nil {
			t.logger.Error("error to call image/taksa service", "err", err)

			return ctxTb.Send(errCommon.Error())
		}

		return ctxTb.Send(resp)
	})
}
