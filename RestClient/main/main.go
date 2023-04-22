package main

import (
	"RestClient"
	"RestClient/pkg/response"
	"bytes"
	"fmt"
	"log"
	"net/http"
)

func main() {
	var (
		err  error
		menu []string
		n    int
		auth bool
	)
	auth = false
	user := RestClient.User{}

	client := new(RestClient.Client).InitClient()
	r := new(response.Response)

	menu = []string{
		"1. Зарегистрироваться\n",
		"2. Войти\n",
	}
	for true {
		if auth == false {
			for _, v := range menu {
				log.Printf("%s", v)
			}

			fmt.Scan(&n)

			switch n {
			case 1:
				fmt.Print("Введите логин: ")
				fmt.Scan(&user.Username)
				fmt.Print("Введите пароль: ")
				fmt.Scan(&user.Password)

				req, _ := http.NewRequest("POST", "http://localhost:8080/api/user/login", bytes.NewReader(user.ToJSON()))
				req.Header.Set("Content-Type", "application/json")
				r.Resp, err = client.Do(req)
				if err != nil {
					log.Printf("Client hasn't been do smth!: %s\n", err)
				}
				r.GetResponse()
				auth = r.Auth
			case 2:

			}
		}
	}

	r.GetResponse()
	r.Msg.Print()
}
