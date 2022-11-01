package bd

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
)

var MongoC = Connection()
var clientOptions = options.Client().ApplyURI("mongodb+srv://jormoreno:opcN9Bl1iB3w5NDV@cluster0.6rgun0c.mongodb.net/?retryWrites=true&w=majority")

func Connection() *mongo.Client {
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.Fatal(err.Error())
		return client
	}
	err = client.Ping(context.TODO(), nil)
	if err != nil {
		log.Fatal(err)
		return client
	}
	log.Printf("Sucess Connection DB")
	return client
}
func CheckConnection() int {
	err := MongoC.Ping(context.TODO(), nil)
	if err != nil {
		log.Fatal(err.Error())
		return 0
	}
	return 1
}
