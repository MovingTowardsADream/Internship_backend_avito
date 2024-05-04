package service

import (
	"Internship_backend_avito/internal/repository/postgresdb"
)

type Authorization interface {
}

type Account interface {
}

type Product interface {
}

type Reservation interface {
}

type Operation interface {
}

type Service struct {
	Authorization
	Account
	Product
	Reservation
	Operation
}

func NewService(repos *postgresdb.Repository) *Service {
	return &Service{}
}
