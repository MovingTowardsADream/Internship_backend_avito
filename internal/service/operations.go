package service

import (
	"Internship_backend_avito/internal/repository/postgresdb"
	"context"
)

type OperationServices struct {
	repo postgresdb.Operation
}

func NewOperationServices(repo postgresdb.Operation) *OperationServices {
	return &OperationServices{repo: repo}
}

func (s *OperationServices) OperationsHistory(ctx context.Context, input OperationHistoryInput) ([]OperationHistoryOutput, error) {
	operations, productNames, err := s.repo.OperationsHistory(ctx, input.AccountId, input.SortType, input.Offset, input.Limit)
	if err != nil {
		return nil, err
	}

	output := make([]OperationHistoryOutput, 0, len(operations))
	for i, operation := range operations {

		output = append(output, OperationHistoryOutput{
			Amount:      operation.Amount,
			Operation:   operation.OperationType,
			Time:        operation.CreatedAt,
			Product:     productNames[i],
			Order:       operation.OrderId,
			Description: operation.Description,
		})
	}
	return output, nil
}
