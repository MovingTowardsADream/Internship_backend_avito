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
	CreateProduct(ctx context.Context, name string) (int, error)
	GetProductById(ctx context.Context, id int) (entity.Product, error)
}

type Reservation interface {
	CreateReservation(ctx context.Context, reservation entity.Reservation) (int, error)
}

type Operation interface {
	OperationsHistory(ctx context.Context, accountId int, sortType string, offset int, limit int) ([]entity.Operation, []string, error)
	OperationsFile(ctx context.Context, month, year int) ([]string, []int, error)
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
		Product:       NewProductPostgres(db),
		Operation:     NewOperationPostgres(db),
		Reservation:   NewReservationPostgres(db),
	}
}
