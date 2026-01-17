package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"sync"
)

type User struct {
	id int 
	name string
	age int 

}


// for thread safty use sync.Mutex or sync.RWMutex

var cacheMutex sync.RWMutex

func main (){

	users := []User{}
	mux := http.NewServeMux()

	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello, World!"))
		fmt.Print("Hello world")
	})

	mux.HandleFunc("POST /users", func(w http.ResponseWriter, r *http.Request) {
		var user User;
		err := json.NewDecoder(r.Body).Decode(&user)

		if err != nil {
			http.Error(w, "Unable to read request body", http.StatusBadRequest)
			return
		}

		if user.name == "" || user.age <= 0 {
			http.Error(w, "Invalid user data", http.StatusBadRequest)
			return
		}

		fmt.Println(user)

		users[len(users) + 1] = user
		w.WriteHeader(http.StatusNoContent)

	})

	fmt.Println("Server is listening at http://localhost:8080")
	http.ListenAndServe(":8080", mux)
}