package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"github.com/gorilla/mux"
	"log"
	//"encoding/json"
)
/**
	Estructura para almacenar el request recibido
 */
type DnaRequest struct {
	Dna [] string `json:"Dna,omitempty"`
}
/**
	Estructura que almacena la estructura y los datos necesarios para el dna
 */
type Dna struct {
	N int `json:"N,omitempty"`
	Size int `json:"Size,omitempty"`
	Pattern int `json:"Pattern,omitempty"`
	Sequence []rune `json:"Sequence,omitempty"`
}

/**
	@method Inicializa los valores del struct dna
 */
func (d Dna) init() {
	d.N = 0
	d.Size = 0
	d.Pattern = 0
	d.Sequence = make([]rune, 0)
}

/**
	@method Almacena los valores necesarios para operar
 */
func (d * Dna) register(data[]string ) {
	d.N = len(data)

	for _, value := range data {
		for _, letter := range value {
			d.Sequence = append(d.Sequence, letter)
		}
	}
	d.Size = len(d.Sequence)
}

/**
	@method  Verifica que un humano es mutando
 */
func(d * Dna) isMutant(data[] string) bool {
	const Sequences = 2 // Cantidad de secuencias requeridas para comprobar que es un mutante

	d.init()
	d.register(data)
	// Recorro los elementos para verificar si se registra un patron
	for index,letter := range d.Sequence {
		if d.Pattern > Sequences {
			return true
		}
		d.iterate(index, letter)
	}
	return false
}
/**
	@method Funcion auxiliar que recorre un array de funciones.
			Este array de funciones contiene la logica que define
			con que valores comparar segun la ubicacion
 */
func(d *Dna) iterate(index int, letter rune) {

	functions := []func(index int, n int) int {
		func(index int, n int ) int { return index + 1 },
		func(index int , n int) int { return index + n },
		func(index int , n int) int { return index - n - 1 },
		func(index int , n int) int { return index - n + 1 },
		func(index int , n int) int { return index + n - 1 },
		func(index int , n int) int { return index + n + 1 },
	}

	for _, function:=range functions {
		if d.check(function , index ,  letter , 0) {
			d.Pattern++
		}
	}
}

/**
	@method Funcion recursiva que segun la direccion calculada a travez del indica
			verifica si existe un patron y lo acumula.
			Se comprueba un patron si este tiene 4 letras contiguas segun su direcicon
 */
func(d *Dna) check(function func(index int, n int) int, index int, letter rune , pattern int) bool {
	result := false
	newIndex := function(index, d.N)

	if pattern == 3 {
		result = true
	}

	if (newIndex > 0 && newIndex < d.Size)  && (d.Sequence[newIndex] == letter && pattern < 4 )  {
		pattern++
		return d.check(function, newIndex , letter, pattern)
	}
	return result
}

func main() {

	router := mux.NewRouter()
	/**
		@method  Ruta donde se realiza la verificacion si un ADN recibido pertenece a un humano o mutante
	 */
	router.HandleFunc("/mutant", func(writer http.ResponseWriter, request *http.Request) {

		writer.Header().Set("Content-Type", "application/json")

		var DnaSequence DnaRequest
		err := json.NewDecoder(request.Body).Decode(&DnaSequence)
		if err != nil {
			panic(err)
		}
		sequence := Dna{}
		fmt.Println(DnaSequence.Dna)
		result := sequence.isMutant(DnaSequence.Dna)

		writer.Header().Set("Content-Type", "application/json")
		if result {
			writer.WriteHeader(http.StatusOK)
			return
		}
		writer.WriteHeader(http.StatusForbidden)

	}).Methods("POST")

	log.Fatal(http.ListenAndServe(":5000", router))
}
