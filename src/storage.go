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

func getValue(db *badger.DB, key []byte) error {
	return db.View(func(txn *badger.Txn) (string, error) {
		item, err := txn.Get([]byte(key))
		if err != nil {
			return "", err
		}
		val, err := item.ValueCopy(nil)
		if err != nil {
			return "", err
		}
		return string(val), nil
	})
}

func getElement(db *badger.DB, key []byte) error {

}
