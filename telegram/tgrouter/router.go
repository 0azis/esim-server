package tgrouter

import (
	"context"
	"esim/internal/app"
	"esim/internal/store"

	"github.com/0azis/bot"
	"github.com/0azis/bot/models"
)

type tgRoutes interface {
	Init()

	// handlers
	startHandler(ctx context.Context, bot *bot.Bot, update *models.Update)
}

type tgRouter struct {
	app app.App
}

func (tg tgRouter) startHandler(ctx context.Context, b *bot.Bot, update *models.Update) {
	user := store.NewTelegramUser(int(update.Message.From.ID))
	_, err := tg.app.Store.User().Create(user)
	if err != nil {

	}

	b.SendMessage(ctx, &bot.SendMessageParams{
		Text:   "Hello. User Created!",
		ChatID: update.Message.From.ID,
	})
}

func (tg tgRouter) Init() {
	tg.app.Bot.Instance().RegisterHandler(bot.HandlerTypeMessageText, "/start", bot.MatchTypeExact, tg.startHandler)
}

func NewTgRouter(app app.App) tgRoutes {
	return tgRouter{app}
}
