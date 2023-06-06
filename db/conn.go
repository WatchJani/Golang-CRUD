package db

import (
	"log"

	"github.com/jmoiron/sqlx"
)

var DB *sqlx.DB

func Connect() {
	var err error
	DB, err = sqlx.Open("postgres", "user=janko dbname=to-do password=JankoKondic72621@ sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}
}
