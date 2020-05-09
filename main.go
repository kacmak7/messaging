package main

import (
	"fmt"
	"log"
	"os"

	"github.com/akamensky/argparse"
	"github.com/sevlyar/go-daemon"
)

func main() {
	// Main arguments parser
	parser := argparse.NewParser("commands", "Available sosimple commands")

	// Init command
	initCmd := parser.NewCommand("init", "Initialize sosimple node listener")
	//initCmdPort := initCmd.String("p", "port", &argparse.Options{Required: false, Help: "Port to allocate"})

	// Daemon command
	daemonCmd := parser.NewCommand("daemon", "Start daemon process")

	// Shutdown command
	shutdownCmd := parser.NewCommand("shutdown", "Shutdown sosimple node listener")

	// Ping command
	pingCmd := parser.NewCommand("ping", "Ping connected Node")
	pingCmdNode := pingCmd.String("n", "node", &argparse.Options{Required: true, Help: "Node to ping"})

	// Join command
	joinCmd := parser.NewCommand("join", "Join to node and his friends")
	joinCmdNode := joinCmd.String("n", "node", &argparse.Options{Required: true, Help: "Node to ping"})

	// Send command
	sendCmd := parser.NewCommand("send", "Send a message")
	sendCmdMessage := sendCmd.String("m", "message", &argparse.Options{Required: true, Help: "Message to send"})

	// Log command
	logCmd := parser.NewCommand("log", "View messages")

	// List command
	listCmd := parser.NewCommand("list", "List all your friends")

	err := parser.Parse(os.Args)
	if err != nil {
		fmt.Print(parser.Usage(err))
		return
	}

	if initCmd.Happened() {
		initialize()
		log.Print("Node successfully initialized")
	} else if daemonCmd.Happened() {
		cntxt := &daemon.Context{
			PidFileName: "daemon.pid",
			PidFilePerm: 0644,
			LogFileName: "daemon.log",
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

		launchServer()
	} else if shutdownCmd.Happened() {
		log.Print("Shutting down")
		// TODO
		log.Print("not yet implemented")
	} else if pingCmd.Happened() {
		for i := 0; i < 6; i++ {
			ping(pingCmdNode)
		}
	} else if joinCmd.Happened() {
		join(joinCmdNode)
	} else if sendCmd.Happened() {
		send(sendCmdMessage)
	} else if logCmd.Happened() {
		viewLog()
	} else if listCmd.Happened() {
		list()
	}
}
