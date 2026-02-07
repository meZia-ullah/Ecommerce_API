package routes

import (
	"ecommerce-api/internal/handlers"
	"net/http"
)

func RegisterRoutes() {
	http.HandleFunc("/user", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodPost {
			handlers.CreateUser(w, r)
		} else if r.Method == http.MethodGet {
			handlers.GetUsers(w, r)
		}
	})
	http.HandleFunc("/order", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodPost {
			handlers.CreateOrder(w, r)
		}
	})
}
