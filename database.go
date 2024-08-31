package golang_database

import (
	"database/sql"
)

func GetConnection() *sql.DB {
	db, err := sql.Open("pgx", "user=postgres password=postgres dbname=belajar_golang_database sslmode=disable")

	if err != nil {
		panic(err)
	}

	db.SetMaxIdleConns(10)
	db.SetMaxOpenConns(100)
	db.SetConnMaxIdleTime(5)
	db.SetConnMaxLifetime(60 * 60 * 24)

	return db
}
