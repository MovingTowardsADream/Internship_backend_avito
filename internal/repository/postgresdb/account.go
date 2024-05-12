package postgresdb

import (
	"Internship_backend_avito/internal/entity"
	"Internship_backend_avito/internal/repository/operations"
	"Internship_backend_avito/internal/repository/repository_errors"
	"Internship_backend_avito/pkg/postgres"
	"context"
	"errors"
	"fmt"
	"github.com/Masterminds/squirrel"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
	"github.com/sirupsen/logrus"
)

type AccountPostgres struct {
	db *postgres.Postgres
}

func NewAccountPostgres(db *postgres.Postgres) *AuthPostgres {
	return &AuthPostgres{db: db}
}

func (r *AuthPostgres) CreateAccount(ctx context.Context) (int, error) {
	sql, args, _ := r.db.Builder.
		Insert("accounts").
		Values(squirrel.Expr("DEFAULT")).
		Suffix("RETURNING id").
		ToSql()

	var id int
	err := r.db.Pool.QueryRow(ctx, sql, args...).Scan(&id)
	if err != nil {
		logrus.Debugf("err: %v", err)
		var pgErr *pgconn.PgError
		if ok := errors.As(err, &pgErr); ok {
			if pgErr.Code == "23505" {
				return 0, repository_errors.ErrAlreadyExists
			}
		}
		return 0, fmt.Errorf("AccountRepo.CreateAccount - r.Pool.QueryRow: %v", err)
	}

	return id, nil
}

func (r *AuthPostgres) AccountDeposit(ctx context.Context, id, amount int) error {
	tx, err := r.db.Pool.Begin(ctx)
	if err != nil {
		return fmt.Errorf("AccountRepo.Deposit - r.Pool.Begin: %v", err)
	}
	defer func() { _ = tx.Rollback(ctx) }()

	sql, args, _ := r.db.Builder.
		Update("accounts").
		Set("balance", squirrel.Expr("balance + ?", amount)).
		Where("id = ?", id).
		ToSql()

	_, err = tx.Exec(ctx, sql, args...)
	if err != nil {
		return fmt.Errorf("AccountRepo.Deposit - tx.Exec: %v", err)
	}

	sql, args, _ = r.db.Builder.
		Insert("operations").
		Columns("account_id", "amount", "operation_type").
		Values(id, amount, operations.OperationTypeDeposit).
		ToSql()

	_, err = tx.Exec(ctx, sql, args...)
	if err != nil {
		return fmt.Errorf("AccountRepo.Deposit - tx.Exec: %v", err)
	}

	err = tx.Commit(ctx)
	if err != nil {
		return fmt.Errorf("AccountRepo.Deposit - tx.Commit: %v", err)
	}

	return nil
}

func (r *AuthPostgres) Withdraw(ctx context.Context, id, amount int) error {
	tx, err := r.db.Pool.Begin(ctx)
	if err != nil {
		return fmt.Errorf("AccountRepo.Withdraw - r.Pool.Begin: %v", err)
	}
	defer func() { _ = tx.Rollback(ctx) }()

	sql, args, _ := r.db.Builder.
		Select("balance").
		From("accounts").
		Where("id = ?", id).
		ToSql()

	var balance int
	err = tx.QueryRow(ctx, sql, args...).Scan(&balance)
	if err != nil {
		return fmt.Errorf("AccountRepo.Withdraw - tx.QueryRow: %v", err)
	}

	if balance < amount {
		return repository_errors.ErrNotEnoughBalance
	}

	sql, args, _ = r.db.Builder.
		Update("accounts").
		Set("balance", squirrel.Expr("balance - ?", amount)).
		Where("id = ?", id).
		ToSql()

	_, err = tx.Exec(ctx, sql, args...)
	if err != nil {
		return fmt.Errorf("AccountRepo.Withdraw - tx.Exec: %v", err)
	}

	sql, args, _ = r.db.Builder.
		Insert("operations").
		Columns("account_id", "amount", "operation_type").
		Values(id, amount, operations.OperationTypeWithdraw).
		ToSql()

	_, err = tx.Exec(ctx, sql, args...)
	if err != nil {
		return fmt.Errorf("AccountRepo.Withdraw - tx.Exec: %v", err)
	}

	err = tx.Commit(ctx)
	if err != nil {
		return fmt.Errorf("AccountRepo.Withdraw - tx.Commit: %v", err)
	}

	return nil
}

func (r *AuthPostgres) Transfer(ctx context.Context, id_from, id_to, amount int) error {
	tx, err := r.db.Pool.Begin(ctx)
	if err != nil {
		return fmt.Errorf("AccountRepo.Transfer - r.Pool.Begin: %v", err)
	}
	defer func() { _ = tx.Rollback(ctx) }()

	sql, args, _ := r.db.Builder.
		Select("balance").
		From("accounts").
		Where("id = ?", id_from).
		ToSql()

	var balance int
	err = tx.QueryRow(ctx, sql, args...).Scan(&balance)
	if err != nil {
		return fmt.Errorf("AccountRepo.Transfer - tx.QueryRow: %v", err)
	}

	if balance < amount {
		return repository_errors.ErrNotEnoughBalance
	}

	sql, args, _ = r.db.Builder.
		Update("accounts").
		Set("balance", squirrel.Expr("balance - ?", amount)).
		Where("id = ?", id_from).
		ToSql()

	_, err = tx.Exec(ctx, sql, args...)
	if err != nil {
		return fmt.Errorf("AccountRepo.Transfer - tx.Exec: %v", err)
	}

	sql, args, _ = r.db.Builder.
		Update("accounts").
		Set("balance", squirrel.Expr("balance + ?", amount)).
		Where("id = ?", id_to).
		ToSql()

	_, err = tx.Exec(ctx, sql, args...)
	if err != nil {
		return fmt.Errorf("AccountRepo.Transfer - tx.Exec: %v", err)
	}

	sql, args, _ = r.db.Builder.
		Insert("operations").
		Columns("account_id", "amount", "operation_type").
		Values(id_from, amount, operations.OperationTypeTransferFrom).
		ToSql()

	_, err = tx.Exec(ctx, sql, args...)
	if err != nil {
		return fmt.Errorf("AccountRepo.Transfer - tx.Exec: %v", err)
	}

	sql, args, _ = r.db.Builder.
		Insert("operations").
		Columns("account_id", "amount", "operation_type").
		Values(id_to, amount, operations.OperationTypeTransferTo).
		ToSql()

	_, err = tx.Exec(ctx, sql, args...)
	if err != nil {
		return fmt.Errorf("AccountRepo.Transfer - tx.Exec: %v", err)
	}

	err = tx.Commit(ctx)
	if err != nil {
		return fmt.Errorf("AccountRepo.Transfer - tx.Commit: %v", err)
	}

	return nil
}

func (r *AuthPostgres) GetAccountById(ctx context.Context, accountId int) (entity.Account, error) {
	sql, args, _ := r.db.Builder.
		Select("*").
		From("accounts").
		Where("id = ?", accountId).
		ToSql()

	var account entity.Account
	err := r.db.Pool.QueryRow(ctx, sql, args...).Scan(
		&account.Id,
		&account.Balance,
		&account.CreatedAt,
	)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return entity.Account{}, repository_errors.ErrNotFound
		}
		return entity.Account{}, fmt.Errorf("AccountRepo.GetAccountById - r.Pool.QueryRow: %v", err)
	}

	return account, nil
}
