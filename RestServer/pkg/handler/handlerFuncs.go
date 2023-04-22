package handler

import (
	"RestService"
	"RestService/pkg/auth"
	"RestService/pkg/database"
	"github.com/gin-gonic/gin"
	"log"
)

func (h *Handler) SelectAll(c *gin.Context) {
	var (
		db database.Database
	)

	db.Connect()
	db.Farmers.GetAll(db.Pool)
	err := db.Pool.Close()
	if err != nil {
		log.Printf("Connection don't close with database: %s\n", err)
	}

	for _, v := range db.Farmers.Data {
		v.PrintFarmers()
	}
}

func (h *Handler) SelectOrderByID(c *gin.Context) {
	var (
		db database.Database
	)

	db.Connect()
	db.Farmers.GetOrderByID(db.Pool)
	err := db.Pool.Close()
	if err != nil {
		log.Printf("Connection don't close with database: %s\n", err)
	}

	for _, v := range db.Farmers.Data {
		v.PrintFarmers()
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

	if db.Users.Contains(user) == 0 {
		if user.Username == "admin" {
			user.Role = 0
			db.Users.Insert(db.Pool, user)

			token, err := authen.GetToken(user.Username)

			resp.ResponseJWT(token, err)
		} else {
			user.Role = 1
			db.Users.Insert(db.Pool, user)

			token, err := authen.GetToken(user.Username)

			resp.ResponseJWT(token, err)
		}

	} else {
		resp.ResponseWrongUsername()
	}
}
