package controller

import (
	"encoding/json"
	"html/template"
	"io/ioutil"
	"net/http"
	"strconv"
	"time"

	"github.com/rockaflare/taskwithgo/database"
	"github.com/rockaflare/taskwithgo/entity"

	"github.com/gorilla/mux"
)

var (
	id       int
	pegawai  string
	detail   string
	deadline time.Time
	IsDone   bool

	view = template.Must(template.ParseFiles("./pages/Index.html"))
)

func Index(w http.ResponseWriter, r *http.Request) {
	var tasks []entity.Task
	database.Connector.Find(&tasks)

	data := entity.View{
		Tasks: tasks,
	}

	_ = view.Execute(w, data)
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

func UpdateTask(w http.ResponseWriter, r *http.Request) {
	requestBody, _ := ioutil.ReadAll(r.Body)
	var task entity.Task
	json.Unmarshal(requestBody, &task)
	database.Connector.Save(&task)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(task)
}

func DeleteTask(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	key := vars["id"]

	var task entity.Task
	id, _ := strconv.ParseInt(key, 10, 64)
	database.Connector.Where("id = ?", id).Delete(&task)
	w.WriteHeader(http.StatusNoContent)
}

func SetDone(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	key := vars["id"]

	id, _ := strconv.ParseInt(key, 10, 64)
	database.Connector.Where("id = ?", id).Update("isdone", 1)
	w.WriteHeader(http.StatusNoContent)
}
