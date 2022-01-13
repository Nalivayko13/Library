package Model

import (
	_ "database/sql"
	"fmt"
	_ "fmt"
	_ "github.com/go-sql-driver/mysql"
	"library/dao"
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
func SaveReeder(reader *dao.Reeder){
	if reader.Name=="" || reader.Surname=="" || reader.DateOfBirth=="" || reader.Email=="" ||reader.Address=="" {
		fmt.Println("!!enter data")
	}
	dao.Save_reeder_toDB(*reader)
}

func GetReeders(AllReeders []dao.Reeder) []dao.Reeder{
	AllReeders = []dao.Reeder{}
	dao.Get_reeders_fronDB(&AllReeders)
	return AllReeders
}
