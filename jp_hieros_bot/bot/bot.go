package bot

import (
	"log"
	"time"

	"gitlab.com/jp_hieros_bot/bot/handlers"
	"gitlab.com/jp_hieros_bot/config"
	tb "gopkg.in/telebot.v3"
)

func JpHierosBot() *tb.Bot {

	cfg := config.Load()

	prefSettings := tb.Settings{
		Token: cfg.BotToken,
		Poller: &tb.LongPoller{
			Timeout: 60 * time.Second,
		},
		ParseMode: tb.ModeHTML,
	}

	bot, err := tb.NewBot(prefSettings)
	if err != nil {
		log.Fatal(err.Error())
	}

	handV1 := handlers.NewHandlerV1(cfg)

	bot.Handle("/start", handV1.StartBot)
	bot.Handle("/challenge", handV1.Challenge)
	bot.Handle(tb.OnText, handV1.Answer)
	bot.Handle(tb.OnCallback, handV1.ReturnAnswer)

	return bot

}
