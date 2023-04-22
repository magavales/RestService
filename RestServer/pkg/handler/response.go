package handler

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

type Response struct {
	rw gin.ResponseWriter
}

func (resp *Response) ResponseJWT(token string, mis error) {
	if mis == nil {
		resp.rw.Header().Set("Content-Type", "application/json")
		resp.rw.Header().Set("Authorization", "Bearer "+token)
		resp.rw.WriteHeader(http.StatusOK)
		_, err := resp.rw.Write([]byte(`{"message": "Registration has been made!"}`))
		if err != nil {
			log.Printf("Couldn't write message! : %s\n", err)
		}
	} else {
		resp.rw.Header().Set("Content-Type", "application/json")
		resp.rw.WriteHeader(http.StatusInternalServerError)
		_, err := resp.rw.Write([]byte(`{"message": "Something wrong!"}`))
		if err != nil {
			log.Printf("Couldn't write message! : %s\n", err)
		}
	}
}

func (resp *Response) ResponseWrongUsername() {
	resp.rw.Header().Set("Content-Type", "application/json")
	resp.rw.WriteHeader(http.StatusBadRequest)
	_, err := resp.rw.Write([]byte(`{"message": "This username is already in use!"}`))
	if err != nil {
		log.Printf("Couldn't write message! : %s\n", err)
	}
}
