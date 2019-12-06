package main

import (
	"github.com/gorilla/mux"
	"github.com/mateogreil/hello-world-42go/controllers"
)

func InitializeRouter() *mux.Router {
	router := mux.NewRouter().StrictSlash(true)

	router.Methods("GET").Path("/users/{id}").Name("Show").HandlerFunc(controllers.UsersShow)
	return router
}