package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"

	_ "github.com/lib/pq" // PostgreSQL driver
)

func main() {
	// 1. Get the Port from DigitalOcean (Defaults to 8080)
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	// 2. Get the Database URL provided by DigitalOcean
	dbURL := os.Getenv("DATABASE_URL")

	// 3. Simple Health Check Handler
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "<h1>WMS1 is Live!</h1>")
		
		if dbURL != "" {
			fmt.Fprintf(w, "<p>Database connection string detected.</p>")
		} else {
			fmt.Fprintf(w, "<p>Waiting for database connection...</p>")
		}
	})

	// 4. Database Test Handler
	http.HandleFunc("/db-test", func(w http.ResponseWriter, r *http.Request) {
		db, err := sql.Open("postgres", dbURL)
		if err != nil {
			http.Error(w, "Failed to open database", 500)
			return
		}
		defer db.Close()

		err = db.Ping()
		if err != nil {
			fmt.Fprintf(w, "Could not reach database: %v", err)
		} else {
			fmt.Fprintf(w, "Successfully connected to PostgreSQL!")
		}
	})

	fmt.Printf("Server starting on port %s...\n", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
