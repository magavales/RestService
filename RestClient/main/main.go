package main

import (
	"RestClient"
	"RestClient/pkg/response"
	"bytes"
	"log"
	"net/http"
)

func main() {
	var err error
	client := new(RestClient.Client).InitClient()

	user := RestClient.User{
		Login:    "Car",
		Password: "!2saf",
	}

	req, _ := http.NewRequest("POST", "http://localhost:8080/user/signOn", bytes.NewReader(user.ToJSON()))
	req.Header.Set("Content-Type", "application/json")
	r := new(response.Response)

	r.Resp, err = client.Do(req)
	if err != nil {
		log.Printf("Client hasn't been do smth!: %s\n", err)
	}

	r.GetResponse()
	r.Msg.Print()
}
