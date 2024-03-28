package repository

import (
	"applicationDesignTest/entities"
	"time"
)

type RoomsRepository interface {
	Store(room *entities.Room) error
	GetAvailability() []entities.Room
	UpdateAvailability(room entities.Room) error
	Book(from, to time.Time) map[time.Time]struct{}
}

type DefaultRoomsRepository struct {
	rooms []entities.Room
}

// Book implements RoomsRepository.
func (d *DefaultRoomsRepository) Book(from, to time.Time) map[time.Time]struct{} {
	daysToBook := daysBetween(from, to)

	unavailableDays := make(map[time.Time]struct{})
	for _, day := range daysToBook {
		unavailableDays[day] = struct{}{}
	}

	for _, dayToBook := range daysToBook {
		for _, availability := range d.GetAvailability() {
			if !availability.Date.Equal(dayToBook) || availability.Quota < 1 {
				continue
			}
			availability.Quota -= 1
			d.UpdateAvailability(availability)
			delete(unavailableDays, dayToBook)
		}
	}

	return unavailableDays
}

// UpdateAvailability implements RoomsRepository.
func (d *DefaultRoomsRepository) UpdateAvailability(room entities.Room) error {
	for i, r := range d.rooms {
		if r.RoomID == room.RoomID {
			d.rooms[i] = room
		}
	}

	return nil
}

// GetAvailability implements RoomsRepository.
func (d *DefaultRoomsRepository) GetAvailability() []entities.Room {
	return []entities.Room{
		{HotelID: "reddison", RoomID: "lux", Date: date(2024, 1, 1), Quota: 1},
		{HotelID: "reddison", RoomID: "lux", Date: date(2024, 1, 2), Quota: 1},
		{HotelID: "reddison", RoomID: "lux", Date: date(2024, 1, 3), Quota: 1},
		{HotelID: "reddison", RoomID: "lux", Date: date(2024, 1, 4), Quota: 1},
		{HotelID: "reddison", RoomID: "lux", Date: date(2024, 1, 5), Quota: 0},
	}
}

// store implements RoomsRepository.
func (d *DefaultRoomsRepository) Store(room *entities.Room) error {
	d.rooms = append(d.rooms, *room)

	return nil
}

func NewRoomsRepository() (RoomsRepository, error) {
	return &DefaultRoomsRepository{
		rooms: make([]entities.Room, 0),
	}, nil
}

func date(year, month, day int) time.Time {
	return time.Date(year, time.Month(month), day, 0, 0, 0, 0, time.UTC)
}

func daysBetween(from time.Time, to time.Time) []time.Time {
	if from.After(to) {
		return nil
	}

	days := make([]time.Time, 0)
	for d := toDay(from); !d.After(toDay(to)); d = d.AddDate(0, 0, 1) {
		days = append(days, d)
	}

	return days
}

func toDay(timestamp time.Time) time.Time {
	return time.Date(timestamp.Year(), timestamp.Month(), timestamp.Day(), 0, 0, 0, 0, time.UTC)
}
