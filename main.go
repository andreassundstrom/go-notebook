package main

import (
	"andreassundstrom/go-notebook/controllers"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	router.GET("/api/v1/notebooks", controllers.GetNotebooks)
	router.POST("/api/v1/notebooks", controllers.CreateNotebook)

	router.Run("localhost:5000")
}
