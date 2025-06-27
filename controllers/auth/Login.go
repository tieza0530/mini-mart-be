package auth

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"log"
	"mini-mart-db/helper"
	"mini-mart-db/services/auth"
	"net/http"
	"os"
)

type LoginStruct struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

func Login(DB *gorm.DB) func(c *gin.Context) {
	return func(c *gin.Context) {
		secretKey := os.Getenv("SECRET_KEY")
		var input LoginStruct
		if err := c.ShouldBindJSON(&input); err != nil {
			c.JSON(400, gin.H{"error": err.Error()})
			return
		}
		user, err := auth.GetUserByAccount(DB, input.Username)
		if err != nil {
			c.JSON(404, gin.H{"error": err.Error()})
			return
		}

		if user.Password == nil || !helper.CheckPasswordHash(input.Password, *user.Password) {
			c.JSON(401, gin.H{"error": "Invalid credentials"})
			return
		}
		refreshToken, err := helper.GenerateJWT(secretKey, user.ID, 7*24*60)
		accessToken, err := helper.GenerateJWT(secretKey, user.ID, 15)

		if err != nil {
			log.Println("generate JWT:", err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Server error"})
			return
		}
		if err := DB.Model(&user).Update("token", refreshToken).Error; err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		c.SetCookie("refresh_token", refreshToken, 7*24*60*60, "/", "", false, true)
		c.JSON(http.StatusOK, gin.H{"data": user, "accessToken": accessToken})
	}
}
