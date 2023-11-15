package database

import (
	"database/sql"
	"fmt"

	go_ora "github.com/sijms/go-ora/v2"
)

type DB struct {
	database *sql.DB
}

// Sets up connection to UF Oracle database using user and password
func (db *DB) Initalize(user string, password string) {
	// Holds error values
	var err error

	// Form connection with database (db)
	db.database, err = sql.Open("oracle", go_ora.BuildUrl("oracle.cise.ufl.edu", 1521, "orcl", user, password, nil))
	if err != nil {
		panic(fmt.Errorf("error in sql.Open: %w", err))
	}

	// Make sure connection is solid by pinging
	err = db.database.Ping()
	if err != nil {
		panic(fmt.Errorf("error in db.Ping: %w", err))
	}
}
