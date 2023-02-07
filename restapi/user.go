package main

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	FirstName string `json:"firstname"`
	LastName  string `json:"lastname"`
	Email     string `json:"email"`
	Balance   int    `json:"balance"`

	//stock positions, etc.
}

func GetUsers(writer http.ResponseWriter, rout *http.Request) {
	writer.Header().Set("Content-Type", "application/json")
	var users []User
	//gets table info from db
	DB.Find(&users)
	//sends information
	json.NewEncoder(writer).Encode(users)
}

func GetUser(writer http.ResponseWriter, rout *http.Request) {
	writer.Header().Set("Content-Type", "application/json")
	params := mux.Vars(rout)
	var user []User
	DB.First(&user, params["id"])
	json.NewEncoder(writer).Encode(user)
}

func CreateUser(writer http.ResponseWriter, rout *http.Request) {
	writer.Header().Set("Content-Type", "application/json")
	var user User
	//handles data received from request (json data)
	json.NewDecoder(rout.Body).Decode(&user)
	DB.Create(&user)
	//sends info to browser
	json.NewEncoder(writer).Encode(user)
}

func UpdateUser(writer http.ResponseWriter, rout *http.Request) {
	writer.Header().Set("Content-Type", "application/json")
	params := mux.Vars(rout)
	var user User
	DB.First(&user, params["id"])
	json.NewDecoder(rout.Body).Decode(&user)
	DB.Save(&user)
	json.NewEncoder(writer).Encode(user)
}

func DeleteUser(writer http.ResponseWriter, rout *http.Request) {
	writer.Header().Set("Content-Type", "application/json")
	params := mux.Vars(rout)
	var user []User
	DB.Delete(&user, params["id"])
	json.NewEncoder(writer).Encode("User deleted.")
}

//func - update prices if possible when starting application
