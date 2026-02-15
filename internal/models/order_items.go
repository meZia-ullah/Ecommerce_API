package models

type OrderItem struct {
	ID          uint    `gorm:"primaryKey"`
	OrderID     uint    `json:"order-id"`
	ProductName string  `json:"product-name"`
	Quantity    int     `json:"quantity"`
	Price       float64 `json:"price"`
}
