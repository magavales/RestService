package main

import (
	"RestClient"
	"RestClient/pkg/request"
	"RestClient/pkg/response"
	"encoding/json"
	"fmt"
	"log"
)

func main() {
	var (
		err    error
		menu   []string
		menu1  []string
		menu2  []string
		menu21 []string
		n      int
		auth   bool
		req    request.Request
	)
	auth = false
	user := RestClient.User{}

	client := new(RestClient.Client).InitClient()
	r := new(response.Response)

	menu = []string{
		"1. Зарегистрироваться\n",
		"2. Войти\n",
	}

	menu1 = []string{
		"1. Получить данные\n",
		"2. Получить отсортированные данные\n",
	}

	menu2 = []string{
		"1. Следующая страница\n",
		"2. Предыдущая страница\n",
		"3. Назад\n",
	}

	menu21 = []string{
		"1. Следующая страница\n",
		"2. Предыдущая страница\n",
		"3. Назад\n",
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

				req.Logon(user.ToJSON())
				r.Resp, err = client.Do(req.Request)
				if err != nil {
					log.Printf("Client hasn't been do smth!: %s\n", err)
				}
				r.GetResponse()
				auth = r.Auth
			case 2:
				fmt.Print("Введите логин: ")
				fmt.Scan(&user.Username)
				fmt.Print("Введите пароль: ")
				fmt.Scan(&user.Password)

				req.Login(user.ToJSON())
				r.Resp, err = client.Do(req.Request)
				if err != nil {
					log.Printf("Client hasn't been do smth!: %s\n", err)
				}
				r.GetResponse()
				auth = r.Auth
			}
		} else {
			for _, v := range menu1 {
				log.Printf("%s", v)
			}

			fmt.Scan(&n)

			switch n {
			case 1:
				var (
					farmers  []RestClient.Farmer
					page     int
					per_page int
					offset   int
				)
				page = 1
				per_page = 5
				offset = 0
				for true {
					req.GetData(r.Resp.Header.Get("Authorization"), page, per_page, offset)
					r.Resp, err = client.Do(req.Request)
					if err != nil {
						log.Printf("Client hasn't been do smth!: %s\n", err)
					}
					decoder := json.NewDecoder(r.Resp.Body)

					err = decoder.Decode(&farmers)
					if err != nil {
						log.Printf("Can't decode data! : %s", err)
					}
					for _, v := range farmers {
						v.Print()
					}
					if page == 1 {
						for _, v := range menu2 {
							log.Printf("%s", v)
						}
						fmt.Scan(&n)
						switch n {
						case 1:
							page++
							offset = offset + per_page
						case 2:
							break
						}
					} else {
						for _, v := range menu21 {
							log.Printf("%s", v)
						}
						fmt.Scan(&n)
						switch n {
						case 1:
							page++
							offset = offset + per_page
						case 2:
							page--
							offset = offset - per_page
						case 3:
							break
						}
					}
				}
			}
		}
	}
}
