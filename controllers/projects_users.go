package controllers

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"strconv"
	"github.com/mateogreil/hello-world-42go/models"
)

func ProjectsUsersShow(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json;charset=UTF-8")
	w.WriteHeader(http.StatusOK)

	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])

	if err != nil {
		log.Fatal(err)
	}

	projects_user := models.FindProjectsUserById(id)

	json.NewEncoder(w).Encode(projects_user)
}

func ProjectsUsersShowByUserId(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json;charset=UTF-8")
	w.WriteHeader(http.StatusOK)

	vars := mux.Vars(r)
	user_id, err := strconv.Atoi(vars["user_id"])

	if err != nil {
		log.Fatal(err)
	}

	projects_user := models.FindProjectsUserByUserId(user_id)

	json.NewEncoder(w).Encode(projects_user)
}
