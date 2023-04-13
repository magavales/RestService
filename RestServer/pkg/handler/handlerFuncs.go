package handler

import (
	"RestService"
	"RestService/pkg/database"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

func (h *Handler) SelectAll(c *gin.Context) {
	var db database.Database
	var farmersData []RestService.Farmer

	db.Connect()
	db.Farmers.GetAll(db.Pool)
	err := db.Pool.Close()
	if err != nil {
		log.Printf("Connection don't close with database: %s\n", err)
	}

	for _, v := range farmersData {
		v.PrintFarmers()
	}
}

func (h *Handler) SelectOrderByID(c *gin.Context) {
	var db database.Database
	var farmersData []RestService.Farmer

	db.Connect()
	db.Farmers.GetOrderByID(db.Pool)
	err := db.Pool.Close()
	if err != nil {
		log.Printf("Connection don't close with database: %s\n", err)
	}

	for _, v := range farmersData {
		v.PrintFarmers()
	}
}

func (h *Handler) SignOn(c *gin.Context) {
	user := new(RestService.User)
	user.FromJSON(c.Request.Body)

	var db database.Database

	db.Connect()
	db.Users.Insert(db.Pool, user)

	c.Writer.Header().Set("Content-Type", "application/json")
	c.Writer.WriteHeader(http.StatusOK)
	_, err := c.Writer.Write([]byte(`{"message": "Well done!"}`))
	if err != nil {
		return
	}
}
