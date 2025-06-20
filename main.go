package main

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"mini-mart-db/config"
	"mini-mart-db/routes"
	"time"
)

func main() {
	config.InitDB()

	r := gin.Default()

	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:3000"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	v1 := r.Group("/v1")
	routes.CategoryRoute(v1, config.DB)

	r.Run(":8080")
}
