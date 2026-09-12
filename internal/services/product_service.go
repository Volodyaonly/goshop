package services

import (
	"context"
	"fmt"
	"goshop/internal/cache"
	"goshop/internal/db"
	"goshop/internal/models"
)

type ProductService struct {
	repo  *db.ProductRepo
	cache *cache.ProductCache
}

func NewProductService(
	repo *db.ProductRepo,
	cache *cache.ProductCache,
) *ProductService {

	return &ProductService{
		repo:  repo,
		cache: cache,
	}
}

func (s *ProductService) GetProduct(
	ctx context.Context,
	id int,
) (*models.Product, error) {

	product, err := s.cache.Get(ctx, id)

	if err == nil && product != nil {
		fmt.Println("Из Redis")
		return product, nil
	}

	fmt.Println("Из PostgreSQL")

	product, err = s.repo.GetByID(ctx, id)
	if err != nil {
		return nil, err
	}

	if err := s.cache.Set(ctx, product); err != nil {
		fmt.Println("не удалось сохранить товар в Redis:", err)
	}

	return product, nil
}
