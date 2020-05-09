package main

import (
	"log"
	"os"

	badger "github.com/dgraph-io/badger"
)

var DBPath = os.Getenv("HOME") + "/.sosimple/"

func updateNodes(db *badger.DB, node string) error {
	// open read-write transaction
	err := db.Update(func(txn *badger.Txn) error {
		item, err := txn.Get([]byte("nodes"))
		if err != nil {
			return err
		}
		val, err := item.ValueCopy(nil)
		if err != nil {
			return err
		}
		newList := string(val) + ":" + node

		txn = db.NewTransaction(true)
		err = txn.SetEntry(badger.NewEntry([]byte("nodes"), []byte(newList)))
		if err != nil {
			log.Panic(err)
		}
		err = txn.Commit()
		if err != nil {
			log.Panic(err)
		}

		return nil
	})

	return err
}
