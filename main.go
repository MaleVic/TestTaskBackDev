package main

import (
	middleware "TestTaskBackDev/middleware"
	routes "TestTaskBackDev/routes"
	"github.com/gin-gonic/gin"
	_ "github.com/heroku/x/hmetrics/onload"
	"os"
)

func main() {
	port := os.Getenv("PORT")

	if port == "" {
		port = "8000"
	}

	router := gin.New()
	router.Use(gin.Logger())
	routes.UserRoutes(router)

	router.Use(middleware.Authentication())

	router.GET("/api-1", func(c *gin.Context) {

		c.JSON(200, gin.H{"success": "Доступ разрешён"})

	})

	router.Run(":" + port)
}
