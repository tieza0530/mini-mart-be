package auth

import (
	"gorm.io/gorm"
	"mini-mart-db/models"
)

func GetUserByAccount(DB *gorm.DB, account string) (models.Users, error) {
	var user models.Users
	err := DB.Where("account =?  ", account).First(&user).Error
	if err != nil {
		return user, err
	}
	return user, nil
}
