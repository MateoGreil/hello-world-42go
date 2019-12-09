package main

import (
	"github.com/gorilla/mux"
	"github.com/mateogreil/hello-world-42go/controllers"
)

func InitializeRouter() *mux.Router {
	router := mux.NewRouter().StrictSlash(true)

	// User
	router.Methods("GET").Path("/users").Name("Index").HandlerFunc(controllers.UsersIndex)
	router.Methods("GET").Path("/users/{id}").Name("Show").HandlerFunc(controllers.UsersShow)
	
	// Project
	router.Methods("GET").Path("/projects").Name("Index").HandlerFunc(controllers.ProjectsIndex)
	router.Methods("GET").Path("/projects/{id}").Name("Show").HandlerFunc(controllers.ProjectsShow)
	
	// ProjectsUser
	router.Methods("GET").Path("/projects_users/{id}").Name("Show").HandlerFunc(controllers.ProjectsUsersShow)

	return router
}