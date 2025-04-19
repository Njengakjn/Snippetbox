package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
)

// handler for writing response and executing application logic
func home(w http.ResponseWriter, r *http.Request) {
	// create a fixed path home page
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}

	w.Write([]byte("Hello, sailor"))
}

func snippetboxView(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil || id < 1 {
		http.NotFound(w, r)
		return
	}
	//data := fmt.Sprintf("Displaying the snippet for id: %d", id)
	//w.Write([]byte(data))
	fmt.Fprintf(w, "Displaying the snippet for id: %d", id)
}

func snippetboxCreate(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		w.Header().Set("Allow", http.MethodPost)
		//w.WriteHeader(405)
		//w.Write([]byte("METHOD NOT ALLOWED"))

		// line below combines the write header and write functions
		http.Error(w, "METHOD NOT ALLOWED", http.StatusMethodNotAllowed)
		return
	}
	w.Write([]byte("Creating a new snippet"))
}

func main() {
	mux := http.NewServeMux()

	// routes
	mux.HandleFunc("/", home)
	mux.HandleFunc("/snippet/view", snippetboxView)
	mux.HandleFunc("/snippet/create", snippetboxCreate)

	// star the server
	log.Println("Starting the server")
	err := http.ListenAndServe(":4000", mux)
	log.Fatal(err)
}
