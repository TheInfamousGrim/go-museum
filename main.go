package main

import (
	"fmt"
	"log"
	"net/http"
)

// Port the server is running on
var port = ":3333"

func handleHello(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello from a Go program"))
}

func handleStartup(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Server running at http://localhost%v/", port)
}

func main() {
	// Create the server
	server := http.NewServeMux()
	server.HandleFunc("/hello", handleHello)

	// Set up the public folder
	fs := http.FileServer(http.Dir("./public"))
	server.Handle("/", fs)

	// Set up the error listening
	log.Printf("Server starting at http://localhost%v/", port)
	err := http.ListenAndServe(port, server)
	server.HandleFunc("/", handleStartup)
	log.Println("Hello")

	if err == nil {
		fmt.Printf("Error while starting the server %v", err)
	}
}
