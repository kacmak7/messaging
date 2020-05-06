package main

import (
	"fmt"
	"html"
	"log"
	"net/http"
	"os"

	"github.com/akamensky/argparse"
	"github.com/sevlyar/go-daemon"
)

// To terminate the daemon use:
//  kill `cat sample.pid`
func main() {
	parser := argparse.NewParser("commands", "Available sosimple commands")
	daemonCmd := parser.NewCommand("daemon", "Start daemon process")

	err := parser.Parse(os.Args)

	if err != nil {
		fmt.Print(parser.Usage(err))
		return
	}

	if daemonCmd.Happened() {
		cntxt := &daemon.Context{
			PidFileName: "sample.pid",
			PidFilePerm: 0644,
			LogFileName: "sample.log",
			LogFilePerm: 0640,
			WorkDir:     "./",
			Umask:       027,
			Args:        nil,
		}

		d, err := cntxt.Reborn()
		if err != nil {
			log.Fatal("Unable to run: ", err)
		}
		if d != nil {
			return
		}
		defer cntxt.Release()

		log.Print("- - - - - - - - - - - - - - -")
		log.Print("daemon started")

		serveHTTP()
	}
}

func serveHTTP() {
	http.HandleFunc("/", httpHandler)
	http.ListenAndServe("127.0.0.1:8080", nil)
}

func httpHandler(w http.ResponseWriter, r *http.Request) {
	log.Printf("request from %s: %s %q", r.RemoteAddr, r.Method, r.URL)
	fmt.Fprintf(w, "go-daemon: %q", html.EscapeString(r.URL.Path))
}
