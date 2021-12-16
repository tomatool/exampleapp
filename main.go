package main

import (
	"encoding/json"
	"log"
	"net/http"
	"os"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

func main() {
	db := sqlx.MustOpen("postgres", os.Getenv("DATABASE_DSN"))

	mux := http.NewServeMux()

	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		json.NewEncoder(w).Encode(map[string]interface{}{
			"name":    "mainapp",
			"version": "v0.1.0",
		})
	})

	mux.HandleFunc("/users", func(w http.ResponseWriter, r *http.Request) {
		var users []struct {
			ID    string `db:"id"`
			Email string `db:"email"`
		}

		if err := db.Select(&users, "SELECT id, email FROM users"); err != nil {
			log.Println(err)
			json.NewEncoder(w).Encode(map[string]interface{}{
				"error": err.Error(),
			})
			return
		}

		log.Println(users)
		json.NewEncoder(w).Encode(users)
	})

	srv := &http.Server{
		Addr:    "0.0.0.0:8080",
		Handler: mux,
	}

	if err := srv.ListenAndServe(); err != nil {
		log.Fatal(err)
	}
}
