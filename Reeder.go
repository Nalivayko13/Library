package main

import (
	"database/sql"
	_ "database/sql"
	"fmt"
	_ "fmt"
	_ "github.com/go-sql-driver/mysql"
	 "html/template"
	"net/http"
	_ "net/http"
)
type Reeder struct {
	Id int `json:"id"`
	Name string `json:"name"`
	Surname string `json:"surname"`
	DateOfBirth string `json:"date_of_birth"`
	Address string `json:"address"`
	Email string `json:"email"`
}

var AllReeders=[]Reeder{}

func SaveReeder(w http.ResponseWriter, r *http.Request){

}

func AddReeder(w http.ResponseWriter,r *http.Request){
	t, err:=template.ParseFiles("template/addreeder.html")
	if err!=nil{
		fmt.Fprintf(w,err.Error())
	}
	t.Execute(w,nil)
}

func GetReeders(w http.ResponseWriter,r *http.Request){
	t, err:=template.ParseFiles("template/reeders.html")
	if err!=nil{
		fmt.Fprintf(w,err.Error())
	}

	db, err :=sql.Open("mysql","root:@tcp(127.0.0.1:3306)/library")
	if err!=nil {
		panic(err)
	}
	defer db.Close()

	res,err:=db.Query("SELECT * FROM `reeders`")
	if err!=nil{
		panic(err)
	}
	for res.Next(){
		var reeder Reeder
		err = res.Scan(&reeder.Id, &reeder.Name,&reeder.Surname, &reeder.DateOfBirth, &reeder.Address, &reeder.Email)
		if err!=nil{
			panic(err)
		}
		AllReeders = append(AllReeders, reeder)
	}

	t.Execute(w,AllReeders)

}

