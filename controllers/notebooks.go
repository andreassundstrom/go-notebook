package controllers

import (
	"andreassundstrom/go-notebook/database"

	"github.com/gin-gonic/gin"
)

func Notebooks(context *gin.Context) {
	notebooks, err := database.GetData()
	if err != nil {
		context.String(500, "Failed to get notebooks")
		return
	}

	context.JSON(200, notebooks)
	return
}
