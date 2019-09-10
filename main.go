package main

import (
	"log"
	"net/http"
)

// Define a home handler function which writes a byte slice containing
// "Hello from Snippetbox" as the response body.
func home(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}
	w.Write([]byte("Hello from Snippetbox"))
}

func showSnippet(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Display a specific snippet..."))
}

func createSnippet(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Create a new snippet..."))
}

func main() {
	// Use the http.NewServeMux() function to initialize a new servemux, then
	// register the home function as the handler for the "/" URL pattern.
	mux := http.NewServeMux()
	mux.HandleFunc("/", home)
	mux.HandleFunc("/snippet", showSnippet)	
	mux.HandleFunc("/snippet/create", createSnippet)

	// http.ListenAndServe() -> start a new web server
	// pass in 2 params: TCP address to listen on and servemux.
	// error logged by log.Fatal()
	log.Println("Starting server on :4000")
	err := http.ListenAndServe(":4000", mux)
	log.Fatal(err)
}

