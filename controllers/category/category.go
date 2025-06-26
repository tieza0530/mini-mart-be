package category

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"mini-mart-db/services/category"
	"net/http"
)

func GetCategoryAll(DB *gorm.DB) func(c *gin.Context) {
	return func(c *gin.Context) {
		categorySlug := c.Query("slug")
		if categorySlug != "" {
			categoryBySlug, err := category.GetCategoryBySlug(DB, categorySlug)
			if err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"Error": err.Error()})
				return
			}
			c.JSON(http.StatusOK, gin.H{"data": categoryBySlug})
			return
		}
		categories, err := category.GetAllCategories(DB)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"Error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, gin.H{"data": categories})
	}
}
