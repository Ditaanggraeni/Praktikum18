package router

import (
	"create-migration/controller"

	"github.com/gorilla/mux"
)

func Router() *mux.Router {

	router := mux.NewRouter()

	router.HandleFunc("/api/mahasiswa", controller.AmbilSemuaMahasiswa).Methods("GET", "OPTIONS")
	router.HandleFunc("/api/mahasiswa/{id}", controller.AmbilMahasiswa).Methods("GET", "OPTIONS")
	router.HandleFunc("/api/mahasiswa", controller.TmbhMahasiswa).Methods("POST", "OPTIONS")
	router.HandleFunc("/api/mahasiswa/{id}", controller.UpdateMahasiswa).Methods("PUT", "OPTIONS")
	router.HandleFunc("/api/mahasiswa/{id}", controller.HapusMahasiswa).Methods("DELETE", "OPTIONS")

	return router
}