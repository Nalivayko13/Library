package Model

import (
	_ "database/sql"
	"errors"
	"fmt"
	_ "fmt"
	_ "github.com/go-sql-driver/mysql"
	"library/dao"
	"log"
	_ "net/http"
)
type Reeder struct {
	IdReeder int `json:"id_reeder"`
	Name string `json:"name"`
	Surname string `json:"surname"`
	DateOfBirth string `json:"date_of_birth"`
	Address string `json:"address"`
	Email string `json:"email"`
}
func SaveReeder(reader *dao.Reeder) error{
	if reader.Name=="" || reader.Surname=="" || reader.DateOfBirth=="" || reader.Email=="" ||reader.Address=="" {
		fmt.Println("!!enter data")
		return errors.New("No data")
	}
	var emails []string
	dao.Get_Allemails_fromDB(&emails)
	correct:=0
	for _,e := range emails{
		if e==reader.Email{
			correct++
		}
	}
	if correct!=0 {
		log.Println("this email already exists")
		return errors.New("this email already exists")
	}

	dao.Save_reeder_toDB(*reader)
	return nil
}

func GetReeders(AllReeders []dao.Reeder) []dao.Reeder{
	AllReeders = []dao.Reeder{}
	dao.Get_reeders_fronDB(&AllReeders)
	return AllReeders
}
