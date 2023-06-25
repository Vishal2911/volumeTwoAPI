package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type User struct {
	Name    string
	Age     int
	Address string
}

var (
	userdata = make(map[string]User)
)

func main() {

	http.HandleFunc("/users", getusers)
	http.HandleFunc("/createuser",adduser)
	fmt.Println("Server started and running at 8080")
	log.Fatal(http.ListenAndServe(":8080", nil))

}

func adduser(w http.ResponseWriter, r *http.Request) {

	if r.Method != "POST" {
		w.WriteHeader(http.StatusMethodNotAllowed)
	}

	var user User

	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	userdata[user.Name] = user

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(user)

}

func getusers(w http.ResponseWriter, r *http.Request) {

	if r.Method != "GET" {
		w.WriteHeader(http.StatusMethodNotAllowed)
	}

	w.WriteHeader(http.StatusAccepted)
	json.NewEncoder(w).Encode(userdata)

}
