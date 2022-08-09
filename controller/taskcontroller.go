package controller

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"taskwithgo/database"
	"taskwithgo/entity"

	"github.com/gorilla/mux"
)

func GetAllTask(w http.ResponseWriter, r *http.Request) {
	var tasks []entity.Task
	database.Connector.Find(&tasks)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(tasks)
}

func GetTaskById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	key := vars["id"]

	var task entity.Task
	database.Connector.First(&task, key)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(task)
}

func CreateTask(w http.ResponseWriter, r *http.Request) {
	requestBody, _ := ioutil.ReadAll(r.Body)
	var task entity.Task
	json.Unmarshal(requestBody, &task)

	database.Connector.Create(task)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(task)
}