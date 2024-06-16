package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	// Serve static files from the "static" directory
	fileServer := http.FileServer(http.Dir("./static"))
	http.Handle("/", fileServer)

	// Handle /form and /hello endpoints
	http.HandleFunc("/form", formHandler)
	http.HandleFunc("/hello", helloHandler)

	// Start the server on port 8080
	fmt.Println("Starting server at port 8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatalf("Could not start server: %s\n", err)
	}
}

// helloHandler handles requests to the /hello endpoint
func helloHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/hello" {
		http.Error(w, "404 not found", http.StatusNotFound)
		return
	}
	if r.Method != "GET" {
		http.Error(w, "Method not supported", http.StatusMethodNotAllowed)
		return
	}
	if _, err := fmt.Fprint(w, "hello!"); err != nil {
		http.Error(w, "Failed to write response", http.StatusInternalServerError)
	}
}

// formHandler handles POST requests to the /form endpoint
func formHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not supported", http.StatusMethodNotAllowed)
		return
	}
	if err := r.ParseForm(); err != nil {
		http.Error(w, fmt.Sprintf("ParseForm() error: %v", err), http.StatusBadRequest)
		return
	}
	name := r.FormValue("name")
	address := r.FormValue("address")
	response := fmt.Sprintf("POST request successful\nName = %s\nAddress = %s\n", name, address)
	if _, err := fmt.Fprint(w, response); err != nil {
		http.Error(w, "Failed to write response", http.StatusInternalServerError)
	}
}
