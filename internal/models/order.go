package models

type Order struct {
	ID        int     `db:"id" json:"id"`
	ProductID int     `db:"product_id" json:"product_id"`
	Quantity  int     `db:"quantity" json:"quantity"`
	Total     float64 `db:"total" json:"total"`
}
