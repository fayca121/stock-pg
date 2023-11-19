package config

import (
	"database/sql"
	"fmt"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"log"
	"os"
)

func ConnectDB() *sql.DB {

	err := godotenv.Load()

	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	dbDriver := os.Getenv("DB_DRIVER")
	dbUrl := os.Getenv("DB_URL")

	db, err := sql.Open(dbDriver, dbUrl)

	if err != nil {
		log.Fatalln("Failed to connect postgres database")
	}

	// check the connection
	if err = db.Ping(); err != nil {
		panic(err)
	}

	fmt.Println("Successfully connected!")
	return db
}

func DisconnectDB(db *sql.DB) {
	err := db.Close()
	if err != nil {
		return
	}
	fmt.Println("Successfully disconnected!")
}
