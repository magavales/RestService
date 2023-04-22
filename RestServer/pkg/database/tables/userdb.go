package tables

import (
	"RestService"
	"github.com/jackc/pgx"
	"log"
)

type UsersDB struct {
	Data []RestService.User
}

func (udb *UsersDB) Insert(pool *pgx.Conn, user RestService.User) {
	_, err := pool.Query("INSERT INTO users (login, password) VALUES ($1, $2)", user.Username, user.Password)
	if err != nil {
		return
	}
}

func (udb *UsersDB) GetAll(pool *pgx.Conn) {
	rows, err := pool.Query("SELECT * FROM users")
	if err != nil {
		log.Printf("The request was made incorrectly: %s\n", err)
	}

	for rows.Next() {
		value, err := rows.Values()
		if err != nil {
			log.Fatal("error while iterating dataset")
		}
		var userTemp RestService.User

		userTemp.ParseData(value)

		udb.Data = append(udb.Data, userTemp)
	}
}

func (udb *UsersDB) Contains(user RestService.User) int {
	for _, v := range udb.Data {
		if v.Username == user.Username {
			return 1
		}
	}
	return 0
}
