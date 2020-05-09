package main

import (
	"fmt"
	"github.com/dgraph-io/badger"
	"log"
	"os"
	"strconv"
)

func main() {
	// Open the Badger database located in the /tmp/badger directory.
	// It will be created if it doesn't exist.
	db, err := badger.Open(badger.DefaultOptions(os.Getenv("HOME") + "/sosimple-P2P/research/.badg"))
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	err = db.View(func(txn *badger.Txn) error {
		_, err := txn.Get([]byte("key"))
		// We expect ErrKeyNotFound
		fmt.Println(err)
		return nil
	})

	if err != nil {
		panic(err)
	}

	txn := db.NewTransaction(true) // Read-write txn
	err = txn.SetEntry(badger.NewEntry([]byte("key"), []byte(stringGenerator(100000))))
	if err != nil {
		panic(err)
	}
	err = txn.Commit()
	if err != nil {
		panic(err)
	}

	err = db.View(func(txn *badger.Txn) error {
		item, err := txn.Get([]byte("key"))
		if err != nil {
			return err
		}
		val, err := item.ValueCopy(nil)
		if err != nil {
			return err
		}
		fmt.Printf("VALUE: ")
		fmt.Printf("%s\n", string(val))
		return nil
	})

	if err != nil {
		panic(err)
	}

	//

	err = db.View(func(txn *badger.Txn) error {
		item, err := txn.Get([]byte("key"))
		if err != nil {
			return err
		}
		val, err := item.ValueCopy(nil)
		if err != nil {
			return err
		}
		fmt.Printf("%s\n", string(val))
		return nil
	})

	if err != nil {
		panic(err)
	}
}

func stringGenerator(x int) string {
	var str = ""
	for i := 0; i < x; i++ {
		str = str + strconv.Itoa(i)
	}
	return str
}