package main

import (
	"log"
	"net/http"
	"config"
	"models"
)

func	main() {
	config.DatabaseInit()
	router := InitializeRouter()

	log.Fatal(http.ListenAndServe(":8080", router))
}