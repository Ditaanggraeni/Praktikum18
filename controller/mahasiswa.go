package controller

import(
	"encoding/json" // package untuk enkode dan mendekode json menjadi struct dan sebaliknya
	"fmt"
	"strconv" // package yang digunakan untuk mengubah string menjadi tipe int

	"log"
	"net/http" // digunakan untuk mengakses objek permintaan dan respons dari api

	"create-migration/models" //models package dimana Mahasiswa didefinisikan

	"github.com/gorilla/mux" // digunakan untuk mendapatkan parameter dari router
	_ "github.com/lib/pq"    // postgres golang driver
)

type response struct {
	ID      int64  `json:"id,omitempty"`
	Message string `json:"message,omitempty"`
}

type Response struct {
	Status  int           `json:"status"`
	Message string        `json:"message"`
	Data    []models.Mahasiswa `json:"data"`
}

// TambahMahasiswa
func TmbhMahasiswa(w http.ResponseWriter, r *http.Request) {

	// create an empty user of type models.User
	// kita buat empty mahasiswa dengan tipe models.Mahasiswa
	var mahasiswa models.Mahasiswa

	// decode data json request ke mahasiswa
	err := json.NewDecoder(r.Body).Decode(&mahasiswa)

	if err != nil {
		log.Fatalf("Tidak bisa mendecode dari request body.  %v", err)
	}

	// panggil modelsnya lalu insert mahasiswa
	insertID := models.TambahMahasiswa(mahasiswa)

	// format response objectnya
	res := response{
		ID:      insertID,
		Message: "Data mahasiswa telah ditambahkan",
	}

	// kirim response
	json.NewEncoder(w).Encode(res)
}

// AmbilMahasiswa mengambil single data dengan parameter id
func AmbilMahasiswa(w http.ResponseWriter, r *http.Request) {
	// kita set headernya
	w.Header().Set("Context-Type", "application/x-www-form-urlencoded")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	// dapatkan idmahasiswa dari parameter request, keynya adalah "id"
	params := mux.Vars(r)

	// konversi id dari string ke int
	id, err := strconv.Atoi(params["id"])

	if err != nil {
		log.Fatalf("Tidak bisa mengubah dari string ke int.  %v", err)
	}

	// memanggil models ambilsatumahasiswa dengan parameter id yg nantinya akan mengambil single data
	mahasiswa, err := models.AmbilSatuMahasiswa(int64(id))

	if err != nil {
		log.Fatalf("Tidak bisa mengambil data mahasiswa. %v", err)
	}

	// kirim response
	json.NewEncoder(w).Encode(mahasiswa)
}

// Ambil semua data mahasiswa
func AmbilSemuaMahasiswa(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Context-Type", "application/x-www-form-urlencoded")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	// memanggil models AmbilSemuaMahasiswa
	mahasiswas, err := models.AmbilSemuaMahasiswa()

	if err != nil {
		log.Fatalf("Tidak bisa mengambil data. %v", err)
	}

	var response Response
	response.Status = 1
	response.Message = "Success"
	response.Data = mahasiswas

	// kirim semua response
	json.NewEncoder(w).Encode(response)
}

func UpdateMahasiswa(w http.ResponseWriter, r *http.Request) {

	// kita ambil request parameter idnya
	params := mux.Vars(r)

	// konversikan ke int yang sebelumnya adalah string
	id, err := strconv.Atoi(params["id"])

	if err != nil {
		log.Fatalf("Tidak bisa mengubah dari string ke int.  %v", err)
	}

	// buat variable mahasiswa dengan type models.Mahasiswa
	var mahasiswa models.Mahasiswa

	// decode json request ke variable mahasiswa
	err = json.NewDecoder(r.Body).Decode(&mahasiswa)

	if err != nil {
		log.Fatalf("Tidak bisa decode request body.  %v", err)
	}

	// panggil updatemahasiswa untuk mengupdate data
	updatedRows := models.UpdateMahasiswa(int64(id), mahasiswa)

	// ini adalah format message berupa string
	msg := fmt.Sprintf("Mahasiswa telah berhasil diupdate. Jumlah yang diupdate %v rows/record", updatedRows)

	// ini adalah format response message
	res := response{
		ID:      int64(id),
		Message: msg,
	}

	// kirim berupa response
	json.NewEncoder(w).Encode(res)
}

func HapusMahasiswa(w http.ResponseWriter, r *http.Request) {

	// kita ambil request parameter idnya
	params := mux.Vars(r)

	// konversikan ke int yang sebelumnya adalah string
	id, err := strconv.Atoi(params["id"])

	if err != nil {
		log.Fatalf("Tidak bisa mengubah dari string ke int.  %v", err)
	}

	// panggil fungsi hapusmahasiswa , dan convert int ke int64
	deletedRows := models.HapusMahasiswa(int64(id))

	// ini adalah format message berupa string
	msg := fmt.Sprintf("Mahasiswa sukses di hapus. Total data yang dihapus %v", deletedRows)

	// ini adalah format reponse message
	res := response{
		ID:      int64(id),
		Message: msg,
	}

	// send the response
	json.NewEncoder(w).Encode(res)
}