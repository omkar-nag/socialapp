package main

import (
	"log"

	"github.com/omkar-nag/socialapp/internal/db" // Uncomment if you need to use db package for database connection
	"github.com/omkar-nag/socialapp/internal/env"
	"github.com/omkar-nag/socialapp/internal/store"
)

func main() {
	addr := env.GetString("DB_ADDR", "postgres://postgres:postgres@localhost/socialnetwork?sslmode=disable")
	dbc, err := db.New(addr, 3, 3, "15m")
	if err != nil {
		log.Fatalf("Failed to connect to the database: %v", err)
	}

	defer dbc.Close()
	ststore := store.NewStorage(dbc) // Replace nil with your actual database connection
	// Call the seed function with the store
	db.Seed(ststore)
}
