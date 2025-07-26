package main

import (
	"context"
	"log"
	"net/http"
	"os"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/jackc/pgx/v5"
	"github.com/psankar/sqlc-demo/sqlc/db"
)

var queries *db.Queries

func main() {
	pgURL := os.Getenv("POSTGRES_URL")
	migrationsDir := os.Getenv("MIGRATIONS_DIR")

	log.Println("POSTGRES_URL:", pgURL, "MIGRATIONS_DIR:", migrationsDir)

	conn, err := pgx.Connect(context.Background(), pgURL)
	if err != nil {
		log.Fatal("Database connection failure", err)
		return
	}

	m, err := migrate.New(migrationsDir, pgURL)
	if err != nil {
		log.Fatal("Error creating migrate instance", err)
		return
	}

	if err := m.Up(); err != nil && err != migrate.ErrNoChange {
		log.Fatal("Error applying migrations", err)
		return
	}

	queries = db.New(conn)

	http.HandleFunc("/add-post", addPostHandler)
	http.HandleFunc("/get-post/{post_id}", getPostHandler)

	err = http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal(err)
	}
}
