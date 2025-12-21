package data

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
)

var DB *sql.DB

func InitializePostgres(connection string) {
	db, err := sql.Open("postgres", connection)

	if err != nil {
		log.Fatal(err)
	}

	err = db.Ping()

	if err != nil {
		log.Fatal(err)
	}
	DB = db
}

func GetDB() *sql.DB {
	return DB
}
