package main

import (
	"bytes"
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"

	badger "github.com/dgraph-io/badger"
	//auth "github.com/abbot/go-http-auth"
)

// TODO maybe in future
//type node struct {
//	Addr string
//	Name string
//}

func initializeNode() {
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
		log.Fatal(err)
	}
	defer db.Close()

	//local := Node{GetPrivateIP(), "local"} // TEMP change to Public IP
	//nodes = append(nodes, local)
}

func ping(w http.ResponseWriter, r *http.Request) {
	log.Print("TEST CALL") // TODO attach IP of requester
	w.Write([]byte("pong\n"))
}

func authorize(w http.ResponseWriter, r *http.Request) {
	log.Print("authorizing new Node")
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
		log.Print(string(val))
		list := strings.Split(string(val), ":")
		for index, node := range list {
			log.Print(string(index) + " " + node)
			var jsonMessage = []byte(fmt.Sprintf(`{"message": %s}`, message))
			resp, err := http.Post("https://"+node, "application/json", bytes.NewBuffer(jsonMessage))
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

func addNode(node string) {
	log.Print("add new participant")

	// add to DB
}

func listNodes() {

}

func connect() {
	log.Print("connect")
}
