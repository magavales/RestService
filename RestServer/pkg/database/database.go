package database

import (
	utilites "RestService"
	"github.com/jackc/pgx"
	"log"
	"time"
)

type Database struct {
	Pool *pgx.Conn
}

func (db *Database) Connect(config pgx.ConnConfig) {
	var err error
	db.Pool, err = pgx.Connect(config)
	if err != nil {
		log.Printf("I can't connect to database: %s\n", err)
	}
}

func (db *Database) GetAll(farmersData *[]utilites.Farmer) {
	rows, err := db.Pool.Query("SELECT * FROM farmers")
	if err != nil {
		log.Printf("The request was made incorrectly: %s\n", err)
	}

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

		*farmersData = append(*farmersData, farmersTemp)
	}
}

func (db *Database) GetOrderByID(farmersData *[]utilites.Farmer) {
	rows, err := db.Pool.Query("SELECT * FROM farmers ORDER BY id")
	if err != nil {
		log.Printf("The request was made incorrectly: %s\n", err)
	}

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

		*farmersData = append(*farmersData, farmersTemp)
	}
}
