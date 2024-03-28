package application

func (app *App) RegisterRoutes() error {
	app.r.Get("/orders", app.HandleCreateOrder)

	return nil
}
