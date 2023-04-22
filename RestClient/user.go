package RestClient

import (
	"encoding/json"
	"log"
)

type User struct {
	ID       int64  `json:"ID"`
	Username string `json:"username"`
	Password string `json:"password"`
	Role     int16  `json:"role"`
}

func (u *User) CreateNewUser(Username string, Password string, Role int16) {
	u.Username = Username
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
