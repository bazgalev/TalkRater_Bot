package main

import (
	tele "gopkg.in/telebot.v3"
	"log/slog"
	"time"
)

func (app *application) recoverPanic(next tele.HandlerFunc) tele.HandlerFunc {
	const op = "middleware.logRequest"

	return func(c tele.Context) error {
		defer func() {
			if r := recover(); r != nil {
				app.logger.Warn(op,
					slog.String("username", c.Sender().Username),
					slog.String("panic", r.(error).Error()))

				c.Send("internal server error")
			}
		}()

		return next(c)
	}
}

func (app *application) measureTime(next tele.HandlerFunc) tele.HandlerFunc {
	const op = "middleware.measureTime"

	return func(c tele.Context) error {
		timeStart := time.Now()

		result := next(c)

		timeEnd := time.Now()

		app.logger.Info(op,
			slog.String("username", c.Sender().Username),
			slog.String("text", c.Text()),
			slog.Int64("time ms", timeEnd.Sub(timeStart).Milliseconds()),
		)

		return result
	}
}

func (app *application) checkAdmin(next tele.HandlerFunc) tele.HandlerFunc {
	const op = "middleware.checkAdmin"

	return func(c tele.Context) error {
		username := c.Sender().Username
		if app.adminDB.IsAdmin(username) {
			return next(c)
		} else {
			app.logger.Info(op,
				slog.String("username", username),
				slog.String("info", "failed authorization"))
			return c.Send("access denied") // TODO: add return access denied
		}
	}
}
