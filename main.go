package main

import (
	"encoding/json"
	"github.com/BrenQ/Mutant/models"
	"github.com/BrenQ/Mutant/service"
	"github.com/BrenQ/Mutant/mongodb"
	"github.com/gorilla/mux"
	"gopkg.in/mgo.v2/bson"

	//"gopkg.in/mgo.v2/bson"
	"log"
	"net/http"
)

type DnaRequest struct {
	Dna [] string `json:"Dna,omitempty"`
}

/**
@method Funcion recursiva que segun la direccion calculada a travez del indica
		verifica si existe un patron y lo acumula.
		Se comprueba un patron si este tiene 4 letras contiguas segun su direcicon
*/

func main() {

	var Db mongodb.Database
	// Variable donde obtengo la sesion de la base de datos

	var Database, err = Db.Init()

	if err != nil {
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

		var DnaSequence *DnaRequest
		err := json.NewDecoder(request.Body).Decode(&DnaSequence)
		if err != nil {
			panic(err)
		}

		var sequenceDna *service.DnaService

		var model = &models.Dna{}
		sequenceDna = &service.DnaService{model}

		var response = sequenceDna.IsMutant(DnaSequence.Dna)

		writer.Header().Set("Content-Type", "application/json")
		err = Sess.DB("dna").C("sequence").Insert(model)

		if err != nil {
			log.Print(err)
		}
		_ = json.NewEncoder(writer).Encode(response)

	}).Methods("POST")

	   	router.HandleFunc("/stats", func(writer http.ResponseWriter, request *http.Request) {

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
