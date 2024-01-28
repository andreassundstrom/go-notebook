package controllers

import (
	"andreassundstrom/go-notebook/database"
	"andreassundstrom/go-notebook/models"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetNotebooks(context *gin.Context) {
	notebooks, err := database.GetNotebooks()
	if err != nil {
		context.String(500, "Failed to get notebooks")
		return
	}

	context.JSON(200, notebooks)
	return
}

func CreateNotebook(context *gin.Context) {
	var newNotebook models.Notebook
	if err := context.Bind(&newNotebook); err != nil {
		log.Fatalf("Failed to parse json data: %v", err)
		context.String(http.StatusBadRequest, "Failed to parse json data")
		return
	}

	newId, err := database.CreateNotebook(&newNotebook)

	if err != nil {
		log.Fatalf("Failed to create new notebook: %v", err)
		context.String(http.StatusInternalServerError, "Failed to create new notebook")
		return
	}

	context.String(http.StatusCreated, "Created new notebook %d", newId)
	return
}
