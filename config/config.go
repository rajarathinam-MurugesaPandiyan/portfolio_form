package config

import (
	"context"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

var DB *mongo.Database
var Client *mongo.Client

func InitializeDb(mongoUrl string) (*mongo.Database, error) {
	databaseName := Envs.DatabaseName
	serverAPI := options.ServerAPI(options.ServerAPIVersion1)
	opts := options.Client().ApplyURI(mongoUrl).SetServerAPIOptions(serverAPI).SetMaxPoolSize(100). // Ensure we don't run out of connections
													SetConnectTimeout(10 * time.Second)
	client, err := mongo.Connect(context.Background(), opts)
	if err != nil {
		panic(err)
	}

	err = client.Ping(context.Background(), readpref.Primary())
	if err != nil {
		panic(err)
	}
	Client = client
	DB = client.Database(databaseName)

	return DB, nil

}

func DisconnectMongoDB() {
	if Client != nil {
		err := Client.Disconnect(context.Background())
		if err != nil {
			log.Println("Error disconnecting MongoDB:", err)
		} else {
			log.Println("MongoDB disconnected successfully")
		}
	}
}
