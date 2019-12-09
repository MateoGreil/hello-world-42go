package main

import (
	"github.com/gorilla/mux"
	"github.com/mateogreil/hello-world-42go/controllers"
)

func InitializeRouter() *mux.Router {
	router := mux.NewRouter().StrictSlash(true)

	// User
	router.Methods("GET").Path("/users/{id}").Name("Show").HandlerFunc(controllers.UsersShow)
	router.Methods("GET").Path("/users/login/{login}").Name("Show").HandlerFunc(controllers.UsersShowByLogin)
	
	// Project
	router.Methods("GET").Path("/projects/{id}").Name("Show").HandlerFunc(controllers.ProjectsShow)
	
	// ProjectsUser
	router.Methods("GET").Path("/projects_users/{id}").Name("Show").HandlerFunc(controllers.ProjectsUsersShow)
	router.Methods("GET").Path("/projects_users/user_id/{user_id}").Name("Show").HandlerFunc(controllers.ProjectsUsersShowByUserId)

	return router
}