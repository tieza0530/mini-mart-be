package models

type Service struct {
	ID         string `gorm:"column:id"`
	CategoryId string `gorm:"column:category_id"`
	Name       string `gorm:"column:name"`
	Slug       string `gorm:"column:slug"`
}

func (Service) TableName() string {
	return "services"
}
