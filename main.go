package main

import (
	"log"
	"net/http"

	"github.com/djengua/rifa-api/faas"
)

func main() {
	addr := ":8080"
	mux := http.NewServeMux()

	mux.HandleFunc("/hello", faas.Translate)
	log.Printf("listening on %s\n", addr)
	log.Fatal(http.ListenAndServe(addr, mux))
}

type Resp struct {
	Language    string `json:"language"`
	Translation string `json:"translation"`
}
