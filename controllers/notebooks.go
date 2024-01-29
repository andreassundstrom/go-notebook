package controllers

import (
	"andreassundstrom/go-notebook/database"
	"andreassundstrom/go-notebook/models"
	"database/sql"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

type NotebookController struct {
	Database *sql.DB
}

func (c *NotebookController) GetNotebooks(context *gin.Context) {
	notebooks, err := database.NotebookRepo.GetNotebooks()
	if err != nil {
		context.String(500, "Failed to get notebooks")
		return
	}

	context.JSON(200, notebooks)
	return
}

func (c *NotebookController) CreateNotebook(context *gin.Context) {
	var newNotebook models.Notebook
	if err := context.Bind(&newNotebook); err != nil {
		log.Fatalf("Failed to parse json data: %v", err)
		context.String(http.StatusBadRequest, "Failed to parse json data")
		return
	}

	newId, err := database.NotebookRepo.CreateNotebook(&newNotebook)

	if err != nil {
		log.Fatalf("Failed to create new notebook: %v", err)
		context.String(http.StatusInternalServerError, "Failed to create new notebook")
		return
	}

	context.String(http.StatusCreated, "Created new notebook %d", newId)
	return
}
