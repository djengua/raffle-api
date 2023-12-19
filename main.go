package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/djengua/rifa-api/handlers"
)

func main() {
	addr := fmt.Sprintf(":%s", os.Getenv("PORT"))
	if addr == ":" {
		addr = ":8080"
	}
	mux := http.NewServeMux()

	mux.HandleFunc("/hello", handlers.TranslateHandler)
	mux.HandleFunc("/health", handlers.HealthCheck)
	log.Printf("listening on %s\n", addr)
	log.Fatal(http.ListenAndServe(addr, mux))
}

type Resp struct {
	Language    string `json:"language"`
	Translation string `json:"translation"`
}
