package service

import (
	"Internship_backend_avito/internal/entity"
	"Internship_backend_avito/internal/repository/postgresdb"
	"context"
	"fmt"
)

type ReservationServices struct {
	repo postgresdb.Reservation
}

func NewReservationServices(repo postgresdb.Reservation) *ReservationServices {
	return &ReservationServices{repo: repo}
}

func (s *ReservationServices) CreateReservation(ctx context.Context, input CreateReservationInput) (int, error) {
	reservation := entity.Reservation{
		AccountId: input.AccountId,
		ProductId: input.ProductId,
		OrderId:   input.OrderId,
		Amount:    input.Amount,
	}

	id, err := s.repo.CreateReservation(ctx, reservation)
	if err != nil {
		return 0, fmt.Errorf("Cant create reservation")
	}

	return id, nil
}
