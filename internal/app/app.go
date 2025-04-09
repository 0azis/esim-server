package app

import (
	"esim/config"
	"esim/internal/http"
	"esim/internal/mail"
	"esim/internal/redis"
	"esim/internal/store"
	"esim/pkg"
	"esim/telegram"

	"github.com/charmbracelet/log"
)

type App struct {
	Store  store.Store
	Redis  redis.Redis
	Server http.Server
	Mailer mail.Mailer
	Bot    telegram.Bot

	Logger *log.Logger
	Jwt    pkg.Jwt

	Config config.Config
}

func New(cfg config.Config) (App, error) {
	var app App
	store, err := store.New(cfg)
	if err != nil {
		return app, err
	}

	redis := redis.New(cfg)
	http := http.New(cfg.Http.Addr())
	bot, err := telegram.NewBot(cfg, store)
	if err != nil {
		return app, err
	}

	jwtBuilder := pkg.NewJwtBuilder(cfg)

	app.Store = store
	app.Redis = redis
	app.Server = http
	app.Bot = bot

	app.Jwt = jwtBuilder

	app.Config = cfg

	return app, nil
}

func (a *App) SetLogger(l *log.Logger) {
	a.Logger = l
}
