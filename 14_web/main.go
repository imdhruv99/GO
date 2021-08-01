package main

import (
	"fmt"
	"net/http"
)

func index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello Dhruv Prajapati")
}

func main() {

	// router
	http.HandleFunc("/", index)
	fmt.Println("Server is Running...")
	// Exposing
	http.ListenAndServe(":3030", nil)
}