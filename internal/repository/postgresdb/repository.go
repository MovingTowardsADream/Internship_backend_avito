package postgresdb

import (
	"Internship_backend_avito/internal/entity"
	"Internship_backend_avito/pkg/postgres"
	"context"
)

type Authorization interface {
	CreateUser(ctx context.Context, user entity.User) (int, error)
	GetUser(ctx context.Context, username, password string) (entity.User, error)
}

type Account interface {
	CreateAccount(ctx context.Context) (int, error)
	AccountDeposit(ctx context.Context, id, amount int) error
	Withdraw(ctx context.Context, id, amount int) error
	Transfer(ctx context.Context, id_from, id_to, amount int) error
	GetAccountById(ctx context.Context, accountId int) (entity.Account, error)
}

type Product interface {
}

type Reservation interface {
}

type Operation interface {
}

type Repository struct {
	Authorization
	Account
	Product
	Reservation
	Operation
}

func NewRepository(db *postgres.Postgres) *Repository {
	return &Repository{
		Authorization: NewAuthPostgres(db),
		Account:       NewAccountPostgres(db),
	}
}
