package main

import (
	"andreassundstrom/go-notebook/controllers"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	router.GET("/notebooks", controllers.Notebooks)

	router.Run("localhost:5000")
}
