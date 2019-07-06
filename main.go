package main

import (
	"log"
	"needyd/routing"
	"net/http"
)

func main() {
	handler := routing.NewRouter()
	log.Println("Listening on port 8080")
	log.Fatal(http.ListenAndServe(":8080", handler))
}
