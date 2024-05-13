package service

import (
	"Internship_backend_avito/internal/entity"
	"Internship_backend_avito/internal/repository/postgresdb"
	"context"
	"time"
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
	CreateProduct(ctx context.Context, name string) (int, error)
	GetProductById(ctx context.Context, id int) (entity.Product, error)
}

type Reservation interface {
}

type OperationHistoryInput struct {
	AccountId int    `json:"account_id" binding:"required"`
	SortType  string `json:"sort_type" binding:"required"`
	Offset    int    `json:"offset" binding:"required"`
	Limit     int    `json:"limit" binding:"required"`
}

type OperationHistoryOutput struct {
	Amount      int       `json:"amount"`
	Operation   string    `json:"operation"`
	Time        time.Time `json:"time"`
	Product     string    `json:"product,omitempty"`
	Order       *int      `json:"order,omitempty"`
	Description string    `json:"description,omitempty"`
}

type Operation interface {
	OperationsHistory(ctx context.Context, input OperationHistoryInput) ([]OperationHistoryOutput, error)
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
		Product:       NewProductServices(repos.Product),
		Operation:     NewOperationServices(repos.Operation),
	}
}
