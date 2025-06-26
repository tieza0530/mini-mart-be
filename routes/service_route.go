package routes

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"mini-mart-db/controllers/category"
)

func ServerRoute(r *gin.RouterGroup, DB *gorm.DB) {
	service := r.Group("/service")
	{
		service.GET("", category.GetServiceByCategoryID(DB))
	}
}
