package main

import (
	"net/http"
	"bytes"
	"log"
)


type Node struct {
	Addr string
	Name string
}
	
// list of all connected nodes
var nodes []Node

func initializeNode() {
	local := Node{GetPrivateIP(), "local"} // TEMP change to Public IP
	nodes = append(nodes, local)
}

func test(w http.ResponseWriter, r *http.Request) {
	log.Print("works")
}

func send(message string) {
	// POST a message to everyone
	log.Print("send")
	for index, node := range nodes {
		log.Print(string(index))
		log.Print(node.Name)
		var jsonMessage = []byte(fmt.Sprintf(`{"message": %s}`, message))
		resp, err := http.Post("https://" + node.Addr, "application/json", bytes.NewBuffer(jsonMessage))
		if err != nil {
			log.Fatal(err)
		}
		log.Print(string(resp.StatusCode))
	}
}

func addNode(node Node) {
	log.Print("add new participant")
	nodes = append(nodes, node)
}

func connect() {
	log.Print("connect")

}