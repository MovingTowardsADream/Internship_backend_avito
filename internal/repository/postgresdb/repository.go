package postgresdb

import (
	"Internship_backend_avito/internal/entity"
	"github.com/jmoiron/sqlx"
)

type Authorization interface {
	CreateUser(user entity.User) (int, error)
	GetUser(username, password string) (entity.User, error)
}

type Account interface {
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

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Authorization: NewAuthPostgres(db),
	}
}
