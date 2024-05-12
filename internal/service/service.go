package service

import (
	"Internship_backend_avito/internal/entity"
	"Internship_backend_avito/internal/repository/postgresdb"
	"context"
)

type AuthCreateUserInput struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type AuthGenerateTokenInput struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type Authorization interface {
	CreateUser(ctx context.Context, input AuthCreateUserInput) (int, error)
	GenerateToken(ctx context.Context, input AuthGenerateTokenInput) (string, error)
	ParseToken(token string) (int, error)
}

type AccountDepositInput struct {
	Id     int `json:"id" binding:"required"`
	Amount int `json:"amount" binding:"required"`
}

type AccountWithdrawInput struct {
	Id     int `json:"id" binding:"required"`
	Amount int `json:"amount" binding:"required"`
}

type AccountTransferInput struct {
	IdFrom int `json:"id_from" binding:"required"`
	IdTo   int `json:"id_to" binding:"required"`
	Amount int `json:"amount" binding:"required"`
}

type Account interface {
	CreateAccount(ctx context.Context) (int, error)
	AccountDeposit(ctx context.Context, input AccountDepositInput) error
	Withdraw(ctx context.Context, input AccountWithdrawInput) error
	Transfer(ctx context.Context, input AccountTransferInput) error
	GetAccountById(ctx context.Context, accountId int) (entity.Account, error)
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
		Account:       NewAccountService(repos.Account),
	}
}
