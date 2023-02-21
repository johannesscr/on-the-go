package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
)

func main() {
	li, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Panic(err)
	}
	fmt.Printf("TCP server listening %v\n", li.Addr())

	defer func(l net.Listener){
		err := l.Close()
		if err != nil {
			log.Panic(err)
		}
	}(li)

	for {
		conn, err := li.Accept()
		if err != nil {
			log.Println(err)
		}
		go handle(conn)
	}
}

func handle(c net.Conn) {
	defer func(co net.Conn){
		err := co.Close()
		if err != nil {
			log.Panic(err)
		}
	}(c)

	// see bufio implementation
	// https://play.golang.org/p/fCpbtxH-2XK
	scanner := bufio.NewScanner(c)
	for scanner.Scan() {
		ln := scanner.Text()
		fmt.Println(ln)
	}
	// now
	// curl http://localhost:8080

	// we never get here
	// we have an open stream connection
	// how does the above reader know when it's done?
	fmt.Println("Code got here.")
}
