package Controller

import (
	"fmt"
	"html/template"
	"library/Model"
	"library/dao"
	"net/http"
	"strconv"
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
	limit,_:= strconv.Atoi(r.URL.Query().Get("limit"))
	page,_:= strconv.Atoi(r.URL.Query().Get("page"))
	AllReeders,_ = Model.GetReeders(AllReeders,limit,page)
	t.Execute(w, AllReeders)

}