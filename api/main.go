package main

import (
	"api/routes"
	"fmt"
	"log"
	"net/http"
)

var hostName = "localhost"
var serverPort = "8080"

func main() {
	fmt.Println("Server listening on port 8080...")

	// Initialize routes
	routes.InitRoutes()

	err := http.ListenAndServe(hostName+":"+serverPort, nil)

	if err != nil {
		log.Fatal(err)
	}
}

//import (
//	"encoding/json"
//	"fmt"
//	"log"
//	"net/http"
//)
//
//type User struct {
//	Id        int64
//	Username  string
//	Password  string
//	CreatedAt string
//	UpdatedAt string
//}
//
//var users []User
//
//var hostName = "localhost"
//var serverPort = "8080"
//
//func main() {
//	http.HandleFunc("/users", UserController)
//	http.HandleFunc("/heath-check", HealthCheckController)
//
//	err := http.ListenAndServe(hostName+":"+serverPort, nil)
//
//	if err != nil {
//		log.Fatal(err)
//	}
//}
//
//func UserController(w http.ResponseWriter, r *http.Request) {
//	switch r.Method {
//	case http.MethodGet:
//		getUsers(w, r)
//	case http.MethodPost:
//		UpdateUser(w, r)
//	default:
//		http.Error(w, "Invalid request", http.StatusMethodNotAllowed)
//	}
//}
//
//func getUsers(w http.ResponseWriter, r *http.Request) {
//	json.NewEncoder(w).Encode(users)
//
//	fmt.Fprintf(w, "Get users: '%v'", users)
//}
//
//func UpdateUser(w http.ResponseWriter, r *http.Request) {
//	var user User
//	err := json.NewDecoder(r.Body).Decode(&user)
//
//	if err != nil {
//		http.Error(w, err.Error(), http.StatusInternalServerError)
//	}
//
//	users = append(users, user)
//
//	fmt.Fprintf(w, "Create an user: '%v'", user)
//}
//
//func HealthCheckController(w http.ResponseWriter, r *http.Request) {
//	fmt.Fprintf(w, "Everything is okay")
//}
