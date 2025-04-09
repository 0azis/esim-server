package main

import (
	"esim/config"
	"esim/internal/app"
	"esim/internal/http/httprouter"
	"esim/telegram/tgrouter"
	"os"
	"time"

	"github.com/charmbracelet/log"
	"github.com/subosito/gotenv"
)

func main() {
	// set default logger
	l := log.NewWithOptions(os.Stdout, log.Options{
		ReportCaller:    true,
		ReportTimestamp: true,
		TimeFormat:      time.Kitchen,
	})
	l.Info("Configured logger...")

	// load env
	if err := gotenv.Load("../.env"); err != nil {
		l.Error("Error while loading .env variables", err.Error())
		return
	}
	l.Info(".env loaded...")

	// get config
	cfg := config.New()
	l.Info("Config loaded...")

	// setup main app obj
	app, err := app.New(cfg)
	if err != nil {
		l.Error("Error while configuring main App struct.", err.Error())
		return
	}
	l.Info("App is initialized...")
	app.SetLogger(l)

	// run bot
	tgrouter := tgrouter.NewTgRouter(app)
	tgrouter.Init()
	go app.Bot.Run()
	l.Info("Telegram Bot is running...")

	// run http with routes
	httprouter := httprouter.NewRouter(app)
	httprouter.Init()
	l.Info("HTTP server with routes is running...")

	err = app.Server.Run()
	if err != nil {
		app.Logger.Error(err.Error())
		return
	}
}
