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

func GetNotebooks() ([]models.Notebook, error) {
	var notebooks []models.Notebook
	rows, err := Db.Query(`
		SELECT
			"Id",
			"Name",
			"Author"
		FROM "Notebooks";
		`)
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

// Returns the created notebook
func CreateNotebook(newNotebook *models.Notebook) (int64, error) {
	var Id int64

	res := Db.QueryRow(`
	INSERT INTO "Notebooks" ("Id", "Name", "Author")
	VALUES (DEFAULT, $1, $2)
	RETURNING "Id";
	`, newNotebook.Name, newNotebook.Author)

	err := res.Scan(&Id)

	if err != nil {
		log.Fatalf("Error when getting last id: %v", err)
		return 0, err
	}

	return Id, nil
}
