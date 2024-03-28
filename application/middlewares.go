package application

import "github.com/go-chi/chi/v5/middleware"

func (app *App) RegisterMiddlewares() {
	app.r.Use(middleware.Logger)
}
