package auth

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"mini-mart-db/helper"
	"mini-mart-db/services/auth"
	"net/http"
	"os"
)

type LogoutRequest struct {
	AccessToken string `json:"accessToken"`
}

func LogOut(DB *gorm.DB) func(c *gin.Context) {
	return func(c *gin.Context) {
		secretkey := os.Getenv("SECRET_KEY")
		var req LogoutRequest
		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		userID, err := helper.VerifyJWT(secretkey, req.AccessToken)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		user, err := auth.GetUserALL(DB, userID)
		if err != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": "Missing or invalid refresh token"})
			return
		}
		c.SetCookie("refresh_token", "", -1, "/", "", false, true)
		if err := DB.Model(&user).Update("token", "").Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Cannot update user token"})
			return
		}
		c.JSON(http.StatusOK, gin.H{"message": "Logout successful"})
	}
}
