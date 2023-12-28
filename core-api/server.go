package coreapi

import (
	"fmt"

	"github.com/djengua/raffle-api/util"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
)

type Server struct {
	database *mongo.Database
	router   *gin.Engine
	config   util.Config
}

func NewServer(config util.Config, database *mongo.Database) (*Server, error) {
	fmt.Println("Config Server ...")
	server := &Server{
		config:   config,
		database: database,
		// tokenMaker: tokenMaker,
	}

	server.setupRouter()

	return server, nil
}

func (s *Server) setupRouter() {
	router := gin.Default()

	router.GET("/hello", s.hello)

	router.GET("/raffle/all", s.getAllRaffle)
	router.POST("/raffle", s.createRaffle)
	router.GET("/raffle/:id", s.getRaffleById)
	router.PUT("/raffle/add-participant", s.addParticipant)
	router.PUT("/raffle/add-ticket-to-participant", s.addTicketToParticipant)
	router.PUT("/raffle/delete-participant", s.deleteParticipant)
	router.POST("/raffle/winner/:id", s.winner)

	s.router = router
}

func (s *Server) Start(address string) error {
	return s.router.Run(address)
}

func errorResponse(err error) gin.H {
	return gin.H{"error": err.Error(), "data": nil}
}
