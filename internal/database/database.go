package database

import (
	"database/sql"
	"log"
)

func ConnectToDB() *sql.DB {
	var db *sql.DB
	var err error

	//TODO: this should be changed with environment variables~
	db, err = sql.Open("sqlite3", "../../assets/POPULATION_CENSUS_DB")

	if err != nil {
		log.Fatal(err)
	}

	return db
}

func CloseDatabase(db *sql.DB) {
	db.Close()
}
