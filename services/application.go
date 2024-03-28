package services

import (
	"applicationDesignTest/entities"
	"applicationDesignTest/repository"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"
	"os"
)

type App struct {
	Orders repository.OrdersRepository
	Rooms  repository.RoomsRepository
	Server *HTTPServer
	logger *log.Logger

	Registerable
}

func NewApp(
	Orders repository.OrdersRepository,
	Rooms repository.RoomsRepository,
	Server *HTTPServer,
	logger *log.Logger,
) *App {
	return &App{
		Orders: Orders,
		Rooms:  Rooms,
		Server: Server,
		logger: logger,
	}
}

func (app *App) CreateOrder(w http.ResponseWriter, r *http.Request) {
	var newOrder entities.Order
	json.NewDecoder(r.Body).Decode(&newOrder)

	if unavailableDays := app.Rooms.Book(newOrder.From, newOrder.To); len(unavailableDays) != 0 {
		http.Error(w, "Hotel room is not available for selected dates", http.StatusInternalServerError)
		app.LogErrorf("Hotel room is not available for selected dates:\n%v\n%v", newOrder, unavailableDays)
		return
	}

	if err := app.Orders.Store(&newOrder); err != nil {
		http.Error(w, "Failed to create order", http.StatusInternalServerError)
		app.LogErrorf("Failed to create order:\n%v\n%v", newOrder, err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(newOrder)

	app.LogInfo("Order successfully created: %v", newOrder)
}

func (app *App) GetRouteHandlers() map[string]http.HandlerFunc {
	return map[string]http.HandlerFunc{
		"/orders": app.CreateOrder,
	}
}

func (app *App) Run() error {
	app.LogInfo("Server listening on localhost:8080")
	err := app.Server.Serve()
	if errors.Is(err, http.ErrServerClosed) {
		app.LogInfo("Server closed")
	} else if err != nil {
		return err
	}

	return nil
}

func (app *App) LogErrorf(format string, v ...any) {
	msg := fmt.Sprintf(format, v...)
	app.logger.Printf("[Error]: %s\n", msg)
}

func (app *App) LogInfo(format string, v ...any) {
	msg := fmt.Sprintf(format, v...)
	app.logger.Printf("[Info]: %s\n", msg)
}

func (app *App) LogFatalf(format string, v ...any) {
	msg := fmt.Sprintf(format, v...)
	app.logger.Printf("[Error]: %s\n", msg)
	os.Exit(1)
}
