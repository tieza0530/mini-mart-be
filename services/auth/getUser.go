package auth

import (
	"gorm.io/gorm"
	"mini-mart-db/models"
)

func GetUserALL(DB *gorm.DB, userID string) (models.Users, error) {
	var user models.Users
	err := DB.Where("id =?  ", userID).First(&user).Error
	if err != nil {
		return user, err
	}
	return user, nil
}
