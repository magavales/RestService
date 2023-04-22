package RestService

import (
	"encoding/json"
	"io"
	"log"
)

type User struct {
	ID       int64
	Username string `json:"username"`
	Password string `json:"password"`
	Role     int16
}

func (u *User) GetData(data io.ReadCloser) {
	decoder := json.NewDecoder(data)

	err := decoder.Decode(&u)
	if err != nil {
		log.Printf("Can't decode data! : %s", err)
	}

	u.ID = u.ID
}

func (u *User) ParseData(values []interface{}) {
	u.ID = values[0].(int64)
	u.Username = values[1].(string)
	u.Password = values[2].(string)
	u.Role = values[3].(int16)
}
