package main

import (
	"ecommerce-api/internal/db"
	"ecommerce-api/internal/models"
	"ecommerce-api/internal/routes"
	"log"
	"net/http"

	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(".env"); err != nil {
		log.Print("error while loading .env", err)
	}
	if err := db.Connect(); err != nil {
		log.Fatal("Failed to connect to Database", err)
	}
	db.DB.AutoMigrate(
		&models.Order{},
		&models.User{},
		&models.OrderItem{},
	)
	routes.RegisterRoutes()
	log.Print("database connected")
	log.Print("server started at :8082")
	if err := http.ListenAndServe(":8082", nil); err != nil {
		log.Fatal("Failed to start server", err)
	}
}
