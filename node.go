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

func initialize() {
	// Remove old storage directory
	if _, err := os.Stat(DBPath); os.IsExist(err) {
		os.Remove(DBPath)
	}
	// initialize storage directory
	if _, err := os.Stat(DBPath); os.IsNotExist(err) {
		os.Mkdir(DBPath, os.ModeDir)
	}

	// check connection with DB
	db, err := badger.Open(badger.DefaultOptions(DBPath))
	if err != nil {
		log.Panic(err)
	}
	defer db.Close()

	// write key to DB
	fmt.Print("Enter Group Key: ")
	pass, err := terminal.ReadPassword(int(syscall.Stdin))
	if err != nil {
		log.Panic(err)
	}
	encryptedPass := sha256.Sum256(pass)

	txn := db.NewTransaction(true)
	err = txn.SetEntry(badger.NewEntry([]byte("key"), encryptedPass[:]))
	if err != nil {
		log.Panic(err)
	}
	err = txn.Commit()
	if err != nil {
		panic(err)
	}

	//local := Node{GetPrivateIP(), "local"} // TEMP change to Public IP
	//nodes = append(nodes, local)
}

func ping(node *string) {
	resp, err := http.Get("https://" + *node + "/ping")
	if err != nil {
		log.Panic(err)
	} else if resp.StatusCode != 200 {
		log.Print("Returned status code: " + string(resp.StatusCode))
	}
}

func send(message *string) {

	// TODO iterate through all friends and send a message

	// open DB
	db, err := badger.Open(badger.DefaultOptions(DBPath))
	if err != nil {
		log.Panic(err)
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
			var jsonMessage = []byte(fmt.Sprintf(`{"message": %s}`, *message))
			resp, err := http.Post("https://"+node+"/send", "application/json", bytes.NewBuffer(jsonMessage))
			if err != nil {
				log.Print(err)
			}
			log.Print(string(resp.StatusCode))
		}
		return nil
	})
	if err != nil {
		panic(err)
	}
}

func viewLog() {

	// openDB
	db, err := badger.Open(badger.DefaultOptions(DBPath))
	if err != nil {
		log.Panic(err)
	}
	defer db.Close()

	// open read-only transaction
	err = db.View(func(txn *badger.Txn) error {
		item, err := txn.Get([]byte("messages"))
		if err != nil {
			return err
		}
		val, err := item.ValueCopy(nil)
		if err != nil {
			return err
		}
		if val == nil {
			fmt.Println("Your log history is empty")
		} else {
			list := strings.Split(string(val), ":")
			for _, message := range list {
				fmt.Println(string(message))
			}
		}
		return nil
	})
}

func join(node *string) {
	db, err := badger.Open(badger.DefaultOptions(DBPath))
	if err != nil {
		log.Panic(err)
	}
	defer db.Close()

	err = db.View(func(txn *badger.Txn) error {
		item, err := txn.Get([]byte("key"))
		if err != nil {
			return err
		}
		val, err := item.ValueCopy(nil)
		if err != nil {
			return err
		}
		key := string(val)
		url := "http://" + *node + "/join" // TODO change to HTTPS
		var jsonMessage = []byte(fmt.Sprintf(`{"key": %s}`, key))
		resp, err := http.Post(url, "application/json", bytes.NewBuffer([]byte(jsonMessage)))
		if err != nil {
			log.Panic(err)
		} else if resp.StatusCode == 403 {
			log.Panic("WRONG KEY")
		} else if resp.StatusCode == 200 {
			log.Print("CONNECTION SUCCESSFUL")
		}
		return nil
	})

	// TODO synchronize databases with a friend

}

func list() {
	// open DB
	db, err := badger.Open(badger.DefaultOptions(DBPath))
	if err != nil {
		log.Panic(err)
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
		if val == nil {
			fmt.Println("You're not connected to anyone")
		} else {
			list := strings.Split(string(val), ":")
			for _, node := range list {
				fmt.Println(string(node))
			}
		}
		return nil
	})
}
