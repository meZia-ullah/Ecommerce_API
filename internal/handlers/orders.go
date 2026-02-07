package handlers

import (
	"ecommerce-api/internal/db"
	"ecommerce-api/internal/models"
	"encoding/json"
	"net/http"
)

func CreateOrder(w http.ResponseWriter, r *http.Request) {
	var neworder models.Order
	if err := json.NewDecoder(r.Body).Decode(&neworder); err != nil {
		http.Error(w, "invalid request body", http.StatusBadRequest)
		return
	}
	var user models.User
	if err := db.DB.First(&user, neworder.UserID).Error; err != nil {
		http.Error(w, "user not found", http.StatusBadRequest)
		return
	}
	if err := db.DB.Create(&neworder).Error; err != nil {
		http.Error(w, "Order could not be created", http.StatusConflict)
		return
	}
	db.DB.Preload("User").First(&neworder, neworder.ID)
	w.Header().Set("content-type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(neworder)
}
