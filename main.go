package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type Message struct {
	Name    string `json:"name"`
	Content string `json:"content"`
}

func getHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}
	fmt.Fprintf(w, "Welcome to sachin Api")

}

func postHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}

	var msg Message

	body, err := io.ReadAll(r.Body) // Read body of the POST request
	if err != nil {

		http.Error(w, "Invalid read request body", http.StatusBadRequest)
		return
	}

	// Parse the JSON
	err = json.Unmarshal(body, &msg)
	if err != nil {
		http.Error(w, "Invalid JSON format", http.StatusBadRequest)
		return
	}
	// Response back with the same MSG
	w.Header().Set("Content/Type", "application/json")
	json.NewEncoder(w).Encode(msg)

}
func main() {
	http.HandleFunc("/get", getHandler)
	http.HandleFunc("/post", postHandler)
	fmt.Println("Server running")
	http.ListenAndServe(":8090", nil)

}
