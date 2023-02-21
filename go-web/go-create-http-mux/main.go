package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"strings"
	"time"
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
		time.Sleep(10 * time.Millisecond)
		err := c.Close()
		printErr(err)
	}(c)

	// write response
	//response(c)

	// read request
	request(c)
	// write response
	//response(c)
}

type Req struct {
	Headers map[string]string
	Body string
}
func (r Req) String() string {
	return fmt.Sprintf("Req{Headers:\n%v\nBody: %v\n}\n", r.Headers, r.Body)
}

// request parses the incoming request as HTTP
func request(conn net.Conn) {
	r := Req{
		Headers: make(map[string]string),
		Body: "",
	}
	section := "headers"
	j := 0
	i := 0

	scanner := bufio.NewScanner(conn)
	for scanner.Scan() {
		ln := scanner.Text()
		fmt.Println(ln)
		if section == "headers" {
			if ln == "" {
				// headers are done
				section = "body"
			} else if i == 0 {
				requestLine := strings.Fields(ln)
				r.Headers["Method"] = requestLine[0]
				r.Headers["URL"] = requestLine[1]
				r.Headers["Protocol"] = requestLine[2]
			} else {
				headerLine := strings.Split(ln, ":")
				r.Headers[headerLine[0]] = headerLine[1]
			}
		}
		if section == "body" {
			r.Body += ln
		}
		if ln == "" {
			j++
			if j == 2 {
				j = 0
				break
			}
		}
		i++
		//if ln == "" {
		//	break
		//}
	}
	mux(conn, r)
}

func mux(conn net.Conn, req Req) {
	fmt.Printf("> %v %v %v\n",
		req.Headers["Protocol"],
		req.Headers["Method"],
		req.Headers["URL"])
	fmt.Println(req)

	// write response
	response(conn)
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
	s := fmt.Sprintf("HTTP/1.1 200 OK\r\n" +
		"Content-Length: %d\r\n" +
		"Content-Type: text/html\r\n" +
		"\r\n" +
		"%v", len(body), body)
	fmt.Println(s)
	_, err := fmt.Fprint(conn, s)
	panicErr(err)

	fmt.Printf("< HTTP/1.1 200 OK\n\n")
}