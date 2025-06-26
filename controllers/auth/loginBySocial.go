package auth

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"mini-mart-db/services/auth"
	"net/http"
)

type CreateUserInput struct {
	Email      *string `json:"email"      binding:"omitempty,email"`
	ProviderID string  `json:"providerID" binding:"required"`
	Name       string  `json:"name"       binding:"required"`
	AvatarUrl  *string `json:"avatarUrl"  binding:"omitempty,url"`
	Token      *string `json:"token"      binding:"omitempty"`
	Provider   string  `json:"provider"   binding:"required,oneof=google facebook local"`
}

func LoginSocial(DB *gorm.DB) func(c *gin.Context) {
	return func(c *gin.Context) {
		var input CreateUserInput

		if err := c.ShouldBindJSON(&input); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		userData, err := auth.PostLoginUserBySocial(
			DB,
			input.Provider,
			input.Token,
			input.Email,
			&input.ProviderID,
			&input.Name,
			input.AvatarUrl,
		)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, gin.H{"data": userData})
	}
}
