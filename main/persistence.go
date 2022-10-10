package main

import (
	"database/sql"
	"fmt"
	"github.com/go-sql-driver/mysql"
	"log"
	"os"
)

var db *sql.DB

func ConnectToDatabase() {
	// Capture connection properties.
	cfg := mysql.Config{
		User:   os.Getenv("DB_USER_LOCAL"),
		Passwd: os.Getenv("DB_PASSWORD_LOCAL"),
		Net:    "tcp",
		Addr:   os.Getenv("DB_ADDRESS_LOCAL"),
		DBName: os.Getenv("DB_NAME_LOCAL"),
	}
	// Get a database handle.
	var err error
	db, err = sql.Open("mysql", cfg.FormatDSN())
	if err != nil {
		log.Fatal(err)
	}

	pingErr := db.Ping()
	if pingErr != nil {
		log.Fatal(pingErr)
	}
	fmt.Println("Successfully connected to database " + os.Getenv("DB_NAME_LOCAL") + ". Congratulations.")
}
