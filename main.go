package main

import (
	"fmt"
	"log"
	"net/http"
	"text/template"

	"frontendmasters.com/go/museum/api"
	"frontendmasters.com/go/museum/data"
)

// Port the server is running on
var port = ":3333"

func handleHello(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello from a Go program"))
}

func handleTemplate(w http.ResponseWriter, r *http.Request) {
	html, err := template.ParseFiles("templates/index.tmpl")
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Internal Server Error"))
		return
	}

	html.Execute(w, data.GetAll())
}

func main() {
	// Create the server
	server := http.NewServeMux()
	server.HandleFunc("/hello", handleHello)
	server.HandleFunc("/template", handleTemplate)
	server.HandleFunc("/api/exhibitions", api.GetExhibitions)
	server.HandleFunc("/api/exhibitions/new", api.Post)

	// Set up the public folder
	fs := http.FileServer(http.Dir("./public"))
	server.Handle("/", fs)

	// Set up the error listening
	log.Printf("Server starting at http://localhost%v/", port)
	err := http.ListenAndServe(port, server)

	if err == nil {
		fmt.Printf("Error while starting the server %v", err)
	}
}
