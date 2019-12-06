package main

import (
	"log"
	"net/http"
	"github.com/mateogreil/hello-world-42go/config"
)

func	main() {
	config.DatabaseInit()
	router := InitializeRouter()

	log.Fatal(http.ListenAndServe(":8080", router))
}