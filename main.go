package main

import (
	"log"

	"github.com/djengua/rifa-api/api"
	"github.com/djengua/rifa-api/util"
)

// var collection *mongo.Collection
// var ctx = context.TODO()

// func init() {
// 	options := options.Client().ApplyURI("mongodb://localhost:27017/")
// 	client, err := mongo.Connect(ctx, options)

// 	if err != nil {
// 		log.Fatal(err)
// 	}
// }

func main() {

	config, err := util.LoadConfig(".")
	if err != nil {
		log.Fatal("cannot read config. ", err)
	}

	server, err := api.NewServer(config)
	if err != nil {
		log.Fatal("Cannot create server: ", err)
	}

	err = server.Start(config.ServerAddress)
	if err != nil {
		log.Fatal("Cannot start server: ", err)
	}
}
