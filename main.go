package main

import (
	"log"
	"net/http"
)

func home(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}
	w.Write([]byte("this is the home page"))
}

func showSnippet(w http.ResponseWriter, r *http.Request)  {
	w.Write([]byte("showing the snippet"))
}

func createSnippet(w http.ResponseWriter, r * http.Request)  {
	w.Write([]byte("creating the snippet"))
}
func main() {

	mux := http.NewServeMux()
	mux.HandleFunc("/", home)
	mux.HandleFunc("/snippet", showSnippet)
	mux.HandleFunc("/snippet/create", createSnippet)

	log.Println("Starting server on :4000")
	err := http.ListenAndServe(":4000", mux)
	log.Fatal(err)
}