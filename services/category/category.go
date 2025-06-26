package category

import (
	"gorm.io/gorm"
	"mini-mart-db/models"
)

func GetAllCategories(DB *gorm.DB) ([]models.Category, error) {
	var categories []models.Category
	err := DB.Find(&categories).Error
	return categories, err
}
