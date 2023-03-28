package main

import (
	"cooltown/resources"
	"log"
	"net/http"
)

func main() {
	log.Fatal(http.ListenAndServe(":3002", resources.Router()))
}
