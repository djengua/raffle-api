package coreapi

import (
	"fmt"

	"github.com/djengua/raffle-api/util"
	"github.com/djengua/token"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
)

type Server struct {
	database   *mongo.Database
	router     *gin.Engine
	tokenMaker token.Maker

	config util.Config
}

func NewServer(config util.Config, database *mongo.Database) (*Server, error) {
	fmt.Println("Config Server ...")
	tokenMaker, err := token.NewPasetoMaker(config.TokenSymmetricKey)
	if err != nil {
		return nil, fmt.Errorf("cannot create token maker: %w", err)
	}

	server := &Server{
		config:     config,
		database:   database,
		tokenMaker: tokenMaker,
	}

	server.setupRouter()

	return server, nil
}

func (s *Server) setupRouter() {
	router := gin.Default()
	router.LoadHTMLGlob("./templates/*")
	router.Static("/static", "./static")
	go ListenToWebSocketChannel()

	router.GET("/hello", s.hello)
	router.GET("/", s.homePage)
	router.GET("/ticket-suggestion", s.TicketSuggest)
	router.GET("/mel-suggestion", s.MelSuggest)
	router.GET("/ws", s.WebSocketEndpoint)

	authRoutes := router.Group("/").Use(authMiddleware(s.tokenMaker))

	authRoutes.GET("/raffle/all", s.getAllRaffle)
	// router.GET("/raffle/all", s.getAllRaffle)

	router.POST("/raffle", s.createRaffle)
	router.GET("/raffle/:id", s.getRaffleById)
	router.PUT("/raffle/add-participant", s.addParticipant)
	router.PUT("/raffle/add-ticket-to-participant", s.addTicketToParticipant)
	router.PUT("/raffle/delete-participant", s.deleteParticipant)
	router.POST("/raffle/discard-ticket", s.discardTicket)
	router.POST("/raffle/winner", s.winner)

	router.POST("/user", s.createUsers)
	router.GET("/user/all", s.getAllUsers)
	router.POST("/login", s.loginUser)
	router.POST("/refresh", s.refreshToken)

	s.router = router
}

// func (s *Server) setupRouterMux(){
// 	r := mux.NewRouter()
//     r.HandleFunc("/", HomeHandler)
//     r.HandleFunc("/products", ProductsHandler)
//     r.HandleFunc("/articles", ArticlesHandler)
// 	s.router
// }

func (s *Server) Start(address string) error {
	return s.router.Run(address)
}

func errorResponse(err error) gin.H {
	return gin.H{"error": err.Error(), "data": nil}
}
