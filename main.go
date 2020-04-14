package main 

import (
	"fmt"
	"log"
	"net/http"
	"github.com/gorilla/mux"
	
	"github.com/kacmak7/messaging"
)


func main() {
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/test", test).Methods("GET")
	//router.HandleFunc("/event", createEvent).Methods("POST")
	log.Fatal(http.ListenAndServe(":8080", router))
}