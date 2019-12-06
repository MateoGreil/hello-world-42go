package config

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

var db *sql.DB

const (
	host     = ""
	port     = 5432
	user     = "readonly"
	password = ""
	dbname   = ""
)

func DatabaseInit() {
	var err error

	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s",
		host, port, user, password, dbname)
	db, err = sql.Open("postgres", psqlInfo)

	if err != nil {
		log.Fatal(err)
	}
}

func Db() *sql.DB {
	return db
}
