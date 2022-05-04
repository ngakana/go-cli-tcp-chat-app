package main

import (
	"log"
	"net"
)

func main() {

	s := newServer()  // manage connections, clients and rooms
	go s.run()        // handle each interaction in  seperate thread

	li, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatalf("Unable to start server, %s", err.Error())
	}
	defer li.Close()
	log.Printf("Started server on port 8080")

	for {
		conn, err := li.Accept() //accept new connections
		if err != nil {
			log.Fatalf("Could not establish socket connection, %s", err.Error())
			continue
		}

		go s.newClient(conn)
	}
}