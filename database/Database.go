package database

import (
	"andreassundstrom/go-notebook/models"
	"database/sql"
	"log"

	_ "github.com/lib/pq"
)

var Db *sql.DB

func init() {
	connectionString := "user=go-notebook dbname=go-notebook host=localhost sslmode=disable password=password"
	db, err := sql.Open("postgres", connectionString)
	if err != nil {
		log.Fatal(err)
	}
	Db = db
}

func GetData() ([]models.Notebook, error) {
	var notebooks []models.Notebook
	rows, err := Db.Query("SELECT 1 as \"ID\", 'My notes' AS \"Name\", 'Author' AS \"Author\" ")
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	for rows.Next() {
		var notebook models.Notebook
		if err := rows.Scan(&notebook.Id, &notebook.Author, &notebook.Name); err != nil {
			return nil, err
		}

		notebooks = append(notebooks, notebook)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return notebooks, nil
}
