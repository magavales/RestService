package handler

import (
	utilites "RestService"
	"RestService/pkg/database"
	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx"
	"log"
)

func (h *Handler) SelectAll(c *gin.Context) {
	var db database.Database
	var farmersData []utilites.Farmer

	conf := pgx.ConnConfig{
		Host:     "localhost",
		Port:     5432,
		Database: "postgres",
		User:     "postgres",
		Password: "1703",
	}

	db.Connect(conf)
	db.GetAll(&farmersData)
	err := db.Pool.Close()
	if err != nil {
		log.Printf("Don't close connection with database: %s\n", err)
	}

	for _, v := range farmersData {
		v.PrintFarmers()
	}
}

func (h *Handler) SelectOrderByID(c *gin.Context) {
	var db database.Database
	var farmersData []utilites.Farmer

	conf := pgx.ConnConfig{
		Host:     "localhost",
		Port:     5432,
		Database: "postgres",
		User:     "postgres",
		Password: "1703",
	}

	db.Connect(conf)
	db.GetOrderByID(&farmersData)
	err := db.Pool.Close()
	if err != nil {
		log.Printf("Don't close connection with database: %s\n", err)
	}

	for _, v := range farmersData {
		v.PrintFarmers()
	}
}
