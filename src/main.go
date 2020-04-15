package main 

import (
	"log"
	"net/http"
	"github.com/gorilla/mux"
	
)

func main() {
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/test", test).Methods("GET")
	//router.HandleFunc("/event", createEvent).Methods("POST")
	log.Fatal(http.ListenAndServe(":8080", router))
}