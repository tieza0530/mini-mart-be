package auth

import (
	"errors"
	"github.com/google/uuid"
	"gorm.io/gorm"
	"mini-mart-db/models"
	"strings"
)

func PostLoginUserBySocial(DB *gorm.DB, provider string, token *string, email *string, providerID *string, name *string, avatarUrl *string) (models.Users, error) {
	var user models.Users
	err := DB.Where("email = ?", email).First(&user).Error
	if err == nil {
		return user, nil
	}
	if err != gorm.ErrRecordNotFound {
		return models.Users{}, err
	}

	if providerID == nil {
		return models.Users{}, errors.New("providerID is required")
	}
	defaultRole := `["user"]`
	cleanID := strings.ReplaceAll(uuid.NewString(), "-", "")
	user = models.Users{
		ID:         cleanID,
		Email:      email,
		ProviderID: providerID,
		Name:       name,
		AvatarURL:  avatarUrl,
		Token:      token,
		Provider:   provider,
		Role:       &defaultRole,
		IsVerified: true,
	}
	if err := DB.Create(&user).Error; err != nil {
		return models.Users{}, err
	}
	return user, nil
}
