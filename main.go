package main

import (
	"log"
	"net/http"
	"github.com/mateogreil/hello-world-42go/config"
	"github.com/mateogreil/hello-world-42go/models"
)

func	main() {
	config.DatabaseInit()
	router := InitializeRouter()

	log.Fatal(http.ListenAndServe(":8080", router))
}