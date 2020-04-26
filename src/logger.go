package main

import "io/ioutil"

// Path : to hidden directory
const Path = "/home/john/.sosimple/"

// NodesFile : file where nodes information is stored
const NodesFile = "nodes"

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func saveNodes() {
	d1 := []byte("hello\ngo\n")
	err := ioutil.WriteFile(Path+NodesFile, d1, 0644)
	check(err)
}
