package mongodb

/**
	Package para obtener los datos de la sesion
	y la base de datos
 */
import (
	"gopkg.in/mgo.v2"
	"time"
)
const (
	DATABASE_HOST = "127.0.0.1:27017"
	DATABASE_NAME = "dna"
	DATABASE_USER = "brenq"
	DATABASE_PASS = "mutant"
)

type Database struct {
	Sess *mgo.Session
	Database *mgo.Database
}

func (db * Database) Init() (*Database )  {

	session, err := mgo.DialWithInfo(&mgo.DialInfo{
		Addrs:    []string{DATABASE_HOST},
		Timeout:  30 * time.Second,
		Database: DATABASE_NAME,
		Username: DATABASE_USER,
		Password: DATABASE_PASS,
	})

	if err != nil {
		panic(err)
	}

	database := session.DB(DATABASE_NAME)

	return &Database{session, database}
}
