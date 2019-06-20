package main

import (
	"encoding/json"
	"github.com/BrenQ/Mutant/mongodb"
	"github.com/gorilla/mux"
	"gopkg.in/mgo.v2/bson"
	"log"
	"net/http"
)


/**
Estructura para almacenar el request recibido
*/
type DnaRequest struct {
	Dna [] string `json:"Dna,omitempty"`
}

/**
Estructura auxiliar para almacenar la respuestas
*/
type Response struct {
	Code    int    `json:"code,omitempty"`
	Message string `json:"message,omitempty"`
}

/**
Estructura que almacena la estructura y los datos necesarios para el dna
*/
type Dna struct {
	ID         bson.ObjectId `json:"_id,omitempty" bson:"_id,omitempty" `
	N          int           `json:"N,omitempty" bson:"n,omitempty" `
	Size       int           `json:"Size,omitempty" bson:"size,omitempty"`
	Pattern    int           `json:"Pattern,omitempty" bson:"pattern,omitempty"`
	Sequence   []rune        `json:"Sequence,omitempty"  bson:"sequence,omitempty"`
	IsMutant   bool          `json:"IsMutant,omitempty"  bson:"IsMutant,omitempty"`
	Response   Response      `json:"Response,omitempty"  bson:"response,omitempty"`
	Iterations int           `json:"iterations,omitempty"  bson:"iterations,omitempty"`
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

func (r *Response) add(code int, message string) {
	if r != nil {
		r.Code = code
		r.Message = message
	}
}

/**
@method Almacena los valores necesarios para operar
*/
func (d *Dna) register(data []string) bool {

	if len(data) == 0 {
		d.Response.add(400, "La secuencia está vacía")
		return false
	}
	d.N = len(data)
	for _, value := range data {
		if len(value) != d.N {
			d.Response.add(400, "La matriz ingresada debe ser cuadrada")
			return false
		}
		for _, letter := range value {

			if ! validateLetter(letter) {
				d.Response.add(400, "Las letras permitidas son A,T,G,C")
				return false
			}
			d.Sequence = append(d.Sequence, letter)
		}
	}
	d.Size = len(d.Sequence)
	return true
}

/**
@method  Verifica que un humano es mutando
*/
func (d *Dna) isMutant(data [] string) bool {
	const Sequences = 2 // Cantidad de secuencias requeridas para comprobar que es un mutante

	d.init()
	registered := d.register(data)
	// Recorro los elementos para verificar si se registra un patron
	if ! registered {
		return false
	}

	for index, letter := range d.Sequence {
		if d.Pattern > Sequences {
			d.Response.add(200, "La secuencia pertenece a un mutante")
			d.IsMutant = true
			return true
		}
		d.iterate(index, letter)
	}

	d.IsMutant = false
	d.Response.add(403, "La secuencia ingresada no pertenece a un mutante")
	return false
}

/**
@method Funcion auxiliar que recorre un array de funciones.
		Este array de funciones contiene la logica que define
		con que valores comparar segun la ubicacion
*/
func (d *Dna) iterate(index int, letter rune) {

	functions := []func(index int, n int) int{
		func(index int, n int) int { return index + 1 },
		func(index int, n int) int { return index + n },
		func(index int, n int) int { return index - n - 1 },
		func(index int, n int) int { return index - n + 1 },
		func(index int, n int) int { return index + n - 1 },
		func(index int, n int) int { return index + n + 1 },
	}

	for _, function := range functions {
		if d.check(function, index, letter, 0) {
			d.Pattern++
		}
	}
}

/**
Funcion auxiliar para verificar si una letra es valida
*/

func validateLetter(letter rune) bool {
	letters := []rune{'A', 'T', 'C', 'G'}

	for _, value := range letters {
		if value == letter {
			return true
		}
	}
	return false
}

/**
@method Funcion recursiva que segun la direccion calculada a travez del indica
		verifica si existe un patron y lo acumula.
		Se comprueba un patron si este tiene 4 letras contiguas segun su direcicon
*/
func (d *Dna) check(function func(index int, n int) int, index int, letter rune, pattern int) bool {
	result := false
	newIndex := function(index, d.N)

	if pattern == 3 {
		result = true
	}

	if (newIndex > 0 && newIndex < d.Size) && (d.Sequence[newIndex] == letter && pattern < 4) {
		pattern++
		return d.check(function, newIndex, letter, pattern)
	}
	d.Iterations++
	return result
}

func main() {

	var Db mongodb.Database
	// Variable donde obtengo la sesion de la base de datos

	var Database , err = Db.Init()

	if err != nil{
		panic(err)
	}
	
	var Sess = Database.Sess.Copy()

	defer Sess.Close()

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

		_ = sequence.isMutant(DnaSequence.Dna)

		writer.Header().Set("Content-Type", "application/json")
		err = Sess.DB("dna").C("sequence").Insert(sequence)

		if err != nil {
			log.Print(err)
		}
		response := json.NewEncoder(writer).Encode(sequence.Response)

		if response != nil {
			writer.WriteHeader(sequence.Response.Code)
		}
	}).Methods("POST")

	router.HandleFunc("/stats", func(writer http.ResponseWriter, request *http.Request) {

		sequence := Dna{}
		sequence.init()
		pipeline := []bson.M{
			bson.M{"$group":
				bson.M{"_id": 0,
					"count_mutant_dna": bson.M{
						"$sum":
						bson.M{"$cond":
						[]interface{}{bson.M{"$eq": []interface{}{"$IsMutant", true}}, 1, 0},
						},
					},
					"count_human_dna": bson.M{
						"$sum":
						bson.M{"$cond":
						[]interface{}{bson.M{"$ifNull": []interface{}{"$IsMutant", false}}, 0, 1},
						},
					},
				},
			},
			bson.M{
				"$project": bson.M{"_id":0,
					"count_mutant_dna": "$count_mutant_dna",
					"count_human_dna": "$count_human_dna",
					"ratio": bson.M{
						"$divide": []interface{}{"$count_mutant_dna", "$count_human_dna"}},
				},
			},
		}

var result []bson.M
_ = Sess.DB("dna").C("sequence").Pipe(pipeline).All(&result)

_ = json.NewEncoder(writer).Encode(result)

}).Methods("GET")

log.Fatal(http.ListenAndServe(":6000", router))
}
