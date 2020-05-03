package main

import (
	//"flag"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func launchServer() { // TODO expose to the world
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/ping", pong).Methods("GET")
	//router.HandleFunc("/event", createEvent).Methods("POST")
	log.Print(http.ListenAndServe(":8080", router))
}

func pong(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("PONG\n"))
}

func authorize(w http.ResponseWriter, r *http.Request) {
	log.Print("authorizing new Node")
}
