package utilites

import (
	"encoding/json"
	"fmt"
	"log"
	"time"
)

type Farmer struct {
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

func (farmer *Farmer) PrintFarmers() {
	farmersJSON, err := json.MarshalIndent(farmer, "", " ")
	if err != nil {
		log.Fatal(err.Error())
	}
	fmt.Printf("%s\n", string(farmersJSON))
}
