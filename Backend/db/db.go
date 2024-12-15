package db

import (
	"fmt"
	"log"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

func Connect(host, port, user, password, dbname string) (*sqlx.DB, error) {
	log.Println("Connecting to database...")
	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		host, "5432", user, password, dbname) // Note: added port parameter

	db, err := sqlx.Connect("postgres", dsn)
	if err != nil {
		log.Println("Failed to connect to database:" + err.Error() + " " + dsn)
		return nil, err
	}

	return db, nil
}
