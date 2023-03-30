package handler

import (
	utilites "RestService"
	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx"
	"log"
	"time"
)

func GetAllFromDB(conn *pgx.Conn) *[]utilites.Farmer {
	var farmersData []utilites.Farmer

	rows, _ := conn.Query("SELECT * FROM farmers")

	for rows.Next() {
		values, err := rows.Values()
		if err != nil {
			log.Fatal("error while iterating dataset")
		}
		var farmersTemp utilites.Farmer

		farmersTemp.ID = values[0].(int64)
		farmersTemp.Name = values[1].(string)
		farmersTemp.Surname = values[2].(string)
		farmersTemp.Country = values[3].(string)
		farmersTemp.DateOfBirth, _ = time.Parse("2006-01-02", values[4].(time.Time).Format("2006-01-02"))
		farmersTemp.Email = values[5].(string)
		farmersTemp.Village = values[6].(string)
		farmersTemp.Land = values[7].(int32)
		farmersTemp.AvgIncome = values[8].(string)

		farmersData = append(farmersData, farmersTemp)
	}

	return &farmersData
}

func (h *Handler) SelectAll(c *gin.Context) {
	conf := pgx.ConnConfig{
		Host:     "localhost",
		Port:     5432,
		Database: "postgres",
		User:     "postgres",
		Password: "1703",
	}

	datapool, _ := pgx.Connect(conf)

	farmersData := new([]utilites.Farmer)

	farmersData = GetAllFromDB(datapool)

	for _, v := range farmersData {
		v.PrintFarmers()
	}
}
