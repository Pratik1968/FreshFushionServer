package db

import (
	"context"
	"os"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)
var client *mongo.Client

func  Init() {
	uri := os.Getenv("MONGO_URI")
	serverAPI := options.ServerAPI(options.ServerAPIVersion1)
	opts:= options.Client().ApplyURI(uri).SetServerAPIOptions(serverAPI)
	var err error
	client,err = mongo.Connect(context.TODO(),opts)
	if err!=nil{
panic(err)
	}
	
}
