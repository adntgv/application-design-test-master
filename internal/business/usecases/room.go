package usecases

import (
	"applicationDesignTest/internal/business/domains"
	"context"
	"fmt"
	"time"
)

type roomUsecase struct {
	repo domains.RoomRepository
}

// Book implements domains.RoomUsecase.
func (r *roomUsecase) Book(ctx context.Context, room *domains.RoomDomain, from time.Time, to time.Time) error {
	unavailableDates := r.repo.GetRoomBookingDaysBetween(ctx, room, from, to)
	if len(unavailableDates) > 0 {
		return fmt.Errorf("room \"%v\" cannot be booked, following dates are unavailable: %v", room, unavailableDates)
	}

	return r.repo.BookRoomForDaysBetween(ctx, room, from, to)
}

func NewRoomUsecase(repo domains.RoomRepository) domains.RoomUsecase {
	return &roomUsecase{
		repo: repo,
	}
}
