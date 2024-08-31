package golang_database

import (
	"context"
	"fmt"
	"testing"
)

func TestExecSql(t *testing.T) {

	db := GetConnection()
	defer db.Close()

	ctx := context.Background()

	script := "INSERT INTO customer(id, name) VALUES('mia', 'elvie')"
	_, err := db.ExecContext(ctx, script)
	if err != nil {
		panic(err)
	}

	fmt.Println("Success insert new customer")
}

func TestUpdateSql(t *testing.T) {
	db := GetConnection() // Mendapatkan koneksi ke database
	defer db.Close()      // Menutup koneksi setelah fungsi selesai

	ctx := context.Background() // Membuat context kosong

	script := "UPDATE customer SET name = 'Bakso' WHERE id = 'ari'" // Query SQL untuk meng-update data
	_, err := db.ExecContext(ctx, script)                           // Menjalankan query dengan context yang diberikan
	if err != nil {                                                 // Memeriksa jika terjadi error
		panic(err) // Menghentikan eksekusi jika ada error
	}

	fmt.Println("Success update customer") // Output jika operasi berhasil
}

func TestDeleteSql(t *testing.T) {
	db := GetConnection() // Mendapatkan koneksi ke database
	defer db.Close()      // Menutup koneksi setelah fungsi selesai

	ctx := context.Background() // Membuat context kosong

	script := "DELETE FROM customer WHERE id = 'mia'" // Query SQL untuk menghapus data
	_, err := db.ExecContext(ctx, script)             // Menjalankan query dengan context yang diberikan
	if err != nil {                                   // Memeriksa jika terjadi error
		panic(err) // Menghentikan eksekusi jika ada error
	}

	fmt.Println("Success delete customer") // Output jika operasi berhasil
}

func TestQuerySql(t *testing.T) {

	db := GetConnection()
	defer db.Close()

	ctx := context.Background()

	script := "SELECT id, name FROM customer"
	rows, err := db.QueryContext(ctx, script)
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	for rows.Next() {
		var id, name string
		err := rows.Scan(&id, &name)
		if err != nil {
			panic(err)
		}
		fmt.Println("id:", id)
		fmt.Println("name:", name)

	}
	defer rows.Close()

}
