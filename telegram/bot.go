package telegram

import (
	"context"
	"esim/config"
	"esim/internal/store"

	"github.com/0azis/bot"
)

type Bot interface {
	Run()
	Instance() *bot.Bot
	Utils() botUtils
}

type tgBot struct {
	utils botUtils
	b     *bot.Bot
}

func NewBot(cfg config.Config, store store.Store) (Bot, error) {
	b, err := bot.New(cfg.TelegramToken)
	if err != nil {
		return tgBot{}, err
	}

	tgBot := tgBot{
		utils: newBotUtils(cfg, b),
		b:     b,
	}

	return tgBot, nil
}

func (tb tgBot) Run() {
	tb.b.Start(context.Background())
}

func (tb tgBot) Instance() *bot.Bot {
	return tb.b
}

func (tb tgBot) Utils() botUtils {
	return tb.utils
}
