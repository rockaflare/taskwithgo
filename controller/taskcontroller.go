package controller

import (
	"encoding/json"
	"net/http"

	"github.com/rockaflare/taskwithgo/database"
	"github.com/rockaflare/taskwithgo/entity"

	"github.com/gorilla/mux"
)

func Index(w http.ResponseWriter, r *http.Request) {
	var tasks []entity.Task
	database.Connector.Find(&tasks)
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(tasks)
}

func GetTaskById(w http.ResponseWriter, r *http.Request) {
	taskId := mux.Vars(r)["id"]
	if checkIfProductExists(taskId) == false {
		json.NewEncoder(w).Encode("Product Not Found!")
		return
	}
	var task entity.Task
	database.Connector.First(&task, taskId)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(task)
}

func CreateTask(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "application/json")
	var task entity.Task
	json.NewDecoder(r.Body).Decode(&task)
	database.Connector.Create(&task)
	json.NewEncoder(w).Encode(task)
}

func UpdateTask(w http.ResponseWriter, r *http.Request) {
	taskId := mux.Vars(r)["id"]
	if checkIfProductExists(taskId) == false {
		json.NewEncoder(w).Encode("Task tidak ditemukan!")
		return
	}
	var task entity.Task
	database.Connector.First(&task, taskId)
	json.NewDecoder(r.Body).Decode(&task)
	database.Connector.Save(&task)
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(task)
}

func DeleteTask(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "application/json")
	taskId := mux.Vars(r)["id"]
	if checkIfProductExists(taskId) == false {
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode("Task tidak ditemukan!")
		return
	}
	var task entity.Task
	database.Connector.Delete(&task, taskId)
	json.NewEncoder(w).Encode("Task dihapus!")
}

func checkIfProductExists(productId string) bool {
	var task entity.Task
	database.Connector.First(&task, productId)
	if task.ID == 0 {
		return false
	}
	return true
}

func SetDone(w http.ResponseWriter, r *http.Request) {
	taskId := mux.Vars(r)["id"]
	if checkIfProductExists(taskId) == false {
		json.NewEncoder(w).Encode("Task tidak ditemukan!")
		return
	}
	var task entity.Task
	database.Connector.First(&task, taskId)
	task.IsDone = true
	database.Connector.Save(&task)
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(task)
}

func UnDone(w http.ResponseWriter, r *http.Request) {
	taskId := mux.Vars(r)["id"]
	if checkIfProductExists(taskId) == false {
		json.NewEncoder(w).Encode("Task tidak ditemukan!")
		return
	}
	var task entity.Task
	database.Connector.First(&task, taskId)
	task.IsDone = false
	database.Connector.Save(&task)
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(task)
}
