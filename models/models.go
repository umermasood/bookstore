package models

import (
	"database/sql"
	"log"
)

type Book struct {
	Isbn   string
	Title  string
	Author string
	Price  float32
}

// Create a custom BookModel type which wraps the sql.DB connection pool.

type BookModel struct {
	// wraps the db dependency
	DB *sql.DB
}

// Use a method on the custom BookModel type to run the SQL query.

func (m *BookModel) AllBooks() ([]Book, error) {
	rows, err := m.DB.Query("SELECT * FROM books")
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
