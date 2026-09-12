package events

type OrderCreated struct {
	EventID   string  `json:"event_id"`
	Event     string  `json:"event"`
	OrderID   int     `json:"order_id"`
	ProductID int     `json:"product_id"`
	Quantity  int     `json:"quantity"`
	Total     float64 `json:"total"`
}
