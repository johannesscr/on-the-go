package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"strings"
)

func panicErr(e error) {
	if e != nil {
		log.Panicln(e)
	}
}

func printErr(e error) {
	if e != nil {
		log.Println(e)
	}
}

func main() {
	li, err := net.Listen("tcp", ":8080")
	panicErr(err)
	fmt.Printf("TCP: HTTP Go Server Listening on %v\n", li.Addr().String())
	defer func(l net.Listener){
		err := l.Close()
		panicErr(err)
	}(li)

	for {
		conn, err := li.Accept()
		printErr(err)

		go handle(conn)
	}

}

// handle the connection with incoming requests and sending responses
func handle(c net.Conn) {
	defer func(c net.Conn){
		err := c.Close()
		printErr(err)
	}(c)

	// read request
	request(c)
	// write response
	response(c)
}

// request parses the incoming request as HTTP
func request(conn net.Conn) {
	i := 0
	scanner := bufio.NewScanner(conn)
	for scanner.Scan() {
		ln := scanner.Text()
		if i == 0 {
			// request line
			requestLine := strings.Fields(ln)
			method := requestLine[0]
			url := requestLine[1]
			protocol := requestLine[2]
			fmt.Printf("> %v %v %v\n", protocol, method, url)
		}
		fmt.Println(ln)
		if ln == "" {
			// headers are done
			break
		}
		i++
	}
}

// response writes the response to the client following the HTTP standards
func response(conn net.Conn) {
	body := `
<!DOCTYPE html>
<html lang="en">
<head>
	<meta charset="UTF-8">
	<title>res title</title>
</head>
<body>
	<div>
		<strong>My Response from my HTTP Server</strong>
	</div>
</body>
</html>
`
	// can span it over multiple lines or a single line of text
	//_, err := fmt.Fprint(conn, "HTTP/1.1 200 OK\r\n")
	//printErr(err)
	//_, err = fmt.Fprintf(conn, "Content-Length: %d\r\n", len(body))
	//printErr(err)
	//_, err = fmt.Fprint(conn, "Content-Type: text/html\r\n")
	//printErr(err)
	//_, err = fmt.Fprint(conn, "\r\n")
	//printErr(err)
	//_, err = fmt.Fprint(conn, body)
	//printErr(err)
	_, err := fmt.Fprintf(conn, "HTTP/1.1 200 OK\r\n" +
		"Content-Length: %d\r\n" +
		"Content-Type: text/html\r\n" +
		"\r\n" +
		"%v", len(body), body)
	printErr(err)

	fmt.Printf("< HTTP/1.1 200 OK\n\n")
}