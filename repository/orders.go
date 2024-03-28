package repository

import "applicationDesignTest/entities"

type OrdersRepository interface {
	Store(order *entities.Order) error
}

type DefaultOrdersRepository struct {
	orders []entities.Order
}

// store implements OrdersRepository.
func (d *DefaultOrdersRepository) Store(order *entities.Order) error {
	d.orders = append(d.orders, *order)

	return nil
}

func NewOrdersRepository() (OrdersRepository, error) {
	return &DefaultOrdersRepository{
		orders: make([]entities.Order, 0),
	}, nil
}
