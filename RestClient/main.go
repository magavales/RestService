package main

import (
	"log"
	"net/http"
)

func main() {
	client := http.Client{
		Transport:     nil,
		CheckRedirect: nil,
		Jar:           nil,
		Timeout:       0,
	}

	req, _ := http.NewRequest("GET", "http://localhost:8080/ping", nil)

	_, err := client.Do(req)
	if err != nil {
		log.Printf("[client error : %s\n]", err)
	}

}
