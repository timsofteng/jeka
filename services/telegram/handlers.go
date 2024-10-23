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
			t.logger.Error("error to call taksa service", "err", err)

			return ctxTb.Send(errCommon.Error())
		}

		return ctxTb.Send(resp)
	})

	t.bot.Handle("/voice", func(ctxTb tele.Context) error {
		ctx, cancel := context.WithTimeout(context.Background(), defaultTimeout)
		defer cancel()

		voiceID, err := t.services.RandVoice(ctx)
		if err != nil {
			t.logger.Error("error to call image/taksa service", "err", err)

			return ctxTb.Send(errCommon.Error())
		}

		voice := &tele.Voice{File: tele.File{FileID: voiceID}}

		return ctxTb.Reply(voice)
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

// t.bot.Handle(tele.OnVoice, func(ctxTb tele.Context) error {
// 	voiceID := ctxTb.Message().Voice.FileID
// 	file, err := os.OpenFile(
// "voices.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
// 	defer file.Close()
//
// 	if err != nil {
// 		log.Printf("Failed to open file: %v", err)
//
// 		return err
// 	}
//
// 	_, err = file.WriteString(voiceID + "\n")
// 	if err != nil {
// 		log.Printf("Failed to write to file: %v", err)
//
// 		return err
// 	}
//
// 	log.Printf("Voice ID saved: %s", voiceID)
//
// 	return nil
// })
