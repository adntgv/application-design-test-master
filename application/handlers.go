package application

import (
	"applicationDesignTest/entities"

	"encoding/json"
	"net/http"
)

type CreateOrderRequest struct {
	From string
	To   string
}

func (app *App) HandleCreateOrder(w http.ResponseWriter, r *http.Request) {
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
