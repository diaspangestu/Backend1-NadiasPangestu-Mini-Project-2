package entities

import "time"

type Customer struct {
	ID        uint   `gorm:"primary_key" `
	FirstName string `gorm:"column:first_name" json:"first_name"`
	LastName  string `gorm:"column:last_name" json:"last_name"`
	Email     string `gorm:"column:email"`
	Avatar    string `gorm:"avatar"`
	CreatedAt time.Time
	UpdatedAt time.Time
}
