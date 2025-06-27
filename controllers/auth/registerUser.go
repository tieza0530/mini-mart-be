package auth

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"mini-mart-db/services/auth"
	"net/http"
)

type UsersRegister struct {
	ID       string  `gorm:"column:id;primaryKey" json:"id"`
	Account  string  `gorm:"column:account"  json:"account,omitempty"`
	Email    string  `gorm:"column:email"    json:"email,omitempty"`
	Password string  `gorm:"column:password" json:"password,omitempty"`
	Role     *string `gorm:"column:role"     json:"role,omitempty"`
	Provider string  `gorm:"column:provider;default:local" json:"provider"`
}

func RegisterUser(DB *gorm.DB) func(c *gin.Context) {
	return func(c *gin.Context) {
		var input UsersRegister
		if err := c.ShouldBindJSON(&input); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		userData, refreshToken, err := auth.PostRegisterUser(DB, input.Account, input.Email, input.Password)
		if err != nil {
			c.JSON(409, gin.H{"error": err.Error()})
			return
		}
		c.SetCookie("refresh_token", refreshToken, 7*24*60*60, "/", "", false, true)
		c.JSON(http.StatusOK, gin.H{"data": userData})
	}
}
