package postgresdb

import (
	"Internship_backend_avito/internal/entity"
	"Internship_backend_avito/pkg/postgres"
	"context"
	"fmt"
)

type OperationPostgres struct {
	db *postgres.Postgres
}

func NewOperationPostgres(db *postgres.Postgres) *OperationPostgres {
	return &OperationPostgres{db: db}
}

func (r *OperationPostgres) OperationsHistory(ctx context.Context, accountId int, sortType string, offset int, limit int) ([]entity.Operation, []string, error) {
	var orderBySql string
	switch sortType {
	case "":
		orderBySql = "created_at DESC"
	case "date":
		orderBySql = "created_at DESC"
	case "amount":
		orderBySql = "amount DESC"
	default:
		return nil, nil, fmt.Errorf("OperationRepo.PaginationOperations: unknown sort type - %s", sortType)
	}

	sqlQuery, args, _ := r.db.Builder.
		Select("operations.id", "account_id", "amount", "operation_type", "created_at", "COALESCE((case when operations.product_id is null then null else products.name end), '') as product_name", "order_id", "COALESCE(description, '')").
		From("operations").
		InnerJoin("products on operations.product_id = products.id or operations.product_id is null").
		Where("account_id = ?", accountId).
		OrderBy(orderBySql).
		Limit(uint64(limit)).
		Offset(uint64(offset)).
		ToSql()

	rows, err := r.db.Pool.Query(ctx, sqlQuery, args...)
	if err != nil {
		return nil, nil, fmt.Errorf("OperationRepo.paginationOperationsByDate - r.Pool.Query: %v", err)
	}
	defer rows.Close()

	var operations []entity.Operation
	var productNames []string
	for rows.Next() {
		var operation entity.Operation
		var productName string
		err = rows.Scan(&operation.Id, &operation.AccountId, &operation.Amount, &operation.OperationType, &operation.CreatedAt, &productName, &operation.OrderId, &operation.Description)
		if err != nil {
			return nil, nil, fmt.Errorf("OperationRepo.paginationOperationsByDate - rows.Scan: %v", err)
		}
		operations = append(operations, operation)
		productNames = append(productNames, productName)
	}

	return operations, productNames, nil
}
