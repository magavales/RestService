package handler

import (
	"RestService"
	"RestService/pkg/auth"
	"RestService/pkg/database"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"log"
	"strconv"
)

func (h *Handler) SelectAll(c *gin.Context) {
	var (
		db     database.Database
		authen auth.Authentication
		resp   Response
	)
	if authen.CheckToken(c.Request.Header.Get("Authorization")) == true {
		db.Connect()
		db.Farmers.GetAll(db.Pool)
		resp.rw = c.Writer
		i, _ := strconv.Atoi(c.Request.URL.Query().Get("offset"))
		j, _ := strconv.Atoi(c.Request.URL.Query().Get("per_page"))
		j = j + i
		body, err := json.Marshal(db.Farmers.Data[i:j])
		if err != nil {
			log.Printf("Convertation has been stopped!: %s", err)
		}
		resp.ResponseGetData(c.Request.Header.Get("Authorization"), body)
	}
}

func (h *Handler) SelectOrderByID(c *gin.Context) {
	var (
		db     database.Database
		authen auth.Authentication
	)

	if authen.CheckToken(c.Request.Header.Get("Authorization")) == true {
		db.Connect()
		db.Farmers.GetAll(db.Pool)
	}

}

func (h *Handler) Logon(c *gin.Context) {
	var (
		db     database.Database
		resp   Response
		authen auth.Authentication
		user   RestService.User
	)
	db.Connect()
	db.Users.GetAll(db.Pool)

	resp.rw = c.Writer
	user.GetData(c.Request.Body)

	if db.Users.Contains(user) == 1 {
		if user.Username == "admin" {
			user.Role = 0
			db.Users.Insert(db.Pool, user)

			token, err := authen.GetToken(user.Username)
			if err != nil {
				resp.ResponseWrongJWT()
			} else {
				resp.ResponseJWT(token)
			}

			resp.ResponseJWT(token)
		} else {
			user.Role = 1
			db.Users.Insert(db.Pool, user)

			token, err := authen.GetToken(user.Username)
			if err != nil {
				resp.ResponseWrongJWT()
			} else {
				resp.ResponseJWT(token)
			}
		}

	} else {
		resp.ResponseWrongUsername()
	}
}

func (h *Handler) Login(c *gin.Context) {
	var (
		db     database.Database
		resp   Response
		authen auth.Authentication
		user   RestService.User
	)
	db.Connect()
	db.Users.GetAll(db.Pool)

	resp.rw = c.Writer
	user.GetData(c.Request.Body)

	if db.Users.PasswordVerification(user) == 0 {
		token, err := authen.GetToken(user.Username)
		if err != nil {
			resp.ResponseWrongJWT()
		} else {
			resp.ResponseJWT(token)
		}
	} else {
		resp.ResponseWrongLogin()
	}
}
