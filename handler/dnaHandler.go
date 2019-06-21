package handler

import (
	"encoding/json"
	"github.com/BrenQ/Mutant/models"
	"github.com/BrenQ/Mutant/mongodb"
	"github.com/BrenQ/Mutant/service"
	"gopkg.in/mgo.v2/bson"
	"log"
	"net/http"
)
/**
	Almaceno el request recibido
 */
type DnaRequest struct {
	Dna [] string `json:"Dna,omitempty"`
}

/**
	Funcion para validar en base a los data enviada si un DNA pertenece a un mutante
 */
func Mutant(writer  http.ResponseWriter, request *http.Request)  {

	var Db mongodb.Database
	// Variable donde obtengo la sesion de la base de datos

	var Database  = Db.Init()

	var Sess = Database.Sess.Copy()

	writer.Header().Set("Content-Type", "application/json")

	var DnaSequence *DnaRequest
	err := json.NewDecoder(request.Body).Decode(&DnaSequence)

	if err != nil {
		panic(err)
	}

	// Obtengo el servicio y le paso la referencia al modelo DNA
	var sequenceDna *service.DnaService
	var model= &models.Dna{}
	sequenceDna = &service.DnaService{model}

	var response = models.Response{}

	 if sequenceDna.IsMutant(DnaSequence.Dna) {
		 response.Add(200 , "La secuencia pertenece a un mutante")
	 } else {
	 	response.Add(403, "La secuencia pertenece a un humano")
	 }


	writer.Header().Set("Content-Type", "application/json")
	err = Sess.DB("dna").C("sequence").Insert(model)

	if err != nil {
		log.Print(err)
	}

	writer.WriteHeader(response.Code)
	_ = json.NewEncoder(writer).Encode(response)

}
/**
	Funcion para calcular la estadistica solicitada
 */
func Stats (writer http.ResponseWriter, request *http.Request) {

	var Db mongodb.Database

	var Database  = Db.Init()

	var Sess = Database.Sess.Copy()

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
			"$project": bson.M{"_id": 0,
				"count_mutant_dna": "$count_mutant_dna",
				"count_human_dna":  "$count_human_dna",
				"ratio": bson.M{
					"$divide": []interface{}{"$count_mutant_dna", "$count_human_dna"}},
			},
		},
	}

	var result []bson.M
	_ = Sess.DB("dna").C("sequence").Pipe(pipeline).All(&result)

	_ = json.NewEncoder(writer).Encode(result)

}