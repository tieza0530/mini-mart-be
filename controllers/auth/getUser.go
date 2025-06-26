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

func GetUser(DB *gorm.DB) func(c *gin.Context) {
	return func(c *gin.Context) {
		secretKey := os.Getenv("SECRET_KEY")
		cookieValue, err := c.Cookie("refresh_token")
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Missing or invalid refresh token"})
			return
		}
		userID, err := helper.VerifyJWT(secretKey, cookieValue)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Missing or invalid refresh token"})
		}

		user, err := auth.GetUserALL(DB, userID)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		}
		accessToken, err := helper.GenerateJWT(secretKey, user.ID, 15)
		if err != nil {
			log.Fatal(err)
		}
		c.JSON(http.StatusOK, gin.H{"user": user, "accessToken": accessToken})
	}
}
