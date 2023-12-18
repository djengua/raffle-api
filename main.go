package main

import (
	"log"
	"net/http"

	"github.com/djengua/rifa-api/handlers"
)

func main() {
	addr := ":8080"
	mux := http.NewServeMux()

	mux.HandleFunc("/hello", handlers.TranslateHandler)
	log.Printf("listening on %s\n", addr)
	log.Fatal(http.ListenAndServe(addr, mux))
}

type Resp struct {
	Language    string `json:"language"`
	Translation string `json:"translation"`
}
