package handlers

import (
	"ecommerce-api/internal/db"
	"ecommerce-api/internal/models"
	"encoding/json"
	"strconv"

	"net/http"
)

func CreateUser(w http.ResponseWriter, r *http.Request) {
	var newuser models.User
	if err := json.NewDecoder(r.Body).Decode(&newuser); err != nil {
		http.Error(w, "invalid request body", http.StatusBadRequest)
		return
	}
	if err := db.DB.Create(&newuser).Error; err != nil {
		http.Error(w, "User already exist", http.StatusConflict)
		return
	}
	w.Header().Set("content-type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(newuser)
}
func GetUsers(w http.ResponseWriter, r *http.Request) {
	strID := r.URL.Query().Get("user_id")
	if strID == "" {
		http.Error(w, "id is empty", http.StatusBadRequest)
		return
	}
	id, err := strconv.ParseUint(strID, 10, 32)
	if err != nil {
		http.Error(w, "invalid user id", http.StatusBadRequest)
		return
	}
	var user models.User
	if err := db.DB.Where("id = ?", id).Find(&user).Error; err != nil {
		http.Error(w, "user not exist", http.StatusBadRequest)
		return
	}
	db.DB.Preload("Orders").First(&user, id)
	w.Header().Set("content-type", "application/json")
	json.NewEncoder(w).Encode(user)

}
