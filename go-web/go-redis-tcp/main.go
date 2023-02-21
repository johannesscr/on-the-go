package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"strings"
)

func main() {
	li, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Panic(err)
	}
	fmt.Printf("TCP listening on %v\n", li.Addr())
	defer func(l net.Listener) {
		err := l.Close()
		if err != nil {
			log.Panic(err)
		}
	}(li)

	for {
		conn, err := li.Accept()
		if err != nil {
			log.Panic(err)
		}
		go handle(conn)
	}

}

func connWrite(conn net.Conn, text string) {
	_, err := conn.Write([]byte("invalid command\n"))
	if err != nil {
		log.Println(err)
	}
}

func handle(conn net.Conn) {
	defer func(net.Conn) {
		err := conn.Close()
		if err != nil {
			log.Panic(err)
		}
	}(conn)

	startMessage := "\nIN-MEMORY DATABASE\n\n" +
		"USE:\n" +
		"\tSET key value\n" +
		"\tGET key\n" +
		"\tDEL key\n\n" +
		"EXAMPLE:\n" +
		"SET fav chocolate\n" +
		"GET fav\n" +
		"> chocolate\n"

	_, err := conn.Write([]byte(startMessage))
	if err != nil {
		log.Println(err)
	}

	// read and write
	data := make(map[string]string)

	scanner := bufio.NewScanner(conn)
	for scanner.Scan() {
		ln := scanner.Text()
		fs := strings.Fields(ln) // splits on spaces
		cmd := fs[0]
		key := fs[1]
		fmt.Println(fs, cmd, key)
		switch cmd {
		case "GET":
			connWrite(conn, fmt.Sprintf("> %v\n", data[key]))
			break
		case "SET":
			data[key] = fs[2]
			connWrite(conn, "\n")
			break
		case "DEL":
			delete(data, key)
			connWrite(conn, fmt.Sprintf("> %v removed\n", key))
			break
		default:
			connWrite(conn, "Invalid command\n")
		}
	}
}
