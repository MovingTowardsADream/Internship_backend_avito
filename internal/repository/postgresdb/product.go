package postgresdb

import (
	"Internship_backend_avito/internal/entity"
	"Internship_backend_avito/pkg/postgres"
	"context"
	"fmt"
)

type ProductPostgres struct {
	db *postgres.Postgres
}

func NewProductPostgres(db *postgres.Postgres) *ProductPostgres {
	return &ProductPostgres{db: db}
}

func (r *ProductPostgres) CreateProduct(ctx context.Context, name string) (int, error) {
	sql, args, _ := r.db.Builder.
		Insert("products").
		Columns("name").
		Values(name).
		Suffix("RETURNING id").
		ToSql()

	var id int
	err := r.db.Pool.QueryRow(ctx, sql, args...).Scan(&id)
	if err != nil {
		return 0, fmt.Errorf("ProductRepo.CreateProduct - r.Pool.QueryRow: %v", err)
	}

	return id, nil
}

func (r *ProductPostgres) GetProductById(ctx context.Context, id int) (entity.Product, error) {
	sql, args, _ := r.db.Builder.
		Select("*").
		From("products").
		Where("id = ?", id).
		ToSql()

	var product entity.Product
	err := r.db.Pool.QueryRow(ctx, sql, args...).Scan(
		&product.Id,
		&product.Name,
	)
	if err != nil {
		return entity.Product{}, fmt.Errorf("ProductRepo.GetProductById - r.Pool.QueryRow: %v", err)
	}

	return product, nil
}
