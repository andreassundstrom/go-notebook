package main

import (
	"andreassundstrom/go-notebook/controllers"
	"andreassundstrom/go-notebook/database"
)

func main() {
	connectionString := "user=go-notebook dbname=go-notebook host=localhost sslmode=disable password=password"
	database.NotebookRepo = &database.NotebookRepository{}
	database.NotebookRepo.InitRepository(connectionString)

	router := controllers.SetupRouter()
	router.Run("localhost:5000")
}
