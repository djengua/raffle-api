package coreapi

import (
	"fmt"
	"log"
	"net/http"
	"sort"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

var webSocketChannel = make(chan WebSocketPayload)

var clients = make(map[WebSocketConnection]string)

var uConn = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin:     func(r *http.Request) bool { return true },
}

type WebSocketConnection struct {
	*websocket.Conn
}

type WebSocketResponse struct {
	Action         string   `json:"action"`
	Message        string   `json:"message"`
	MessageType    string   `json:"message_type"`
	ConnectedUsers []string `json:"connected_users"`
}

type WebSocketPayload struct {
	Action   string              `json:"action"`
	Username string              `json:"username"`
	Message  string              `json:"Message"`
	Conn     WebSocketConnection `json:"-"`
}

func (s *Server) WebSocketEndpoint(ctx *gin.Context) {
	conn, err := uConn.Upgrade(ctx.Writer, ctx.Request, nil)
	if err != nil {
		log.Println(err)
	}
	//defer conn.Close()

	log.Println("Client connected to endpoint")

	response := WebSocketResponse{
		Message: "<em><small>Connected to server.</small></em>",
	}

	cn := WebSocketConnection{conn}
	clients[cn] = ""
	err = conn.WriteJSON(response)
	if err != nil {
		log.Println(err)
	}
	go ListenForWebSocket(&cn)
}

func ListenForWebSocket(conn *WebSocketConnection) {
	fmt.Println("Listening...")
	defer func() {
		if r := recover(); r != nil {
			log.Println("Error ", fmt.Sprintf("%v", r))
		}
	}()

	var payload WebSocketPayload

	for {
		err := conn.ReadJSON(&payload)
		if err != nil {
			// do nothing
		} else {
			payload.Conn = *conn
			webSocketChannel <- payload
		}
	}
}

func ListenToWebSocketChannel() {
	var response WebSocketResponse

	for {
		e := <-webSocketChannel
		switch e.Action {
		case "username":
			clients[e.Conn] = e.Username
			users := getUserList()
			response.Action = "list_users"
			response.ConnectedUsers = users
			broadcastToAll(response)
		case "left":
			response.Action = "list_users"
			delete(clients, e.Conn)
			users := getUserList()
			response.ConnectedUsers = users
			broadcastToAll(response)
		case "broadcast":
			response.Action = "broadcast"
			response.Message = fmt.Sprintf("<strong>%s</strong>: %s", e.Username, e.Message)
			broadcastToAll(response)
		}

		// response.Action = "got here"
		// response.Message = fmt.Sprintf("Some message, and action was %s", e.Action)
		// broadcastToAll(response)
	}
}

func getUserList() []string {
	var userList []string
	for _, x := range clients {
		if x != "" {
			userList = append(userList, x)
		}
	}
	sort.Strings(userList)
	return userList
}

func broadcastToAll(response WebSocketResponse) {
	for client := range clients {
		err := client.WriteJSON(response)
		if err != nil {
			log.Println("websocket error")
			log.Println(err)
			_ = client.Close()
			delete(clients, client)
		}
	}
}
