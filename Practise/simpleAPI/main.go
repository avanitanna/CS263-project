package main

import (
	"encoding/json" // you need this to return as json
	"fmt"
	"log"      // you need this to logging if there's any error
	"net/http" // you need this to interact with server
)

func User(w http.ResponseWriter, r *http.Request) {
	// Set header response to "Content-Type application-json"
	w.Header().Set("Content-Type", "application/json")

	// Make some fake data
	user := map[string]interface{}{"ID": 1, "Name": "Huda Prasetyo"}

	// Send the response to client
	json.NewEncoder(w).Encode(user)
}

func main() {
	// Make a route
	http.HandleFunc("/api/user", User)

	// Some text to be appear for connection
	fmt.Println("Server started at port 8000")

	// Run the server in port 8000
	log.Fatal(http.ListenAndServe(":8000", nil))
}
