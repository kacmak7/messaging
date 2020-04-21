package main 

import (
	//"flag"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	
)

func main() {
	//port := flag.String("port", "")

	initializeNode()
	send("HI HELLOooo")
	
	//
	
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/test", test).Methods("GET")
	//router.HandleFunc("/event", createEvent).Methods("POST")
	log.Fatal(http.ListenAndServe(":8080", router))

	//

}
