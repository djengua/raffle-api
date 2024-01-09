package coreapi

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

var uConn = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin:     func(r *http.Request) bool { return true },
}

type wsResponse struct {
	Action      string `json:"action"`
	Message     string `json:"message"`
	MessageType string `json:"message_type"`
}

func (s *Server) WsEndpoint(ctx *gin.Context) {
	conn, err := uConn.Upgrade(ctx.Writer, ctx.Request, nil)
	if err != nil {
		log.Println(err)
	}

	log.Println("Client connected to endpoint")

	defer conn.Close()

	response := wsResponse{
		Message: "<small>Connected to server.</small>",
	}
	err = conn.WriteJSON(response)
	if err != nil {
		log.Println(err)
	}

}
