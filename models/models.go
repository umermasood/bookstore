package models

import (
	"database/sql"
	"log"
)

// Create an exported global variable to hold the database connection pool

var DB *sql.DB

type Book struct {
	Isbn   string
	Title  string
	Author string
	Price  float32
}

// AllBooks returns a slice of all books in the books table.
func AllBooks() ([]Book, error) {
	// Note that we are calling Query() on the global variable.
	rows, err := DB.Query("SELECT * FROM books")
	if err != nil {
		return nil, err
	}
	defer func(rows *sql.Rows) {
		err := rows.Close()
		if err != nil {
			log.Println("Error while defer closing rows")
		}
	}(rows)

	var bks []Book

	for rows.Next() {
		var bk Book

		err := rows.Scan(&bk.Isbn, &bk.Title, &bk.Author, &bk.Price)
		if err != nil {
			return nil, err
		}

		bks = append(bks, bk)
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}

	return bks, nil
}
