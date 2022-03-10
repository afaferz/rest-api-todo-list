package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"backend/to-do/db"
	"backend/to-do/models"

	"github.com/gorilla/mux"
)

func Home(wr http.ResponseWriter, req *http.Request) {
	fmt.Fprint(wr, "Homepage")
}

func AllTasks(wr http.ResponseWriter, req *http.Request) {
	var tasks []models.Task
	db.DB.Find(&tasks)
	json.NewEncoder(wr).Encode(tasks)
}

func TaskById(wr http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)
	taskId := vars["id"]

	var task models.Task
	db.DB.First(&task, taskId)
	json.NewEncoder(wr).Encode(task)
}

func CreateNewTask(wr http.ResponseWriter, req *http.Request) {
	var newTask models.Task
	json.NewDecoder(req.Body).Decode(&newTask)
	db.DB.Create(&newTask)
	json.NewEncoder(wr).Encode(newTask)
}

func DeleteTask(wr http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)
	taskId := vars["id"]

	var task models.Task
	db.DB.Delete(&task, taskId)
	json.NewEncoder(wr).Encode(task)
}

func EditTask(wr http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)
	taskId := vars["id"]

	var task models.Task
	db.DB.First(&task, taskId)
	json.NewDecoder(req.Body).Decode(&task)
	db.DB.Save(&task)
	json.NewEncoder(wr).Encode(task)
}
