package Controller

import (
	"fmt"
	"html/template"
	"net/http"
)

func Successful_book_reg(w http.ResponseWriter, r *http.Request){
	t, err:=template.ParseFiles("template/successful.html")
	if err!=nil{
		fmt.Fprintf(w,err.Error())
	}
	t.Execute(w,nil)

}
