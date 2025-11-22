package main

import (
	"log"
	"net/http"

	"github.com/go-api/db"
	"github.com/go-api/handler"
	"github.com/go-api/model"
)

func main() {
	db := db.Connect()
	db.AutoMigrate(&model.User{})

	userHandler := handler.NewUserHandler(db)

	http.HandleFunc("/health", handler.HealthHandler)
	http.HandleFunc("/users", userHandler.GetUsers)
	http.HandleFunc("/user/create", userHandler.CreateUser)
	http.HandleFunc("/user/get", userHandler.GetUserByID)
	http.HandleFunc("/user/update", userHandler.UpdateUser)
	http.HandleFunc("/user/delete", userHandler.DeleteUser)

	log.Println("Server running on :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
