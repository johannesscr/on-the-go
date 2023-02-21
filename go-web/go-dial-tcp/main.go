package main

import (
	"log"
	"net"
)

func main() {
	conn, err := net.Dial("tcp", "localhost:8080")
	if err != nil {
		log.Panic(err)
	}
	defer func(c net.Conn){
		err := c.Close()
		if err != nil {
			log.Panic(err)
		}
	}(conn)

	_, err = conn.Write([]byte("I dialed you"))
	if err != nil {
		log.Println(err)
	}
}