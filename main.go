package main

import (
	"context"
	"fmt"
	"log"

	"github.com/djengua/rifa-api/api"
	"github.com/djengua/rifa-api/util"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// var collection *mongo.Collection
var ctx = context.TODO()

// func init() {

// }

func main() {
	config, err := util.LoadConfig(".")
	if err != nil {
		log.Fatal("cannot read config. ", err)
	}

	options := options.Client().ApplyURI(config.DBUri)
	client, err := mongo.Connect(ctx, options)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Ping to DB ...")
	err = client.Ping(ctx, nil)
	if err != nil {
		log.Fatal(err)
	}

	database := client.Database("raffleDB")

	server, err := api.NewServer(config, database)
	if err != nil {
		log.Fatal("Cannot create server: ", err)
	}

	err = server.Start(config.ServerAddress)
	if err != nil {
		log.Fatal("Cannot start server: ", err)
	}
}
