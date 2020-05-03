package main

import (
	//"flag"

	"log"
	"net/http"

	badger "github.com/dgraph-io/badger"
	"github.com/gorilla/mux"
)

func launchServer() { // TODO expose to the world
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/ping", pong).Methods("GET")
	router.HandleFunc("/join", authorize).Methods("POST")
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
		key, err := item.ValueCopy(nil)
		if err != nil {
			return err
		}
		// retrieve key of requester
		foreignKeys, ok := r.URL.Query()["key"]
		foreignKey := foreignKeys[0]

		if !ok || len(foreignKey) < 1 {
			log.Print("WARNING!")
			log.Print("SOMEBODY TRIED TO CONNECT WITH YOU WITHOUT KEY")
			w.Write([]byte("KEY IS MISSING"))
		}

		// compare keys
		if string(key) != foreignKey {
			log.Print("WARNING!")
			log.Print("SOMEBODY TRIED TO CONNECT WITH YOU WITH WRONG KEY")
			w.Write([]byte("WRONG KEY"))
		}
		return nil
	})
}
