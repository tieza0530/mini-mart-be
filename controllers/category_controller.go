package controllers

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"mini-mart-db/services"
	"net/http"
)

func GetCategory(DB *gorm.DB) func(c *gin.Context) {
	return func(c *gin.Context) {
		categories, err := services.GetAllCategories(DB)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"Error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, gin.H{"data": categories})
	}
}
