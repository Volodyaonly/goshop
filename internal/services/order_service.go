package services

import (
	"context"

	"goshop/internal/db"
	appkafka "goshop/internal/kafka"
	"goshop/internal/models"
)

type OrderService struct {
	repo     *db.OrderRepo
	producer *appkafka.Producer
}

func NewOrderService(
	repo *db.OrderRepo,
	producer *appkafka.Producer,
) *OrderService {

	return &OrderService{
		repo:     repo,
		producer: producer,
	}
}

func (s *OrderService) CreateOrder(
	ctx context.Context,
	o *models.Order,
) error {

	if err := s.repo.Create(ctx, o); err != nil {
		return err
	}

	if err := s.producer.PublishOrderCreated(ctx, o); err != nil {
		return err
	}

	return nil
}
