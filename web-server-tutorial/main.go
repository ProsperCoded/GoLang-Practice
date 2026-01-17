package main

import (
	"fmt"
	"net/http"
)


func main (){
	mux := http.NewServeMux()

	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello, World!"))
		fmt.Print("Hello world")
	})

	fmt.Println("Server is listening at http://localhost:8080")
	http.ListenAndServe(":8080", mux)
}