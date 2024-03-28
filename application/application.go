package application

import (
	"applicationDesignTest/repository"

	"errors"
	"fmt"
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
)

type App struct {
	Orders repository.OrdersRepository
	Rooms  repository.RoomsRepository
	logger *log.Logger
	r      *chi.Mux
}

func NewApp(
	Orders repository.OrdersRepository,
	Rooms repository.RoomsRepository,
	logger *log.Logger,
) *App {
	r := chi.NewRouter()

	return &App{
		Orders: Orders,
		Rooms:  Rooms,
		logger: logger,
		r:      r,
	}
}

func (app *App) Run(host, port string) error {
	app.LogInfo("Server listening on %v:%v", host, port)
	app.RegisterMiddlewares()
	err := app.RegisterRoutes()
	if err != nil {
		return fmt.Errorf("could not register routes: %v", err)
	}

	err = http.ListenAndServe(host+":"+port, app.r)

	if errors.Is(err, http.ErrServerClosed) {
		app.LogInfo("Server closed")
	} else if err != nil {
		return err
	}

	return nil
}
