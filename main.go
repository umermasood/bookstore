package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"

	"bookstore/example/models"

	_ "github.com/lib/pq"
)

type Env struct {
	db *sql.DB
}

func main() {
	// Initialize the connection pool
	db, err := sql.Open("postgres", "postgresql://postgres:pTbZBLeksqDVrXotzEag@containers-us-west-33.railway.app:5751/railway")
	if err != nil {
		log.Fatal(err)
	}

	// Create an instance of Env containing the db connection pool
	env := &Env{db: db} // injected the dependency

	// Pass the Env struct as a parameter to booksIndex().
	http.HandleFunc("/books", booksIndex(env))
	log.Fatal(http.ListenAndServe(":3000", nil))
}

// Use a closure to make Env available to the handler logic.
func booksIndex(env *Env) http.HandlerFunc {
	return func(w http.ResponseWriter, _ *http.Request) {
		bks, err := models.AllBooks(env.db)
		if err != nil {
			log.Print(err)
			http.Error(w, http.StatusText(500), 500)
			return
		}

		for _, bk := range bks {
			_, _ = fmt.Fprintf(w, "%s, %s, %s, £%.2f\n", bk.Isbn, bk.Title, bk.Author, bk.Price)
		}
	}
}
