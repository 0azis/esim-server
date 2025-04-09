package telegram

import (
	"context"
	"esim/config"
	"fmt"
	"time"

	"github.com/0azis/bot"
	initdata "github.com/telegram-mini-apps/init-data-golang"
)

type botUtils interface {
	ValidateInitData(initData string) bool
	ParseInitData(initData string) (int, error)
}

type botUtiler struct {
	cfg config.Config

	b *bot.Bot
}

func (bu botUtiler) ValidateInitData(initData string) bool {
	expIn := 24 * time.Hour
	err := initdata.Validate(initData, bu.cfg.TelegramToken, expIn)
	return err == nil
}

func (bu botUtiler) ParseInitData(initData string) (int, error) {
	i, err := initdata.Parse(initData)
	if err != nil {
		return 0, err
	}

	return int(i.User.ID), nil
}

func (bu botUtiler) Error(userID int64, text string) {
	bu.b.SendMessage(context.Background(), &bot.SendMessageParams{
		ChatID: userID,
		Text:   fmt.Sprintf("❌ Произошла ошибка!\n\n%s", text),
	})
}

func newBotUtils(cfg config.Config, b *bot.Bot) botUtils {
	return botUtiler{cfg, b}
}
