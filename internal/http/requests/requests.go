package requests

import (
	"applicationDesignTest/internal/business/domains"
	"time"
)

type CreateOrderRequest struct {
	HotelID   string    `json:"hotel_id"`
	RoomID    string    `json:"room_id"`
	UserEmail string    `json:"email"`
	From      time.Time `json:"from"`
	To        time.Time `json:"to"`
}

func (r *CreateOrderRequest) ToRoom() *domains.RoomDomain {
	return &domains.RoomDomain{
		HotelID: r.HotelID,
		RoomID:  r.RoomID,
	}
}

func (r *CreateOrderRequest) ToOrder() *domains.OrderDomain {
	return &domains.OrderDomain{
		HotelID:   r.HotelID,
		RoomID:    r.RoomID,
		UserEmail: r.UserEmail,
		From:      r.From,
		To:        r.To,
	}
}
