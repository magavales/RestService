package main

import (
	"encoding/json"
	"errors"
	"flag"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx"
	"log"
	"net/http"
	"time"
)

var port *string

func init() {
	port = flag.String("port", "8080", "Port on which server will listen for requests")
}

func GetAllFromDB(conn *pgx.Conn) []Farmers {
	var farmersData []Farmers

	rows, _ := conn.Query("SELECT * FROM farmers")

	for rows.Next() {
		values, err := rows.Values()
		if err != nil {
			log.Fatal("error while iterating dataset")
		}
		var farmersTemp Farmers

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

	return farmersData
}

type Farmers struct {
	ID          int64
	Name        string
	Surname     string
	Country     string
	DateOfBirth time.Time
	Email       string
	Village     string
	Land        int32
	AvgIncome   string
}

func (farmers *Farmers) PrintFarmers() {
	farmersJSON, err := json.MarshalIndent(farmers, "", " ")
	if err != nil {
		log.Fatal(err.Error())
	}
	fmt.Printf("%s\n", string(farmersJSON))
}

func main() {
	flag.Parse()

	router := gin.Default()

	router.GET("/ping", func(c *gin.Context) {
		c.String(http.StatusOK, "Ping: Well done", "well done")
		log.Println("msg accepted")
	})

	router.GET("/database/select", func(c *gin.Context) {
		conf := pgx.ConnConfig{
			Host:     "localhost",
			Port:     5432,
			Database: "postgres",
			User:     "postgres",
			Password: "1703",
		}

		datapool, _ := pgx.Connect(conf)

		var farmersData []Farmers

		farmersData = GetAllFromDB(datapool)

		for _, v := range farmersData {
			v.PrintFarmers()
		}
	})

	server := &http.Server{
		Addr:    fmt.Sprintf(":%s", *port),
		Handler: router,
	}

	err := server.ListenAndServe()
	if err != nil && errors.Is(err, http.ErrServerClosed) {
		log.Printf("[server error]: %s\n", err)
	}
}