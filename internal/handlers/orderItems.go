package handlers

import (
	"ecommerce-api/internal/db"
	"ecommerce-api/internal/models"
	"encoding/json"
	"net/http"
)

func CreateItems(w http.ResponseWriter, r *http.Request) {
	var item models.OrderItem
	if err := json.NewDecoder(r.Body).Decode(&item); err != nil {
		http.Error(w, "invlaid body formate", 400)
		return
	}
	var currentOrder models.Order
	if err := db.DB.First(&currentOrder, item.OrderID).Error; err != nil {
		http.Error(w, "order id not exist", 400)
		return
	}
	if err := db.DB.Create(&item).Error; err != nil {
		http.Error(w, "internal error", 400)
		return
	}
	db.DB.Preload("Order").First(&item, item.ID)
	w.Header().Set("content-type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(item)
}
