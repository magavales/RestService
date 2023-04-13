package RestService

import (
	"encoding/json"
	"io"
	"log"
)

type User struct {
	ID       int64
	Login    string
	Password string
	Role     int16
}

func (u *User) FromJSON(data io.ReadCloser) {
	decoder := json.NewDecoder(data)

	err := decoder.Decode(&u)
	if err != nil {
		log.Printf("Can't decode data! : %s", err)
	}
}
