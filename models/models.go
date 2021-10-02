package models

import (
	"database/sql"
	"fmt"
	"create-migration/config"
	"log"

	_ "github.com/lib/pq" // postgres golang driver
)

// Mahasiswa schema dari tabel Mahasiswa
// kita coba dengan jika datanya null
// jika return datanya ada yg null, silahkan pake NullString, contohnya dibawah
// Penulis       config.NullString `json:"penulis"`
type Mahasiswa struct {
	ID            int64  `json:"id"`
	Nama    string `json:"nama"`
	Nim       string `json:"nim"`
	Jurusan string `json:"jurusan"`
}

func TambahMahasiswa(mahasiswa Mahasiswa) int64 {

	// mengkoneksikan ke db postgres
	db := config.CreateConnection()

	// kita tutup koneksinya di akhir proses
	defer db.Close()

	// kita buat insert query
	// mengembalikan nilai id akan mengembalikan id dari mahasiswa yang dimasukkan ke db
	sqlStatement := `INSERT INTO mahasiswa (nama, nim, jurusan) VALUES ($1, $2, $3) RETURNING id`

	// id yang dimasukkan akan disimpan di id ini
	var id int64

	// Scan function akan menyimpan insert id didalam id id
	err := db.QueryRow(sqlStatement, mahasiswa.Nama, mahasiswa.Nim, mahasiswa.Jurusan).Scan(&id)

	if err != nil {
		log.Fatalf("Tidak Bisa mengeksekusi query. %v", err)
	}

	fmt.Printf("Insert data single record %v", id)

	// return insert id
	return id
}

// ambil satu mahasiswa
func AmbilSemuaMahasiswa() ([]Mahasiswa, error) {
	// mengkoneksikan ke db postgres
	db := config.CreateConnection()

	// kita tutup koneksinya di akhir proses
	defer db.Close()

	var mahasiswas []Mahasiswa

	// kita buat select query
	sqlStatement := `SELECT * FROM mahasiswa`

	// mengeksekusi sql query
	rows, err := db.Query(sqlStatement)

	if err != nil {
		log.Fatalf("tidak bisa mengeksekusi query. %v", err)
	}

	// kita tutup eksekusi proses sql qeurynya
	defer rows.Close()

	// kita iterasi mengambil datanya
	for rows.Next() {
		var mahasiswa Mahasiswa

		// kita ambil datanya dan unmarshal ke structnya
		err = rows.Scan(&mahasiswa.ID, &mahasiswa.Nama, &mahasiswa.Nim, &mahasiswa.Jurusan)

		if err != nil {
			log.Fatalf("tidak bisa mengambil data. %v", err)
		}

		// masukkan kedalam slice mahasiswas
		mahasiswas = append(mahasiswas, mahasiswa)

	}

	// return empty mahasiswa atau jika error
	return mahasiswas, err
}

// mengambil satu mahasiswa
func AmbilSatuMahasiswa(id int64) (Mahasiswa, error) {
	// mengkoneksikan ke db postgres
	db := config.CreateConnection()

	// kita tutup koneksinya di akhir proses
	defer db.Close()

	var mahasiswa Mahasiswa

	// buat sql query
	sqlStatement := `SELECT * FROM mahasiswa WHERE id=$1`

	// eksekusi sql statement
	row := db.QueryRow(sqlStatement, id)

	err := row.Scan(&mahasiswa.ID, &mahasiswa.Nama, &mahasiswa.Nim, &mahasiswa.Jurusan)

	switch err {
	case sql.ErrNoRows:
		fmt.Println("Tidak ada data yang dicari!")
		return mahasiswa, nil
	case nil:
		return mahasiswa, nil
	default:
		log.Fatalf("tidak bisa mengambil data. %v", err)
	}

	return mahasiswa, err
}

// update user in the DB
func UpdateMahasiswa(id int64, mahasiswa Mahasiswa) int64 {

	// mengkoneksikan ke db postgres
	db := config.CreateConnection()

	// kita tutup koneksinya di akhir proses
	defer db.Close()

	// kita buat sql query create
	sqlStatement := `UPDATE mahasiswa SET nama=$2, nim=$3, jurusan=$4 WHERE id=$1`

	// eksekusi sql statement
	res, err := db.Exec(sqlStatement, id, mahasiswa.Nama, mahasiswa.Nim, mahasiswa.Jurusan)

	if err != nil {
		log.Fatalf("Tidak bisa mengeksekusi query. %v", err)
	}

	// cek berapa banyak row/data yang diupdate
	rowsAffected, err := res.RowsAffected()

	//kita cek
	if err != nil {
		log.Fatalf("Error ketika mengecheck rows/data yang diupdate. %v", err)
	}

	fmt.Printf("Total rows/record yang diupdate %v\n", rowsAffected)

	return rowsAffected
}

func HapusMahasiswa(id int64) int64 {

	// mengkoneksikan ke db postgres
	db := config.CreateConnection()

	// kita tutup koneksinya di akhir proses
	defer db.Close()

	// buat sql query
	sqlStatement := `DELETE FROM mahasiswa WHERE id=$1`

	// eksekusi sql statement
	res, err := db.Exec(sqlStatement, id)

	if err != nil {
		log.Fatalf("tidak bisa mengeksekusi query. %v", err)
	}

	// cek berapa jumlah data/row yang di hapus
	rowsAffected, err := res.RowsAffected()

	if err != nil {
		log.Fatalf("tidak bisa mencari data. %v", err)
	}

	fmt.Printf("Total data yang terhapus %v", rowsAffected)

	return rowsAffected
}