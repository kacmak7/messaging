package main

import (
	//"flag"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func launchServer() {
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/ping", ping).Methods("GET")
	//router.HandleFunc("/event", createEvent).Methods("POST")
	log.Print(http.ListenAndServe(":8080", router))
}

// TODO expose to the world
