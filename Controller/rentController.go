package Controller

import (
	"fmt"
	"html/template"
	_ "library/dao"
	"net/http"
)

func GiveBook(w http.ResponseWriter, r *http.Request){
	t, err:=template.ParseFiles("template/givebook.html")
	if err!=nil{
		fmt.Fprintf(w,err.Error())
	}
	t.Execute(w,nil)

}

func SaveRent(w http.ResponseWriter, r *http.Request){
	fmt.Fprintf(w,"click save_rent")
}
