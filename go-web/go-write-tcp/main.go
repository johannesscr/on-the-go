package main

import (
	"fmt"
	"io"
	"log"
	"net"
)

func main() {
	// HTTP runs on top of TCP
	// if the incoming text adheres to the HTTP protocol, then the server
	// can interpret the TCP as HTTP
	li, err := net.Listen("tcp", ":8080")
	fmt.Printf("TCP server listening on: %v %v\n", li.Addr().Network(), li.Addr().String())
	if err != nil {
		log.Panic(err)
	}
	defer func(li net.Listener) {
		err := li.Close()
		if err != nil {
			log.Panic(err)
		}
	}(li)

	for {
		// waiting for a request
		// brew install telnet
		// telnet localhost 8080
		conn, err := li.Accept()
		if err != nil {
			log.Println(err)
		}

		_, err = io.WriteString(conn, "\nHello from TCP server\n")
		if err != nil {
			log.Println(err)
		}
		_, err = fmt.Fprintln(conn, "\tHow is your day?")
		if err != nil {
			log.Println(err)
		}
		_, err = fmt.Fprintf(conn, "%v", "\t\tWell I hope!\n")
		if err != nil {
			log.Println(err)
		}

		err = conn.Close()
		if err != nil {
			log.Panic(err)
		}
	}
}