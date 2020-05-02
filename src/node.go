package main

import (
	"bytes"
	"crypto/sha256"
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"
	"syscall"

	badger "github.com/dgraph-io/badger"
	"golang.org/x/crypto/ssh/terminal"
	//auth "github.com/abbot/go-http-auth"
)

// TODO maybe in future
//type node struct {
//	Addr string
//	Name string
//}

func initialize() {
	log.Print("Initializing node")

	// Remove old storage directory
	if _, err := os.Stat(dbPath); os.IsExist(err) {
		os.Remove(dbPath)
	}
	// initialize storage directory
	if _, err := os.Stat(dbPath); os.IsNotExist(err) {
		os.Mkdir(dbPath, os.ModeDir)
	}

	// check connection with DB
	db, err := badger.Open(badger.DefaultOptions(dbPath))
	if err != nil {
		log.Panic(err)
	}
	defer db.Close()

	// write key to DB
	log.Print("Enter Group Key: ")
	pass, err := terminal.ReadPassword(int(syscall.Stdin))
	if err != nil {
		log.Panic(err)
	}

	txn = db.NewTransaction(true)
	err = txn.SetEntry(badger.NewEntry([]byte("key"), []byte(sha256.Sum256(pass))))
	if err != nil {
		log.Panic(err)
	}
	err txn.Commit()
	if err != nul {
		panic(err)
	}

	//local := Node{GetPrivateIP(), "local"} // TEMP change to Public IP
	//nodes = append(nodes, local)
}

func send(message string) {

	// TODO iterate through all connected nodes and send a message

	// open DB
	db, err := badger.Open(badger.DefaultOptions(dbPath))
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// open read-only transaction
	err = db.View(func(txn *badger.Txn) error {
		item, err := txn.Get([]byte("nodes"))
		if err != nil {
			return err
		}
		val, err := item.ValueCopy(nil)
		if err != nil {
			return err
		}
		list := strings.Split(string(val), ":")
		for index, node := range list {
			log.Print(string(index) + " " + node)
			var jsonMessage = []byte(fmt.Sprintf(`{"message": %s}`, message))
			resp, err := http.Post("https://"+node+"/", "application/json", bytes.NewBuffer(jsonMessage))
			if err != nil {
				log.Print(err)
			} else {
				log.Print(string(resp.StatusCode))
			}
		}
		return nil
	})
	if err != nil {
		panic(err)
	}
}

func join(node string, key string) {
	log.Print("join")

	//resp, err := http.NewRequest("PUT", url)
	url := "https://" + node + "/join?key=" + key
	resp, err := http.Post(url, "", bytes.NewBuffer([]byte("")))
	if err != nil {
		log.Panic(err)
	} else if resp.StatusCode == 403 {
		log.Panic("WRONG KEY")
	} else if resp.StatusCode == 200 {
		log.Print("CONNECTION SUCCESSFUL")

		// TODO synchronize databases with all new friends
	}
}

func list() {
	// open DB
	db, err := badger.Open(badger.DefaultOptions(dbPath))
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// open read-only transaction
	err = db.View(func(txn *badger.Txn) error {
		item, err := txn.Get([]byte("nodes"))
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
