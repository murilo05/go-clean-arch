package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/joho/godotenv"
)

func main() {
	_ = godotenv.Load(".env")

	dbUser := os.Getenv("DB_USER")
	dbPass := os.Getenv("DB_PASSWORD")
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbName := os.Getenv("DB_NAME")

	if dbUser == "" || dbPass == "" || dbHost == "" || dbPort == "" || dbName == "" {
		log.Fatal("Missing database environment variables. Check your .env file.")
	}

	databaseURL := fmt.Sprintf(
		"postgres://%s:%s@%s:%s/%s?sslmode=disable",
		dbUser, dbPass, dbHost, dbPort, dbName,
	)

	sqlDir := os.Getenv("POSTGRES_SQL_DIR")
	if sqlDir == "" {
		log.Fatal("POSTGRES_SQL_DIR not set in .env")
	}

	migrationsPath := "file://" + sqlDir

	action := flag.String("action", "up", "Migration action: up or down")
	flag.Parse()

	m, err := migrate.New(migrationsPath, databaseURL)
	if err != nil {
		log.Fatalf("Error loading migrations: %v", err)
	}

	switch *action {
	case "up":
		err = m.Up()
	case "down":
		err = m.Down()
	default:
		log.Fatalf("Invalid action '%s'. Use: up or down", *action)
	}

	if err != nil {
		if err == migrate.ErrNoChange {
			fmt.Println("No migrations to apply.")
			return
		}
		log.Fatalf("Migration error: %v", err)
	}

	fmt.Println("Migration completed successfully!")
}
