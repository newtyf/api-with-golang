package routes

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/newtyf/go-gorm-restapi/db"
	"github.com/newtyf/go-gorm-restapi/models"
)

// TODO: get all tasks
func GetTasksHandler(w http.ResponseWriter, r *http.Request) {
	var tasks []models.Task

	db.DB.Find(&tasks)
	json.NewEncoder(w).Encode(&tasks)
}

// TODO: get detail task
func GetTaskHandler(w http.ResponseWriter, r *http.Request) {
	var task models.Task

	params := mux.Vars(r)

	db.DB.First(&task, params["id"])

	if task.ID == 0 {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("Task not found"))
		return
	}

	json.NewEncoder(w).Encode(&task)
}

// TODO: create task
func CreateTaskHandler(w http.ResponseWriter, r *http.Request) {

	var task models.Task

	json.NewDecoder(r.Body).Decode(&task)

	createdTask := db.DB.Create(&task)

	if err := createdTask.Error; err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		return
	}

	json.NewEncoder(w).Encode(&task)
}

// TODO: delete task
func DeleteTaskHandler(w http.ResponseWriter, r *http.Request) {

	var task models.Task

	params := mux.Vars(r)
	db.DB.First(&task, params["id"])

	if task.ID == 0 {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Task not found"))
		return
	}

	// db.DB.Delete(&task) //? soft delete not really delete of db
	db.DB.Unscoped().Delete(&task) //? real delete yes really delete of db
	w.WriteHeader(http.StatusNoContent)
	// w.Write([]byte("Task deleted")) //not work because header is status nocontent
}
