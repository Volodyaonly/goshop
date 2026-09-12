package db

import (
	"context"
	"goshop/internal/models"

	"github.com/jmoiron/sqlx"
)

type OrderRepo struct {
	db DBTX
}

func NewOrderRepo(db *sqlx.DB) *OrderRepo {
	return &OrderRepo{db: db}
}

func (r *OrderRepo) Create(
	ctx context.Context,
	o *models.Order,
) error {

	query := `
		INSERT INTO orders (
			product_id,
			quantity,
			total
		)
		VALUES ($1, $2, $3)
		RETURNING id
	`

	err := r.db.GetContext(
		ctx,
		&o.ID,
		query,
		o.ProductID,
		o.Quantity,
		o.Total,
	)

	return err
}
