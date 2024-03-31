package inmem

import (
	"applicationDesignTest/internal/business/domains"
	"context"
	"fmt"
	"time"
)

type inmemRoomRepository struct {
	rooms       map[string]*domains.RoomDomain
	bookingDays map[string][]time.Time
}

// BookRoomForDaysBetween implements domains.RoomRepository.
func (i *inmemRoomRepository) BookRoomForDaysBetween(ctx context.Context, room *domains.RoomDomain, from time.Time, to time.Time) error {
	if err := i.exists(room); err != nil {
		return err
	}

	key := getRoomKey(room)

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
func (i *inmemRoomRepository) GetRoomBookingDaysBetween(ctx context.Context, room *domains.RoomDomain, from time.Time, to time.Time) ([]time.Time, error) {
	if err := i.exists(room); err != nil {
		return nil, err
	}

	key := getRoomKey(room)
	bookedDays := i.bookingDays[key]

	result := make([]time.Time, 0)

	for _, date := range bookedDays {
		if date.After(from) && date.Before(to) {
			result = append(result, date)
		}
	}

	return result, nil
}

func (i *inmemRoomRepository) exists(room *domains.RoomDomain) error {
	key := getRoomKey(room)

	if _, exists := i.rooms[key]; !exists {
		return fmt.Errorf("room with hotel id '%v' and room id '%v' does not exist", room.HotelID, room.RoomID)
	}

	return nil
}

func NewRoomRepository() domains.RoomRepository {
	rooms, bookingDays := initStorages()

	return &inmemRoomRepository{
		bookingDays: bookingDays,
		rooms:       rooms,
	}
}

func initStorages() (rooms map[string]*domains.RoomDomain, bookingDays map[string][]time.Time) {
	rooms = make(map[string]*domains.RoomDomain)
	bookingDays = make(map[string][]time.Time)

	// usually this is done in migrations

	for hotelId := 1; hotelId <= 5; hotelId++ {
		for roomId := 1; roomId <= 5; roomId++ {
			room := &domains.RoomDomain{
				HotelID: fmt.Sprint(hotelId),
				RoomID:  fmt.Sprint(roomId),
			}

			key := getRoomKey(room)
			rooms[key] = room
		}
	}

	for _, room := range rooms {
		key := getRoomKey(room)

		today := time.Now()

		from := today.Add(time.Hour * 24)
		to := from.Add(time.Hour * 24)

		bookingDays[key] = []time.Time{
			from, to,
		}
	}

	return rooms, bookingDays
}

func getRoomKey(room *domains.RoomDomain) string {
	return fmt.Sprintf("%v:%v", room.HotelID, room.RoomID)
}
