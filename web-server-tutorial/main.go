package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"sync"
)

type User struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
	Age  int    `json:"age"`
}

// for thread safty use sync.Mutex or sync.RWMutex

var cacheMutex sync.RWMutex

func main() {

	users := make(map[int]User)
	mux := http.NewServeMux()

	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello, World!"))
		fmt.Print("Hello world")
	})

	mux.HandleFunc("POST /users", func(w http.ResponseWriter, r *http.Request) {
		var user User
		err := json.NewDecoder(r.Body).Decode(&user)

		if err != nil {
			http.Error(w, "Unable to read request body", http.StatusBadRequest)
			return
		}

		if user.Name == "" || user.Age <= 0 {
			http.Error(w, "Invalid user data", http.StatusBadRequest)
			fmt.Println(user)
			return
		}

		fmt.Println(user)

		cacheMutex.Lock()
		newID := len(users) + 1
		user.ID = newID
		users[newID] = user

		cacheMutex.Unlock()
		w.WriteHeader(http.StatusNoContent)
	})

	mux.HandleFunc("/users/{id}", func(w http.ResponseWriter, r *http.Request) {
		cacheMutex.RLock()

		id, err := strconv.Atoi(r.URL.Path[len("/users/"):])
		if err != nil {
			http.Error(w, "Invalid user ID", http.StatusBadRequest)
			return
		}
		defer cacheMutex.RUnlock()

		user, ok := users[id]
		if !ok {
			http.Error(w, "User not found", http.StatusNotFound)
			return
		}
		j, err := json.Marshal(user)

		if err != nil {
			http.Error(w, "Unable to marshal user data", http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		w.Write(j)
	})

	mux.HandleFunc("DELETE /users/{id}", func(w http.ResponseWriter, r *http.Request) {
		id, err := strconv.Atoi(r.URL.Path[len("/users/"):])
		if err != nil {
			http.Error(w, "Invalid user ID", http.StatusBadRequest)
			return
		}

		_, ok := users[id]
		if !ok {
			http.Error(w, "User not found", http.StatusNotFound)
			return
		}
		cacheMutex.Lock()
		delete(users, id)
		cacheMutex.Unlock()
		w.WriteHeader(http.StatusNoContent)

	})

	fmt.Println("Server is listening at http://localhost:8080")
	http.ListenAndServe(":8080", mux)
}