package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"

	"bookstore/example/models"

	_ "github.com/lib/pq"
)

// This time make the models.BookModel a dependency in Env

type Env struct {
	books models.BookModel
}

func main() {
	// Initialize the connection pool
	db, err := sql.Open("postgres", "postgresql://postgres:pTbZBLeksqDVrXotzEag@containers-us-west-33.railway.app:5751/railway")
	if err != nil {
		log.Fatal(err)
	}

	// Initialize Env with a models.BookModel instance (which in turn wraps the connection pool).
	env := &Env{
		books: models.BookModel{DB: db},
	}

	http.HandleFunc("/books", env.booksIndex)
	log.Fatal(http.ListenAndServe(":3000", nil))
}

func (env *Env) booksIndex(w http.ResponseWriter, _ *http.Request) {
	// Execute the SQL query by calling the AllBooks() method.
	bks, err := env.books.AllBooks()
	if err != nil {
		log.Print(err)
		http.Error(w, http.StatusText(500), 500)
		return
	}

	for _, bk := range bks {
		_, _ = fmt.Fprintf(w, "%s, %s, %s, £%.2f\n", bk.Isbn, bk.Title, bk.Author, bk.Price)
	}
}
