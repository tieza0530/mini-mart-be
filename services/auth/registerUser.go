package auth

import (
	"errors"
	"github.com/google/uuid"
	"gorm.io/gorm"
	"log"
	"mini-mart-db/helper"
	"mini-mart-db/models"
	"os"
	"strings"
)

func PostRegisterUser(DB *gorm.DB, account string, email string, password string) (models.Users, string, error) {
	var user models.Users
	secretKey := os.Getenv("SECRET_KEY")
	if account == "" || email == "" || password == "" {
		return user, "", errors.New("account or email or password is empty")
	}
	err := DB.Where("email = ? OR account = ?", email, account).First(&user).Error
	if err == nil {
		return user, "", errors.New("email or account already exists")
	}
	if !errors.Is(err, gorm.ErrRecordNotFound) {
		return user, "", err
	}
	defaultRole := `["user"]`
	cleanID := strings.ReplaceAll(uuid.NewString(), "-", "")

	hash, err := helper.HashPassword(password)
	if err != nil {
		log.Fatal(err)
	}

	refreshToken, err := helper.GenerateJWT(secretKey, cleanID, (7 * 24 * 60))
	if err != nil {
		log.Fatal(err)
	}

	user = models.Users{
		ID:       cleanID,
		Account:  &account,
		Email:    &email,
		Password: &hash,
		Role:     &defaultRole,
		Provider: "local",
		Token:    &refreshToken,
	}

	if err := DB.Create(&user).Error; err != nil {
		return user, "", err
	}
	return user, refreshToken, nil
}
