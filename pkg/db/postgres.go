package db

import (
	"database/sql"
	"fmt"
	"log"
	"time"

	_ "github.com/lib/pq"
)

var DB *sql.DB

func InitDB(connectionString string) {
	var err error
	DB, err = sql.Open("postgres", connectionString)
	if err != nil {
		log.Fatal("Failed to open connection to database: ", err)
	}

	err = DB.Ping()
	if err != nil {
		log.Fatal("Failed to ping database: ", err)
	}

	fmt.Println("Connected to the database successfully!")

	// Connection Pooling Configuration
	// Adjust these based on your server resources and traffic
	DB.SetMaxOpenConns(25)                 // Max open connections to the DB
	DB.SetMaxIdleConns(25)                 // Max idle connections to keep open
	DB.SetConnMaxLifetime(5 * time.Minute) // How long a connection can be reused
}
