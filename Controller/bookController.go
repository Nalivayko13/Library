package Controller

import (
	"fmt"
	"html/template"
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

func SaveBook(w http.ResponseWriter, r *http.Request){

	var book dao.Book
	book.Name=r.FormValue("name")
	book.Genre=r.FormValue("genre")
	book.Authors=r.FormValue("authors")
	book.Num_of_copies=r.FormValue("num_of_copies")
	book.Price_per_day=r.FormValue("price_per_day")
	book.Reg_date=time.Now().Format("01-02-2006")
	book.Price_of_book=r.FormValue("price_of_book")

	_,err1:=strconv.Atoi(book.Price_of_book)
	if err1!=nil{
		fmt.Fprintf(w," or data is incorrect")
	}
	_,err2:=strconv.Atoi(book.Price_per_day)
	if err2!=nil{
		fmt.Fprintf(w," or data is incorrect")
	}
	_,err3:=strconv.Atoi(book.Num_of_copies)
	if err3!=nil{
		fmt.Fprintf(w," or data is incorrect")
	}

	if book.Name=="" || book.Price_of_book=="" || book.Price_per_day=="" || book.Num_of_copies=="" || book.Reg_date=="" || book.Genre=="" || book.Authors==""{
		fmt.Fprintf(w,"enter data")
	}
	dao.Save_book_toDB(book)
	http.Redirect(w,r,"/successful",http.StatusSeeOther)

}
func AddBook(w http.ResponseWriter, r *http.Request){
	t, err:=template.ParseFiles("template/addbook.html")
	if err!=nil{
		fmt.Fprintf(w,err.Error())
	}
	t.Execute(w,nil)

}
var AllBooks = []dao.Book{}
func Home(w http.ResponseWriter,r *http.Request){
	t, err:=template.ParseFiles("template/home.html")
	if err!=nil{
		fmt.Fprintf(w,err.Error())
	}
	AllBooks = []dao.Book{}
	dao.Get_books_fronDB(&AllBooks)
	t.Execute(w,AllBooks)

}

