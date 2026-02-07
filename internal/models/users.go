package models

type User struct {
	ID    uint `gorm:"primaryKey" json:"id"`
	Name  string
	Email string `gorm:"unique"`

	Orders []Order `gorm:"foreignKey:UserID" json:"orders"`
}
