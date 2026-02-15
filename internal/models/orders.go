package models

type Order struct {
	ID        uint        `gorm:"primaryKey" json:"id"`
	UserID    uint        `json:"user_id"`
	User      User        `json:"-"`
	Status    string      `json:"status"`
	OrderItem []OrderItem `gorm:"foreignKey:OrderID"`
}
