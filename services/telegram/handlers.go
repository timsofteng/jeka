package telegram

import (
	"context"
	"time"

	tele "gopkg.in/telebot.v4"
)

const defaultTimeout = 2 * time.Second

//nolint:funlen
func (t *Telegram) handlers() {
	t.bot.Handle("/test", func(c tele.Context) error {
		return c.Send("Hello!")
	})

	t.bot.Handle("/video", func(ctxTb tele.Context) error {
		ctx, cancel := context.WithTimeout(context.Background(), defaultTimeout)
		defer cancel()

		resp, err := t.services.RandVideo(ctx)
		if err != nil {
			t.logger.Error("/video handler", "err", err)

			return ctxTb.Send(errCommon.Error())
		}

		return ctxTb.Send(resp)
	})

	t.bot.Handle("/random-img", func(ctxTb tele.Context) error {
		ctx, cancel := context.WithTimeout(context.Background(), defaultTimeout)
		defer cancel()

		resp, err := t.services.RandImg(ctx)
		if err != nil {
			t.logger.Error("error to call image/randomimg service", "err", err)

			return ctxTb.Send(errCommon.Error())
		}

		return ctxTb.Send(resp)
	})

	t.bot.Handle("/taksa", func(ctxTb tele.Context) error {
		ctx, cancel := context.WithTimeout(context.Background(), defaultTimeout)
		defer cancel()

		resp, err := t.services.Taksa(ctx)
		if err != nil {
			t.logger.Error("error to call image/taksa service", "err", err)

			return ctxTb.Send(errCommon.Error())
		}

		return ctxTb.Send(resp)
	})

	textHandler := func(ctxTb tele.Context) error {
		ctx, cancel := context.WithTimeout(context.Background(), defaultTimeout)
		defer cancel()

		resp, err := t.services.RandText(ctx)
		if err != nil {
			t.logger.Error("error to call rand text service", "err", err)

			return ctxTb.Send(errCommon.Error())
		}

		t.logger.Info("responding to client", "resp", resp)

		return ctxTb.Reply(resp)
	}
	t.bot.Handle("/text", textHandler)
	t.bot.Handle("/jeka", textHandler)

	t.bot.Handle(tele.OnReply, func(ctxTb tele.Context) error {
		if ctxTb.Message().ReplyTo.Sender.Username == t.ownUsername {
			ctx, cancel := context.WithTimeout(context.Background(), defaultTimeout)
			defer cancel()

			resp, err := t.services.RandText(ctx)
			if err != nil {
				t.logger.Error("error to call rand text service", "err", err)

				return ctxTb.Send(errCommon.Error())
			}

			t.logger.Info("responding to client", "resp", resp)

			return ctxTb.Reply(resp)
		}

		return nil
	})
}
