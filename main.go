package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"

	"bookstore/example/models"

	_ "github.com/lib/pq"
)

func main() {
	var err error

	// Initialize the sql.DB connection pool and assign it to the models.DB global variable.
	// temp railway postgres db
	models.DB, err = sql.Open("postgres", "postgresql://postgres:pTbZBLeksqDVrXotzEag@containers-us-west-33.railway.app:5751/railway")
	if err != nil {
		log.Fatal(err)
	}

	http.HandleFunc("/books", booksIndex)
	log.Fatal(http.ListenAndServe(":3000", nil))
}

// booksIndex sends an HTTP response listing all books.
func booksIndex(w http.ResponseWriter, r *http.Request) {
	bks, err := models.AllBooks()
	if err != nil {
		log.Print(err)
		http.Error(w, http.StatusText(500), 500)
		return
	}

	for _, bk := range bks {
		_, _ = fmt.Fprintf(w, "%s, %s, %s, £%.2f\n", bk.Isbn, bk.Title, bk.Author, bk.Price)
	}
}
