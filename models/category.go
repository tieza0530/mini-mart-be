package models

import "time"

type Category struct {
	ID          string     `gorm:"column:id"`
	Category    string     `gorm:"column:category"`
	Icon        string     `gorm:"column:icon"`
	Slug        string     `gorm:"column:slug"`
	Description *string    `gorm:"column:description"`
	CreatedAt   *time.Time `gorm:"column:created_at"`
}

func (Category) TableName() string {
	return "categories"
}
