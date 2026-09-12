package db

import (
	"context"
	"database/sql"
	"goshop/internal/models"

	"github.com/jmoiron/sqlx"
)

type ProductRepo struct {
	db DBTX
}

type DBTX interface {
	GetContext(ctx context.Context, dest any, query string, args ...any) error
	ExecContext(ctx context.Context, query string, args ...any) (sql.Result, error)
}

func NewProductRepo(db *sqlx.DB) *ProductRepo {
	return &ProductRepo{db: db}
}

func (r *ProductRepo) GetByID(ctx context.Context, id int) (*models.Product, error) {
	var p models.Product
	err := r.db.GetContext(ctx, &p, "SELECT * FROM products WHERE id=$1", id)
	return &p, err
}
