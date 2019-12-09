package controllers

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"strconv"
	"github.com/mateogreil/hello-world-42go/models"
)

func ProjectsIndex(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json;charset=UTF-8")
	w.WriteHeader(http.StatusOK)

	pageParam, ok := r.URL.Query()["page"]

	var page int
	if ok {
		var err error
		page, err = strconv.Atoi(pageParam[0])
		if err != nil {
			log.Fatal(err)
		}
	}

	projects := models.AllProject(page)
	json.NewEncoder(w).Encode(projects)
}

func ProjectsShow(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json;charset=UTF-8")
	w.WriteHeader(http.StatusOK)

	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])

	if err != nil {
		log.Fatal(err)
	}

	project := models.FindProjectById(id)

	json.NewEncoder(w).Encode(project)
}