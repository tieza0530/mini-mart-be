package routes

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"mini-mart-db/controllers/auth"
)

func AuthRoute(r *gin.RouterGroup, DB *gorm.DB) {
	users := r.Group("/auth")
	{
		users.POST("/", auth.LoginSocial(DB))
		users.POST("/register", auth.RegisterUser(DB))
		users.GET("/refresh-token", auth.GetUser(DB))
		users.POST("/logout", auth.LogOut(DB))
		users.POST("/login", auth.Login(DB))
	}
}
