package inmem

import (
	"applicationDesignTest/internal/business/domains"
	"context"
	"fmt"
	"time"
)

type inmemRoomRepository struct {
	bookingDays map[string][]time.Time
}

// BookRoomForDaysBetween implements domains.RoomRepository.
func (i *inmemRoomRepository) BookRoomForDaysBetween(ctx context.Context, room *domains.RoomDomain, from time.Time, to time.Time) error {
	key := fmt.Sprintf("%v:%v", room.HotelID, room.RoomID)

	days := i.bookingDays[key]
	if days == nil {
		days = make([]time.Time, 0)
	}

	nextDay := from

	for nextDay.Before(to) {
		days = append(days, nextDay)
		nextDay = nextDay.Add(time.Hour * 24)
	}

	i.bookingDays[key] = days

	return nil
}

// GetRoomBookingDaysBetween implements domains.RoomRepository.
func (i *inmemRoomRepository) GetRoomBookingDaysBetween(ctx context.Context, room *domains.RoomDomain, from time.Time, to time.Time) []time.Time {
	key := fmt.Sprintf("%v:%v", room.HotelID, room.RoomID)
	bookedDays := i.bookingDays[key]

	result := make([]time.Time, 0)

	for _, date := range bookedDays {
		if date.After(from) && date.Before(to) {
			result = append(result, date)
		}
	}

	return result
}

func NewRoomRepository() domains.RoomRepository {
	return &inmemRoomRepository{
		bookingDays: make(map[string][]time.Time),
	}
}
