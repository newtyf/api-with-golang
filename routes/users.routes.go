package routes

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/newtyf/go-gorm-restapi/db"
	"github.com/newtyf/go-gorm-restapi/models"
)

// TODO: get all users
func GetUsersHandler(w http.ResponseWriter, r *http.Request) {
	var users []models.User

	db.DB.Find(&users)

	json.NewEncoder(w).Encode(&users)
}

// TODO: get detail user
func GetUserHandler(w http.ResponseWriter, r *http.Request) {
	var user models.User

	userId := (mux.Vars(r))["id"]

	db.DB.First(&user, userId)

	if user.ID == 0 {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("User not found"))
		return
	}

	db.DB.Model(&user).Association("Tasks").Find(&user.Tasks)

	json.NewEncoder(w).Encode(&user)
}

// TODO: create user
func PostUsersHandler(w http.ResponseWriter, r *http.Request) {

	var user models.User

	json.NewDecoder(r.Body).Decode(&user)

	createdUser := db.DB.Create(&user)

	if err := createdUser.Error; err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		return
	}

	json.NewEncoder(w).Encode(&user)
}

// TODO: delete user
func DeleteUsersHandler(w http.ResponseWriter, r *http.Request) {

	var user models.User

	userId := (mux.Vars(r))["id"]
	db.DB.First(&user, userId)

	if user.ID == 0 {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("User not found"))
		return
	}

	db.DB.Delete(&user) //? soft delete not really delete of db
	// db.DB.Unscoped().Delete(&user) //? real delete yes really delete of db
	w.Write([]byte("User deleted"))
}
