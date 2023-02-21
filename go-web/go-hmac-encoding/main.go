package main

import (
	"crypto/hmac"
	"crypto/sha256"
	"fmt"
	"io"
)

func main() {
	s1 := "name@example.com"
	s2 := "name@exampl.com"
	s3 := "name@example.com"

	fmt.Println("string\t\t\tHMAC SHA256")
	fmt.Printf("%s\t%v\n", s1, getCode(s1))
	fmt.Printf("%s\t\t%v\n", s2, getCode(s2))
	fmt.Printf("%s\t%v\n", s3, getCode(s3))
}

func getCode(s string) string {
	h := hmac.New(sha256.New, []byte(s))
	_, err := io.WriteString(h, s)
	if err != nil {
		fmt.Println(err)
	}
	return fmt.Sprintf("%x", h.Sum(nil))
}