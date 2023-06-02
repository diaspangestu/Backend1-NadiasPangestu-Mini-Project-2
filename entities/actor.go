package entities

import "time"

type Status string

const (
	True  Status = "true"
	False Status = "false"
)

type Actor struct {
	ID         uint   `gorm:"primary_key"`
	Username   string `gorm:"column:username;unique"`
	Password   string `gorm:"column:password"`
	RoleID     int    `gorm:"column:role_id"`
	IsVerified Status `gorm:"column:is_verified;type enum('true','false')"`
	IsActive   Status `gorm:"column:is_active;type enum('true','false')"`
	CreatedAt  time.Time
	UpdatedAt  time.Time
}
