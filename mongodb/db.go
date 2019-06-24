package mongodb

/**
	Package para obtener los datos de la sesion
	y la base de datos
 */
import (
	"go.mongodb.org/mongo-driver/mongo/options"
	"time"
	"context"
	"go.mongodb.org/mongo-driver/mongo"
)

type Database struct {
	Sess *mongo.Client
	Database *mongo.Database
}

func (db * Database) Init() *Database  {

	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)

	clientOptions := options.Client().ApplyURI("mongodb+srv://brenq:mutant@cluster0-sxpqo.gcp.mongodb.net/test?retryWrites=true&w=majority")
	session, _ := mongo.Connect(ctx, clientOptions)


	database := session.Database(DATABASE_NAME)

	return &Database{session, database}
}
