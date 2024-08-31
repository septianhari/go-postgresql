package golang_database

import (
	"database/sql"
	"testing"

	_ "github.com/jackc/pgx/v5/stdlib"
)

func TestEmpty(t *testing.T) {

}

func TestOpenConnection(t *testing.T) {
	// Menggunakan "pgx" sebagai driver PostgreSQL
	db, err := sql.Open("pgx", "user=postgres password=postgres dbname=belajar_golang_database sslmode=disable")

	if err != nil {
		panic(err)
	}

	defer db.Close()
}
