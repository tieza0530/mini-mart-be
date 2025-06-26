package models

import "time"

type Users struct {
	ID                string     `gorm:"column:id"`
	Account           *string    `gorm:"column:account"`
	Password          *string    `gorm:"column:password" json:"-"`
	Provider          string     `gorm:"column:provider"`
	ProviderID        *string    `gorm:"column:provider_id"`
	Email             *string    `gorm:"column:email"`
	Phone             *string    `gorm:"column:phone"`
	Name              *string    `gorm:"column:name"`
	Address           *string    `gorm:"column:address"`
	AvatarURL         *string    `gorm:"column:avatar_url"`
	Role              *string    `gorm:"column:role" json:"-"`
	IsVerified        bool       `gorm:"column:is_verified"`
	IsVerifiedExpires *time.Time `gorm:"column:is_verified_expires"`
	Token             *string    `gorm:"column:token" json:"-"`
	CreatedAt         *time.Time `gorm:"column:created_at"`
}

func (Users) TableName() string {
	return "users"
}
