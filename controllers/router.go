package controllers

import "github.com/gin-gonic/gin"

func SetupRouter() *gin.Engine {
	router := gin.Default()

	NotebookController := NotebookController{}

	router.GET("/api/v1/notebooks", NotebookController.GetNotebooks)
	router.POST("/api/v1/notebooks", NotebookController.CreateNotebook)
	return router
}
