package main

import (
	"fmt"
	"net/http"
	"bytes"
)


type Node struct {
	Addr string
	Name string
}
	
// list of all connected nodes
var nodes []Node

func main() {
	local := Node{"localhost:8080", "local"}
	nodes = append(nodes, local)
}

func test(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("works")
}

func send(message string) {
	// POST a message to everyone
	for index, node := range nodes {
		fmt.Printf(node.Name)
		var jsonMessage = []byte(fmt.Sprintf(`{"message": %s}`, message))
		resp, err := http.Post("https://" + node.Addr, "application/json", bytes.NewBuffer(jsonMessage))
	}
}

func addNode(node Node) {
	fmt.Printf("add new participant")
	nodes = append(nodes, node)
}

func connect() {
	fmt.Printf("connect")

}