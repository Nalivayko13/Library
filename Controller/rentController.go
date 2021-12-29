package Controller

import (
	"fmt"
	"html/template"
	"library/dao"
	_ "library/dao"
	"net/http"
	"time"
)

func GiveBook(w http.ResponseWriter, r *http.Request){
	t, err:=template.ParseFiles("template/givebook.html")
	if err!=nil{
		fmt.Fprintf(w,err.Error())
	}
	t.Execute(w,nil)

}

func SaveRentController(w http.ResponseWriter, r *http.Request){
	//fmt.Fprintf(w,"click save_rent")
	var rent dao.Rent
	rent.Id_reeder=r.FormValue("id_reeder")
	rent.Id_book=r.FormValue("id_book")
	t:=time.Now()
	rent.First_date=t.Format("01-02-2006")
	rent.Last_date=t.Add(720*time.Hour).Format("01-02-2006")
	rent.Fine=0

	if rent.Id_reeder=="" || rent.Id_book=="" {
		fmt.Fprintf(w,"enter data")
	}
	dao.Save_rent_toDB(rent)
	http.Redirect(w,r,"/successful",http.StatusSeeOther)
}

func ReturnBookController(w http.ResponseWriter, r *http.Request)  {
	t, err:=template.ParseFiles("template/returnbook.html")
	if err!=nil{
		fmt.Fprintf(w,err.Error())
	}
	t.Execute(w,nil)
}
func CompleteRentController(w http.ResponseWriter,r *http.Request) {
	var rent dao.Rent
	rent.Id_rent=r.FormValue("id_rent")
	dao.Complete_rent_toDB(rent)
	http.Redirect(w,r,"/successful",http.StatusSeeOther)
}
