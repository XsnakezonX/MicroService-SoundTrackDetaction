package main

import (
	"log"
	"net/http"
	"tracks/repository"
	"tracks/resources"
)

func main() {
	repository.Init()
	repository.Clear()
	repository.Create()
	log.Fatal(http.ListenAndServe(":3000", resources.Router()))
}
