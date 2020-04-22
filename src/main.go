package main 

import (
	//"flag"
	//"log"
	"time"
)

func main() {
	//port := flag.String("port", "")

	go launchServer()

	initializeNode()

	send("HI HELLOooo")
	time.Sleep(100 * time.Second)
}
