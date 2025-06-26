package category

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"mini-mart-db/services/category"
	"net/http"
)

func GetServiceByCategoryID(DB *gorm.DB) func(c *gin.Context) {
	return func(c *gin.Context) {
		categoryID := c.Query("category-id")
		if categoryID == "" {
			c.JSON(http.StatusBadRequest, gin.H{"error": "category-id is required"})
			return
		}
		servicesData, err := category.GetServicesByCategory(DB, categoryID)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, gin.H{"data": servicesData})
	}
}
