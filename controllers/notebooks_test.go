package controllers

import (
	"andreassundstrom/go-notebook/database"
	"andreassundstrom/go-notebook/models"
	"andreassundstrom/go-notebook/test_helpers"
	"context"
	"encoding/json"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetNotebooks(t *testing.T) {
	// Setup
	ctx := context.Background()
	testContainer, err := test_helpers.CreatePostgresContainer(ctx)

	if err != nil {
		log.Fatal(err)
	}

	database.NotebookRepo = &database.NotebookRepository{}
	database.NotebookRepo.InitRepository(testContainer.ConnectionString)

	expectedNotebook := models.Notebook{
		Name:   "My notes",
		Author: "Andreas",
	}

	database.NotebookRepo.CreateNotebook(&expectedNotebook)

	router := SetupRouter()
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/api/v1/notebooks", nil)

	// Act
	router.ServeHTTP(w, req)

	// Assert
	assert.Equal(t, 200, w.Code)

	var actualNotebooks []models.Notebook

	if err := json.Unmarshal(w.Body.Bytes(), &actualNotebooks); err != nil {
		assert.Fail(t, "Failed to parse response")
	}

	assert.Len(t, actualNotebooks, 1)
	assert.Equal(t, expectedNotebook.Author, actualNotebooks[0].Author)
	assert.Equal(t, expectedNotebook.Name, actualNotebooks[0].Name)
}
