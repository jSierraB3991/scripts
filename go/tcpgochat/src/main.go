package main

import (
	"log"
	"net"
)

func main() {
	s := NewServer()
	go s.Run()
	listener, err := net.Listen("tcp", ":8888")
	if err != nil {
		log.Fatalf("Unable to start server: %s", err.Error())
	}

	defer listener.Close()
	log.Printf("Start server on :8888")
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Printf("Unable to accept connection: %s", err.Error())
			continue
		}

		go s.newClient(conn)
	}
}
