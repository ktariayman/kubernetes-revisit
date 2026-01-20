package main

import (
	"database/sql"
	"encoding/json"
	"log"
	"net/http"
	"time"

	_ "github.com/lib/pq"

	"github.com/ktariayman/kubernetes-revisit/internal/config"
	"github.com/ktariayman/kubernetes-revisit/internal/user"
)

func main() {
	cfg := config.Load()

	// 1. Database
	db, err := sql.Open("postgres", cfg.DatabaseURL)
	if err != nil {
		log.Fatalf("DB error: %v", err)
	}
	defer db.Close()
	
	if err := db.Ping(); err != nil {
		log.Fatalf("DB Ping failed: %v", err)
	}

	// 2. Migration
	if _, err := db.Exec(schema); err != nil {
		log.Fatalf("Schema error: %v", err)
	}

	// 3. Wiring
	repo := user.NewRepository(db)
	svc := user.NewService(repo)
	handler := user.NewHandler(svc)

	// 4. Routes
	http.HandleFunc("/health", healthHandler)
	http.HandleFunc("/users", handler.ListUsers)
	http.HandleFunc("/users/get", handler.GetUser)
	http.HandleFunc("/users/create", handler.CreateUser)

	log.Printf("Starting on %s", cfg.Port)
	http.ListenAndServe(":"+cfg.Port, nil)
}

func healthHandler(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(map[string]any{
		"status": "ok",
		"time":   time.Now(),
	})
}

var schema = `
CREATE TABLE IF NOT EXISTS users (
	id SERIAL PRIMARY KEY,
	name VARCHAR(100) NOT NULL,
	email VARCHAR(100) UNIQUE NOT NULL,
	created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);
`
