package main

import (
	"bytes"
	"fmt"
	"log"
	"net/http"
	"os"
	//auth "github.com/abbot/go-http-auth"
)

type node struct {
	Addr string
	Name string
}

func initializeNode() {
	log.Print("Initializing node")

	// TODO 2 DATABASES BADGER

	// Remove old storage directory
	if _, err := os.Stat(dbPath); os.IsExist(err) {
		os.Remove(dbPath)
	}
	// initialize storage directory
	if _, err := os.Stat(dbPath); os.IsNotExist(err) {
		os.Mkdir(dbPath, os.ModeDir)
	}
	// check connection
	openDB()

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

	// TODO fetch nodes from badger DB

	for index, node := range nodes {
		log.Print(string(index))
		log.Print(node.Name)
		var jsonMessage = []byte(fmt.Sprintf(`{"message": %s}`, message))
		resp, err := http.Post("https://"+node.Addr, "application/json", bytes.NewBuffer(jsonMessage))
		if err != nil {
			log.Print(err)
		} else {
			log.Print(string(resp.StatusCode))
		}
	}
}

func addNode(node node) {
	log.Print("add new participant")

	// add to DB
}

func listNodes() {

}

func connect() {
	log.Print("connect")
}
