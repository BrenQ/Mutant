package configuration

import (
	"gopkg.in/mgo.v2"
	"log"
)

/**
	Database
 */
type Db struct {
	client *mgo.Database
}

func (Db *Db) init(){
	client, err := mgo.Dial("mongodb://localhost:27017")

	if err != nil {
		log.Fatal(err)
	}

	Db.client = client.DB("dna")
}