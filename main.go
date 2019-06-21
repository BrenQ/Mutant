package main

import (
	"github.com/gorilla/mux"
	"github.com/BrenQ/Mutant/handler"
	"log"
	"net/http"
)


func main() {

	router := mux.NewRouter()

	// Ruta donde se realiza la verificacion si un ADN recibido pertenece a un humano o mutante/
	router.HandleFunc("/mutant", handler.Mutant).Methods("POST")
	// Ruta ´para obtener la estadistica del ADN
	router.HandleFunc("/stats", handler.Stats).Methods("GET")

	log.Fatal(http.ListenAndServe(":6000", router))
}
