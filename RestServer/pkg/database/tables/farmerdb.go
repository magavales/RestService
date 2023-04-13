package tables

import (
	"RestService"
	"github.com/jackc/pgx"
	"log"
)

type FarmersDB struct {
	Data []RestService.Farmer
}

func (fdb *FarmersDB) GetAll(pool *pgx.Conn) {
	rows, err := pool.Query("SELECT * FROM farmers")
	if err != nil {
		log.Printf("The request was made incorrectly: %s\n", err)
	}

	for rows.Next() {
		value, err := rows.Values()
		if err != nil {
			log.Fatal("error while iterating dataset")
		}
		var farmersTemp RestService.Farmer

		farmersTemp.ParseData(value)

		fdb.Data = append(fdb.Data, farmersTemp)
	}
}

func (fdb *FarmersDB) GetOrderByID(pool *pgx.Conn) {
	rows, err := pool.Query("SELECT * FROM farmers ORDER BY id")
	if err != nil {
		log.Printf("The request was made incorrectly: %s\n", err)
	}

	for rows.Next() {
		value, err := rows.Values()
		if err != nil {
			log.Fatal("error while iterating dataset")
		}
		var farmersTemp RestService.Farmer

		farmersTemp.ParseData(value)

		fdb.Data = append(fdb.Data, farmersTemp)
	}
}
