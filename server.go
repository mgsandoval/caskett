// Basic HTTP server with two routes

package main

import (
	"fmt"
	"errors"
	"io"
	"net/http"
	"os"
)

// Function getRoot: prints a message and returns the output
func getRoot(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("got / request\n")
	io.WriteString(w, "This is my website!!!")
}

// Fucntion getHello: prints a message and returns the output
func getHello(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("got /hello request\n")
	io.WriteString(w, "Hello, HTTP!\n")
}

// mux variable is set with http.NewServeMux 
func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", getRoot)
	mux.HandleFunc("/hello", getHello)
	
	fmt.Printf("Listening on port 3333\n")
	err := http.ListenAndServe(":3333", mux)

	// Server error handling
	if errors.Is(err, http.ErrServerClosed) {
		fmt.Printf("Server closed.\n")
	} else if err != nil {
		fmt.Printf("Error starting server: %s\n", err)
		os.Exit(1)
	}
}
