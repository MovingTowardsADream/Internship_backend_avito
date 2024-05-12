package service

import (
	"Internship_backend_avito/internal/entity"
	"Internship_backend_avito/internal/repository/postgresdb"
	"context"
)

type AccountService struct {
	repo postgresdb.Account
}

func NewAccountService(repo postgresdb.Account) *AccountService {
	return &AccountService{repo: repo}
}

func (s *AccountService) CreateAccount(ctx context.Context) (int, error) {
	return s.repo.CreateAccount(ctx)
}

func (s *AccountService) AccountDeposit(ctx context.Context, input AccountDepositInput) error {
	return s.repo.AccountDeposit(ctx, input.Id, input.Amount)
}

func (s *AccountService) Withdraw(ctx context.Context, input AccountWithdrawInput) error {
	return s.repo.Withdraw(ctx, input.Id, input.Amount)
}

func (s *AccountService) Transfer(ctx context.Context, input AccountTransferInput) error {
	return s.repo.Transfer(ctx, input.IdFrom, input.IdTo, input.Amount)
}

func (s *AccountService) GetAccountById(ctx context.Context, accountId int) (entity.Account, error) {
	return s.repo.GetAccountById(ctx, accountId)
}
