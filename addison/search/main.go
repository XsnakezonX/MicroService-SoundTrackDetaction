package main

import (
	"log"
	"net/http"
	"search/resources"
)

func main() {
	log.Fatal(http.ListenAndServe(":3001", resources.Router()))
}
