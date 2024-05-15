package service

import (
	"Internship_backend_avito/internal/repository/postgresdb"
	"bytes"
	"context"
	"encoding/csv"
	"errors"
	"strconv"
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

func (s *OperationServices) OperationsFile(ctx context.Context, month, year int) ([]byte, error) {
	products, amounts, err := s.repo.OperationsFile(ctx, month, year)
	if err != nil {
		return nil, errors.New("failed to get revenue operations")
	}

	// TODO del this
	products = []string{"Name", "LastName"}
	amounts = []int{3, 7}
	//

	b := bytes.Buffer{}
	w := csv.NewWriter(&b)

	for i := range products {
		err = w.Write([]string{products[i], strconv.Itoa(amounts[i])})
		if err != nil {
			return nil, errors.New("failed to write csv")
		}
	}

	w.Flush()
	if err = w.Error(); err != nil {
		return nil, errors.New("failed to write csv")
	}

	return b.Bytes(), nil
}
