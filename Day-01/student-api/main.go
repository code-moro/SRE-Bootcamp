package main

import (
	"log"
	"os"
	"student-api/routers"
	"student-api/database"
	"github.com/golang-migrate/migrate/v4"

	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	
)

func main() {
	connString := "postgres://postgres:password@localhost:5432/studentdb?sslmode=disable"

	if err := database.Connect(connString); err != nil {
		log.Fatal(err)
	}
	defer database.Close()

	log.Println("Database connected")

    m, err := migrate.New(
	"file://migrations",
	"postgres://postgres:password@localhost:5432/studentdb?sslmode=disable",
)
if err != nil {
	panic(err)
}

err = m.Up()
if err != nil {
	log.Printf("Migration error: %v", err)
} else {
	log.Println("Migrations applied successfully")
}

	router := routers.SetupRouter()
    
	port := os.Getenv("PORT")

	if port == "" {
		port = "8090"
	}

	log.Printf(
		"INFO: Server started on %s",
		port,
	)

	router.Run(":" + port)
}

	
