package service

import (
	"Internship_backend_avito/internal/repository/postgresdb"
)

type AuthCreateUserInput struct {
	Username string
	Password string
}

type AuthGenerateTokenInput struct {
	Username string
	Password string
}

type Authorization interface {
	CreateUser(input AuthCreateUserInput) (int, error)
	GenerateToken(input AuthGenerateTokenInput) (string, error)
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
	return &Service{
		Authorization: NewAuthServices(repos.Authorization),
	}
}
