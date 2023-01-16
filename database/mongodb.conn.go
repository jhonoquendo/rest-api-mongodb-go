package database

import (
	"context"
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var (
	usr      = "root"
	pwd      = ""
	host     = "localhost"
	port     = 27017
	database = "mongodbgo"
)

func GetCollection(collection string) *mongo.Collection {
	//uri := fmt.Sprintf("mongodb://%s:%s@%s:%d", usr, pwd, host, port)
	uri := fmt.Sprintf("mongodb://localhost/mongodbgo")

	client, err := mongo.NewClient(options.Client().ApplyURI(uri))

	if err != nil {
		fmt.Println("Error con obtener la uri de mongo")
	}

	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	err = client.Connect(ctx)

	if err != nil {
		fmt.Println("Error de conexion")
	}

	return client.Database(database).Collection(collection)

}
