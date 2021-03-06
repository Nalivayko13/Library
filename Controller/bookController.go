package Controller

import (
	"fmt"
	"html/template"
	"library/Model"
	"library/dao"
	"net/http"
	"strconv"
	"time"
)

func Successful_book_reg(w http.ResponseWriter, r *http.Request){
	t, err:=template.ParseFiles("template/successful.html")
	if err!=nil{
		fmt.Fprintf(w,err.Error())
	}
	t.Execute(w,nil)

}

func SaveBookController(w http.ResponseWriter, r *http.Request){

	var book dao.Book
	book.Name=r.FormValue("name")
	//book.Genre=r.FormValue("genre")
	book.Authors=r.FormValue("authors")
	book.Num_of_copies=r.FormValue("num_of_copies")
	book.Price_per_day=r.FormValue("price_per_day")
	book.Reg_date=time.Now().Format("01-02-2006")
	book.Price_of_book=r.FormValue("price_of_book")
	Model.SaveBook(&book)
	http.Redirect(w,r,"/successful",http.StatusSeeOther)

}
func AddBookController(w http.ResponseWriter, r *http.Request){
	t, err:=template.ParseFiles("template/addbook.html")
	if err!=nil{
		fmt.Fprintf(w,err.Error())
	}
	t.Execute(w,nil)

}
var AllBooks = []dao.Book{}
func HomeController(w http.ResponseWriter,r *http.Request){
	t, err:=template.ParseFiles("template/home.html")
	if err!=nil{
		fmt.Fprintf(w,err.Error())
	}
	limit,_:= strconv.Atoi(r.URL.Query().Get("limit"))
	page,_:= strconv.Atoi(r.URL.Query().Get("page"))
	AllBooks,_=Model.Home(AllBooks,limit,page)
	t.Execute(w,AllBooks)

}

