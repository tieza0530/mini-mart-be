package category

import (
	"gorm.io/gorm"
	"mini-mart-db/models"
)

func GetCategoryBySlug(DB *gorm.DB, slug string) (*models.Category, error) {
	var category models.Category
	err := DB.Where("slug = ?", slug).First(&category).Error
	if err != nil {
		return nil, err
	}
	return &category, nil
}
