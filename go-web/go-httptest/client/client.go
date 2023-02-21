package main

import (
	"io/ioutil"
	"log"
	"net/http"
)

type Client struct {
	url string
}

func NewClient(url string) Client {
	return Client{url: url}
}

func (c Client) UpperCase(word string) (string, error) {
	res, err := http.Get(c.url + "/upper?word=" + word)
	if err != nil {
		return "", err
	}
	defer func(){
		err := res.Body.Close()
		if err != nil {
			log.Fatal(err)
		}
	}()

	out, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return "", err
	}
	return string(out), nil
}
