package main

import (
	"encoding/base64"
	"fmt"
	"log"
)

func main() {
	fmt.Println("should we want to send data somewhere, but our data\n" +
		"includes data with an invalid format. Such as sending data json\n" +
		"data that contains \" but we cannot normally send \", then\n" +
		"we can encode it to base64 and send it easily")
	jsonString := "{\"name\":\"ben\"}"
	cb64 := customEncodingStandard([]byte(jsonString))
	b64 := standardEncoding([]byte(jsonString))

	fmt.Println("string\t\tcustom encoding\t\tstandard encoding")
	fmt.Printf("%s\t%s\t%s\n\n", jsonString, cb64, b64)

	fmt.Println("custom encoding\t\tstring")
	fmt.Printf("%s\t%s\n\n", cb64, customDecodingStandard(cb64))

	fmt.Println("standard encoding\tstring")
	fmt.Printf("%s\t%s\n", b64, standardDecoding(b64))
}

func customEncodingStandard(bs []byte) string {
	encodeStd := "ABCDEFGHIJKLMNOPQRSTUVWXYZabvdefghijklmnopqrstuvwxyz0123456789+/"
	return base64.NewEncoding(encodeStd).EncodeToString(bs)
}

func customDecodingStandard(s string) string {
	encodeStd := "ABCDEFGHIJKLMNOPQRSTUVWXYZabvdefghijklmnopqrstuvwxyz0123456789+/"
	bs, err := base64.NewEncoding(encodeStd).DecodeString(s)
	if err != nil {
		log.Println(err)
	}
	return fmt.Sprintf("%s", bs)
}

func standardEncoding(bs []byte) string {
	return base64.StdEncoding.EncodeToString(bs)
}
func standardDecoding(s string) string {
	bs, err := base64.StdEncoding.DecodeString(s)
	if err != nil {
		log.Println(err)
	}
	return fmt.Sprintf("%s", bs)
}