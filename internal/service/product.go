package service

import (
	"Internship_backend_avito/internal/entity"
	"Internship_backend_avito/internal/repository/postgresdb"
	"context"
)

type ProductServices struct {
	repo postgresdb.Product
}

func NewProductServices(repo postgresdb.Product) *ProductServices {
	return &ProductServices{repo: repo}
}

func (s *ProductServices) CreateProduct(ctx context.Context, name string) (int, error) {
	return s.repo.CreateProduct(ctx, name)
}

func (s *ProductServices) GetProductById(ctx context.Context, id int) (entity.Product, error) {
	return s.repo.GetProductById(ctx, id)
}
