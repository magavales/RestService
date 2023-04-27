package RestClient

import (
	"fmt"
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

func (f *Farmer) Print() {
	fmt.Printf("ID: %d\n", f.ID)
	fmt.Printf("Name: %s\n", f.Name)
	fmt.Printf("Surname: %s\n", f.Surname)
	fmt.Printf("Country: %s\n", f.Country)
	fmt.Printf("Date of birth: %s\n", f.DateOfBirth.Format("2006-01-02"))
	fmt.Printf("Email: %s\n", f.Email)
	fmt.Printf("Village: %s\n", f.Village)
	fmt.Printf("Land: %d\n", f.Land)
	fmt.Printf("Average Income: %s\n", f.AvgIncome)
}
