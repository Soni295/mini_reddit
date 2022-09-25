package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

const (
	PORT = ":8080"
	HOST = "localhost"
)

func main() {
	router := mux.NewRouter()

	router.HandleFunc("/", func(w http.ResponseWriter, _ *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("It's works"))
	})

	router.HandleFunc("/users", getUsers).Methods(http.MethodGet)
	router.HandleFunc("/users", addUser).Methods(http.MethodPost)

	fmt.Printf("Server listening in the port%s\n\n", PORT)
	fmt.Printf("Link: http://%s%s\n", HOST, PORT)

	if err := http.ListenAndServe(PORT, router); err != nil {
		log.Fatalf("Server error: %s", err.Error())
	}
}
