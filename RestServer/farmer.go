package RestService

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

func (farmer *Farmer) ParseData(values []interface{}) {
	farmer.ID = values[0].(int64)
	farmer.Name = values[1].(string)
	farmer.Surname = values[2].(string)
	farmer.Country = values[3].(string)
	farmer.DateOfBirth, _ = values[4].(time.Time)
	farmer.Email = values[5].(string)
	farmer.Village = values[6].(string)
	farmer.Land = values[7].(int32)
	farmer.AvgIncome = values[8].(string)
}

func (farmer *Farmer) PrintFarmers() {
	farmersJSON, err := json.MarshalIndent(farmer, "", " ")
	if err != nil {
		log.Fatal(err.Error())
	}
	fmt.Printf("%s\n", string(farmersJSON))
}

func (farmer *Farmer) ToJSON() []byte {
	body, err := json.Marshal(farmer)
	if err != nil {
		log.Printf("Convertation has been stopped!: %s", err)
	}

	return body
}
