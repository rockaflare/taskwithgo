package main

import (
	"log"
	"net/http"

	"github.com/rockaflare/taskwithgo/controller"
	"github.com/rockaflare/taskwithgo/database"
	"github.com/rockaflare/taskwithgo/entity"

	"github.com/gorilla/mux"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func main() {
	dbStart()
	log.Println("Menjalankan HTTP server di port 8080")
	router := mux.NewRouter().StrictSlash(true)
	handlerStart(router)
	log.Fatal(http.ListenAndServe(":8080", router))
}

func handlerStart(router *mux.Router) {
	router.HandleFunc("/index", controller.Index)
	router.HandleFunc("/get/{id}", controller.GetTaskById).Methods("GET")
}

func dbStart() {
	config := database.Config{
		ServerName: "localhost:3306",
		User:       "phpmyadmin",
		Password:   "myadminroot",
		DB:         "gotaskapp",
	}

	connectionString := database.GetConnectionString(config)
	err := database.Connect(connectionString)
	if err != nil {
		panic(err.Error())
	}

	database.Migrate(&entity.Task{})
}
