package tables

import (
	"RestService"
	"github.com/jackc/pgx"
)

type UsersDB struct {
	Data []RestService.User
}

func (d *UsersDB) Insert(pool *pgx.Conn, user *RestService.User) {
	_, err := pool.Query("INSERT INTO users (login, password) VALUES ($1, $2)", user.Login, user.Password)
	if err != nil {
		return
	}
}
