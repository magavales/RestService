package response

import (
	"encoding/json"
	"log"
	"net/http"
)

type Message struct {
	Msg    string `json:"message"`
	Status int
}

func (m *Message) Print() {
	log.Printf("Response from server: %s\n"+
		"StatusCode: %d\n", m.Msg, m.Status)
}

type Response struct {
	Resp *http.Response
	Msg  Message
}

func (r *Response) GetResponse() {
	decoder := json.NewDecoder(r.Resp.Body)
	err := decoder.Decode(&r.Msg)
	if err != nil {
		log.Printf("Decoder failed to decode! : %s\n", err)
	}
	r.Msg.Status = r.Resp.StatusCode
}
