package RestClient

import (
	"encoding/json"
	"log"
)

type User struct {
	ID       int64
	Login    string
	Password string
	Role     int16
}

func (u *User) CreateNewUser(Login string, Password string, Role int16) {
	u.Login = Login
	u.Password = Password
	u.Role = Role
}

func (u *User) ToJSON() []byte {
	body, err := json.Marshal(u)
	if err != nil {
		log.Printf("Convertation has been stopped!: %s", err)
	}

	return body
}
