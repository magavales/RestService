package request

import (
	"bytes"
	"fmt"
	"net/http"
)

type Request struct {
	Request *http.Request
}

func (r *Request) Logon(data []byte) {
	r.Request, _ = http.NewRequest("POST", "http://localhost:8080/api/user/logon", bytes.NewReader(data))
	r.Request.Header.Set("Content-Type", "application/json")
}

func (r *Request) Login(data []byte) {
	r.Request, _ = http.NewRequest("POST", "http://localhost:8080/api/user/login", bytes.NewReader(data))
	r.Request.Header.Set("Content-Type", "application/json")
}

func (r *Request) GetData(token string, page, per_page, offset int) {
	url := fmt.Sprintf("http://localhost:8080/api/database/all?page=%d&per_page=%d&offset=%d", page, per_page, offset)
	r.Request, _ = http.NewRequest("GET", url, bytes.NewReader([]byte("")))
	r.Request.Header.Set("Content-Type", "application/json")
	r.Request.Header.Set("Authorization", token)
}
