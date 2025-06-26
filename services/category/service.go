package category

import (
	"gorm.io/gorm"
	"mini-mart-db/models"
)

func GetServicesByCategory(DB *gorm.DB, categoryID string) ([]models.Service, error) {
	var services []models.Service
	err := DB.Where("category_id = ?", categoryID).Find(&services).Error
	if err != nil {
		return nil, err
	}
	return services, nil
}
