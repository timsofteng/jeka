package telegram

import (
	"context"
	"time"

	tele "gopkg.in/telebot.v4"
)

const defaultTimeout = 2 * time.Second

const helpMsg = `/help - инструкция
/rand - голосовое или текстовое
/voice - голосовое
/text - текст
/video - видео
/rand_img - случайное изображение
/taksa - такса`

//nolint:funlen
func (t *Telegram) handlers(ctx context.Context) {
	t.bot.Handle("/test", func(c tele.Context) error {
		return c.Send("Hello!")
	})

	t.bot.Handle("/help", func(c tele.Context) error {
		return c.Send(helpMsg)
	})

	t.bot.Handle("/video", func(ctxTb tele.Context) error {
		ctx, cancel := context.WithTimeout(ctx, defaultTimeout)
		defer cancel()

		video, err := t.services.RandVideo(ctx)
		if err != nil {
			t.logger.Error("/video handler", "err", err)

			return ctxTb.Send(errCommon.Error())
		}

		return ctxTb.Send(video)
	})

	t.bot.Handle("/rand_img", func(ctxTb tele.Context) error {
		ctx, cancel := context.WithTimeout(ctx, defaultTimeout)
		defer cancel()

		photo, err := t.services.RandImg(ctx)
		if err != nil {
			t.logger.Error("error to call image/randomimg service", "err", err)

			return ctxTb.Send(errCommon.Error())
		}

		return ctxTb.Send(photo)
	})

	t.bot.Handle("/taksa", func(ctxTb tele.Context) error {
		ctx, cancel := context.WithTimeout(ctx, defaultTimeout)
		defer cancel()

		photo, err := t.services.Taksa(ctx)
		if err != nil {
			t.logger.Error("error to call taksa service", "err", err)

			return ctxTb.Send(errCommon.Error())
		}

		return ctxTb.Send(photo)
	})

	t.bot.Handle("/rand", func(ctxTb tele.Context) error {
		ctx, cancel := context.WithTimeout(ctx, defaultTimeout)
		defer cancel()

		resp, err := t.services.Rand(ctx)
		if err != nil {
			t.logger.Error("error to call rand service", "err", err)

			return ctxTb.Send(errCommon.Error())
		}

		return ctxTb.Reply(resp)
	})

	t.bot.Handle("/voice", func(ctxTb tele.Context) error {
		ctx, cancel := context.WithTimeout(ctx, defaultTimeout)
		defer cancel()

		voice, err := t.services.RandVoice(ctx)
		if err != nil {
			t.logger.Error("error to call image/taksa service", "err", err)

			return ctxTb.Send(errCommon.Error())
		}

		return ctxTb.Reply(voice)
	})

	textHandler := func(ctxTb tele.Context) error {
		ctx, cancel := context.WithTimeout(ctx, defaultTimeout)
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
			ctx, cancel := context.WithTimeout(ctx, defaultTimeout)
			defer cancel()

			resp, err := t.services.Rand(ctx)
			if err != nil {
				t.logger.Error("error to call rand service", "err", err)

				return ctxTb.Send(errCommon.Error())
			}

			t.logger.Info("responding to client", "resp", resp)

			return ctxTb.Reply(resp)
		}

		return nil
	})
}
