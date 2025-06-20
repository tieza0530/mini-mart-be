package routes

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"mini-mart-db/controllers"
)

func CategoryRoute(r *gin.RouterGroup, DB *gorm.DB) {
	category := r.Group("/category")
	{
		category.GET("", controllers.GetCategory(DB))
	}
}
