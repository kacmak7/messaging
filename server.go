package main

import (
	//"flag"
	"log"
	"net/http"
	"strings"

	badger "github.com/dgraph-io/badger"
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

	// open DB
	db, err := badger.Open(badger.DefaultOptions(dbPath))
	if err != nil {
		log.Panic(err)
	}
	defer db.Close()

	// open read-only transaction
	err = db.View(func(txn *badger.Txn) error {
		item, err := txn.Get([]byte("key"))
		if err != nil {
			return err
		}
		val, err := item.ValueCopy(nil)
		if err != nil {
			return err
		}
		list := strings.Split(string(val), ":")
		for index, node := range list {
			log.Print(string(index) + ": " + node)
		}
		return nil
	})
}
