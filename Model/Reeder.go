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

func GetReeders(AllReeders []dao.Reeder, limit,page int) ([]dao.Reeder,error){
	AllReeders = []dao.Reeder{}
	var Total = []dao.Reeder{}
	dao.Get_reedersWithPage_fronDB(&AllReeders,limit,page)
	dao.Get_reeders_fronDB(&Total)
	totalCount:=len(Total)
	LimitOfPages:=(totalCount/limit)+1
	if LimitOfPages<page{
		return nil,errors.New("No more pages")
	}
	if limit==0 || page==0 {
		//dao.Get_books_fronDB(&AllBooks)
		//log.Println("this is all books")
		return nil,errors.New("no page or limit")
	}
	return AllReeders,nil
}
