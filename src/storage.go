package main

import (
	"log"
	"os"

	badger "github.com/dgraph-io/badger"
)

var dbPath = os.Getenv("HOME") + "/.sosimple/"

func openDB() *badger.DB {
	db, err := badger.Open(badger.DefaultOptions(os.Getenv("HOME") + "/.sosimple/db"))
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	return db
}
