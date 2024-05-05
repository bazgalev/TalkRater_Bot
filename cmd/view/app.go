package view

import (
	tele "gopkg.in/telebot.v3"
	"log/slog"
	"talk_rater_bot/internal/databases"
	"talk_rater_bot/internal/helpers"
	"talk_rater_bot/internal/templates"
)

type Application struct {
	Logger         *slog.Logger
	UserBot        *tele.Bot
	AdminBot       *tele.Bot
	AdminDB        *databases.AdminDB
	AdminTemplates *templates.Templates
	UserTemplates  *templates.Templates
	TimeParser     *helpers.TimeParser
	PathTmp        string
}

func (app *Application) Run() {
	go app.UserBot.Start()
	go app.AdminBot.Start()
}
