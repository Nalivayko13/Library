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
	db, err :=sql.Open("mysql","root:@tcp(127.0.0.1:3306)/library")
	if err!=nil {
		panic(err)
	}
	defer db.Close()

	var reeder Reeder
	reeder.Name=r.FormValue("name")
	reeder.Surname=r.FormValue("surname")
	reeder.DateOfBirth=r.FormValue("date_of_birth")
	reeder.Email=r.FormValue("email")
	reeder.Address=r.FormValue("address")



	if reeder.Name=="" || reeder.Surname=="" || reeder.DateOfBirth=="" || reeder.Email=="" ||reeder.Address=="" {
		fmt.Fprintf(w,"enter data")
	}
	s := fmt.Sprintf("INSERT INTO `reeders` (`name`,`surname`,`date_of_birth`, `address`, `email`) VALUES ('%s', '%s', '%s', '%s', '%s')", reeder.Name, reeder.Surname,reeder.DateOfBirth,reeder.Address,reeder.Email)
	fmt.Println(s)
	insert, err1 := db.Query(s)
	if err1 != nil {
		panic(err1)
	}

	fmt.Println(insert)
	defer insert.Close()
	http.Redirect(w,r,"/successful",http.StatusSeeOther)
	//http.Redirect(w,r,"/home",http.StatusSeeOther)
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
	AllReeders = []Reeder{}
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

