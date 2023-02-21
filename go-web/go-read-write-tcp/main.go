package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"time"
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

	err := c.SetDeadline(time.Now().Add(10 * time.Second))
	if err != nil {
		log.Println(err)
	}

	// https://play.golang.org/p/fCpbtxH-2XK
	scanner := bufio.NewScanner(c)
	for scanner.Scan() {
		ln := scanner.Text()
		fmt.Println(ln)
		_, err := fmt.Fprintf(c, "I heard you say: %v\n", ln)
		if err != nil {
			log.Println(err)
		}
	}
	// curl http://localhost:8080
	// telnet localhost 8080

	fmt.Println("*** Code got here ***")
}
