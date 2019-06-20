package mongodb

import (
	"gopkg.in/mgo.v2"
)
const (
	DATABASEHOST = "mongodb://localhost:27017"
	DATABASENAME= "dna"
)

type Database struct {
	Session *mgo.Session
	Database *mgo.Database
}

func (db * Database) Init() (*Database , error)  {

	session, err := mgo.DialWithInfo(&mgo.DialInfo{
		Addrs: []string{DATABASEHOST},
		Database: DATABASENAME,
	})

	if err != nil {
		return nil , err
	}

	database := session.DB(DATABASENAME)

	return &Database{session, database} , nil
}
