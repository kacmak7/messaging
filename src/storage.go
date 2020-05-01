package main

import (
	"log"
	"os"

	badger "github.com/dgraph-io/badger"
)

func openDB() {
	db, err := badger.Open(badger.DefaultOptions(os.Getenv("HOME") + "/.sosimple/db"))
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
}
