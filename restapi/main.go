package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

// prints message when called from log.Fatal(...).
func logRequest(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(writer http.ResponseWriter, rout *http.Request) {
		log.Println("Data received from", rout.RemoteAddr)
		handler.ServeHTTP(writer, rout)
	})
}

func initializeRouter() {
	rout := mux.NewRouter()

	rout.HandleFunc("/users", GetUsers).Methods("GET")
	rout.HandleFunc("/users/{id}", GetUser).Methods("GET")
	rout.HandleFunc("/users", CreateUser).Methods("POST")
	rout.HandleFunc("/users/{id}", UpdateUser).Methods("PUT")
	rout.HandleFunc("/users/{id}", DeleteUser).Methods("DELETE")

	rout.HandleFunc("/stocks", GetStocks).Methods("GET")
	rout.HandleFunc("/stocks/{tick}", GetStock).Methods("GET")

	log.Fatal(http.ListenAndServe(":9000", logRequest(rout)))

}

func main() {
	InitialMigration()
	initializeRouter()
}
