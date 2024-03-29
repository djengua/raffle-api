package main

import (
	"context"
	"fmt"
	"log"

	coreapi "github.com/djengua/raffle-api/core-api"
	"github.com/djengua/raffle-api/util"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func main() {
	var ctx = context.TODO()
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

	database := client.Database(config.DbName)

	server, err := coreapi.NewServer(config, database)
	if err != nil {
		log.Fatal("Cannot create server: ", err)
		client.Disconnect(ctx)
	}

	err = server.Start(config.ServerAddress)
	if err != nil {
		log.Fatal("Cannot start server: ", err)
		client.Disconnect(ctx)
	}
}
