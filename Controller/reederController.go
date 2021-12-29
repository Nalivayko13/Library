package Controller

import (
	"fmt"
	"html/template"
	"library/Model"
	"library/dao"
	"net/http"

)

func AddReederController(w http.ResponseWriter,r *http.Request){
	t, err:=template.ParseFiles("template/addreeder.html")
	if err!=nil{
		fmt.Fprintf(w,err.Error())
	}
	t.Execute(w,nil)
}

func SaveReederController(w http.ResponseWriter, r *http.Request){
	var reeder dao.Reeder
	reeder.Name=r.FormValue("name")
	reeder.Surname=r.FormValue("surname")
	reeder.DateOfBirth=r.FormValue("date_of_birth")
	reeder.Email=r.FormValue("email")
	reeder.Address=r.FormValue("address")
	Model.SaveReeder(&reeder)
	http.Redirect(w,r,"/successful",http.StatusSeeOther)
}

var AllReeders=[]dao.Reeder{}
func GetReedersController(w http.ResponseWriter,r *http.Request){
	t, err:=template.ParseFiles("template/reeders.html")
	if err!=nil{
		fmt.Fprintf(w,err.Error())
	}
	AllReeders = Model.GetReeders(AllReeders)
	t.Execute(w, AllReeders)

}