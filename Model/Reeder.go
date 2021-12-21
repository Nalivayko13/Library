package Model

import (
	_ "database/sql"
	"fmt"
	_ "fmt"
	_ "github.com/go-sql-driver/mysql"
	"html/template"
	"library/dao"
	"net/http"
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

var AllReeders=[]dao.Reeder{}
func SaveReeder(w http.ResponseWriter, r *http.Request){
	var reeder dao.Reeder
	reeder.Name=r.FormValue("name")
	reeder.Surname=r.FormValue("surname")
	reeder.DateOfBirth=r.FormValue("date_of_birth")
	reeder.Email=r.FormValue("email")
	reeder.Address=r.FormValue("address")
	if reeder.Name=="" || reeder.Surname=="" || reeder.DateOfBirth=="" || reeder.Email=="" ||reeder.Address=="" {
		fmt.Fprintf(w,"!!enter data")
	}
	dao.Save_reeder_toDB(reeder)
	http.Redirect(w,r,"/successful",http.StatusSeeOther)

}

func GetReeders(w http.ResponseWriter,r *http.Request){
	t, err:=template.ParseFiles("template/reeders.html")
	if err!=nil{
		fmt.Fprintf(w,err.Error())
	}
	AllReeders = []dao.Reeder{}
	dao.Get_reeders_fronDB(&AllReeders)
	t.Execute(w, AllReeders)

}

