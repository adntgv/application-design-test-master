// Ниже реализован сервис бронирования номеров в отеле. В предметной области
// выделены два понятия: Order — заказ, который включает в себя даты бронирования
// и контакты пользователя, и RoomAvailability — количество свободных номеров на
// конкретный день.
//
// Задание:
// - провести рефакторинг кода с выделением слоев и абстракций
// - применить best-practices там где это имеет смысл
// - исправить имеющиеся в реализации логические и технические ошибки и неточности
package main

import (
	"applicationDesignTest/config"
	"applicationDesignTest/repository"
	"applicationDesignTest/services"
	"log"
)

func main() {
	logger := log.Default()

	cfg, err := config.New()
	if err != nil {
		logger.Fatalf("Could not load configuration: %v", err)
	}

	OrdersRepo, err := repository.NewOrdersRepository()
	if err != nil {
		logger.Fatalf("Could not init orders repository: %v", err)
	}

	RoomsRepo, err := repository.NewRoomsRepository()
	if err != nil {
		logger.Fatalf("Could not init rooms repository: %v", err)
	}

	httpServer := services.NewHTTPServer(cfg.Host, cfg.Port)

	app := services.NewApp(
		OrdersRepo,
		RoomsRepo,
		httpServer,
		logger,
	)

	if err := app.Run(); err != nil {
		app.LogFatalf("Server failed: %s", err)
	}
}
