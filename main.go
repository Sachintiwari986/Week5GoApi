package main

import (
	"fmt"
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
func main() {
	http.HandleFunc("/get", getHandler)
	fmt.Println("Server running")
	http.ListenAndServe(":8090", nil)

}
