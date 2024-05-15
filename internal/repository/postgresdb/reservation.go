package postgresdb

import (
	"Internship_backend_avito/internal/entity"
	"Internship_backend_avito/internal/repository/repository_errors"
	"Internship_backend_avito/pkg/postgres"
	"context"
	"fmt"
	"github.com/Masterminds/squirrel"
)

type ReservationPostgres struct {
	db *postgres.Postgres
}

func NewReservationPostgres(db *postgres.Postgres) *ReservationPostgres {
	return &ReservationPostgres{db: db}
}

func (r *ReservationPostgres) CreateReservation(ctx context.Context, reservation entity.Reservation) (int, error) {
	tx, err := r.db.Pool.Begin(ctx)
	if err != nil {
		return 0, fmt.Errorf("ReservationRepo.CreateReservation - r.Pool.Begin: %v", err)
	}
	defer func() { _ = tx.Rollback(ctx) }()

	sql, args, _ := r.db.Builder.
		Select("balance").
		From("accounts").
		Where("id = ?", reservation.AccountId).
		ToSql()

	var balance int
	err = tx.QueryRow(ctx, sql, args...).Scan(&balance)
	if err != nil {
		return 0, fmt.Errorf("ReservationRepo.CreateReservation - tx.QueryRow: %v", err)
	}

	if balance < reservation.Amount {
		return 0, repository_errors.ErrNotEnoughBalance
	}

	sql, args, _ = r.db.Builder.
		Update("accounts").
		Set("balance", squirrel.Expr("balance - ?", reservation.Amount)).
		Where("id = ?", reservation.AccountId).
		ToSql()

	_, err = tx.Exec(ctx, sql, args...)
	if err != nil {
		return 0, fmt.Errorf("ReservationRepo.CreateReservation - tx.Exec: %v", err)
	}

	sql, args, _ = r.db.Builder.
		Insert("reservations").
		Columns("account_id", "product_id", "order_id", "amount").
		Values(
			reservation.AccountId,
			reservation.ProductId,
			reservation.OrderId,
			reservation.Amount,
		).
		Suffix("RETURNING id").
		ToSql()

	var id int
	err = tx.QueryRow(ctx, sql, args...).Scan(&id)
	if err != nil {
		return 0, fmt.Errorf("ReservationRepo.CreateReservation - tx.QueryRow: %v", err)
	}

	sql, args, _ = r.db.Builder.
		Insert("operations").
		Columns("account_id", "amount", "operation_type", "product_id", "order_id").
		Values(
			reservation.AccountId,
			reservation.Amount,
			"reservation",
			reservation.ProductId,
			reservation.OrderId,
		).
		ToSql()

	_, err = tx.Exec(ctx, sql, args...)
	if err != nil {
		return 0, fmt.Errorf("ReservationRepo.CreateReservation - tx.Exec: %v", err)
	}

	err = tx.Commit(ctx)
	if err != nil {
		return 0, fmt.Errorf("ReservationRepo.CreateReservation - tx.Commit: %v", err)
	}

	return id, nil
}
