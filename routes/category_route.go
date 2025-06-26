package routes

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"mini-mart-db/controllers/category"
)

func CategoryRoute(r *gin.RouterGroup, DB *gorm.DB) {
	categories := r.Group("/category")
	{
		categories.GET("", category.GetCategoryAll(DB))
	}
}
