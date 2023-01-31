package models

import (
	"context"
	"database/sql"
	"errors"
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

func AllBooks(ctx context.Context) ([]Book, error) {
	// Retrieve the connection pool from the context. Because the
	// r.Context().Value() method always returns an interface{} type, we
	// need to type assert it into a *sql.DB before using it.
	db, ok := ctx.Value("db").(*sql.DB)
	if !ok {
		return nil, errors.New("could not get database connection pool from context")
	}

	rows, err := db.Query("SELECT * FROM books")
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
