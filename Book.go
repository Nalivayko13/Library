package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	_ "github.com/gorilla/mux"
	"html/template"
	"net/http"
	"strconv"
	"time"
)

type Book struct {
	IdBook int `json:"id_book"`
	Name string `json:"name"`
	Genre string `json:"genre"`
	Price_of_book string `json:"price_of_book"`
	Num_of_copies string `json:"num_of_copies"`
	Authors string `json:"authors"`
	Cover_photo string `json:"cover_photo"`
	Price_per_day string `json:"price_per_day"`
	Reg_date string `json:"reg_date"`
}

var AllBooks = []Book{}

func Successful_book_reg(w http.ResponseWriter, r *http.Request){
	t, err:=template.ParseFiles("template/successful.html")
	if err!=nil{
		fmt.Fprintf(w,err.Error())
	}
	t.Execute(w,nil)

}

func SaveBook(w http.ResponseWriter, r *http.Request){
	db, err :=sql.Open("mysql","root:@tcp(127.0.0.1:3306)/library")
	if err!=nil {
		panic(err)
	}
	defer db.Close()

	var book Book
	book.Name=r.FormValue("name")
	book.Genre=r.FormValue("genre")
	book.Authors=r.FormValue("authors")
	book.Num_of_copies=r.FormValue("num_of_copies")
	book.Price_per_day=r.FormValue("price_per_day")
	book.Reg_date=time.Now().Format("01-02-2006")
	book.Price_of_book=r.FormValue("price_of_book")

	_,err1:=strconv.Atoi(book.Price_of_book)
	_,err2:=strconv.Atoi(book.Price_per_day)
	_,err3:=strconv.Atoi(book.Num_of_copies)

	if book.Name=="" || book.Price_of_book=="" || book.Price_per_day=="" || book.Num_of_copies=="" || book.Reg_date=="" || book.Genre=="" || book.Authors==""{
		fmt.Fprintf(w,"enter data")
	}
	if err1==nil && err2==nil && err3==nil{

		s := fmt.Sprintf("INSERT INTO `books` (`name`,`genre`,`price_of_book`,`num_of_copies`,`authors`, `price_per_day`,`reg_date`) VALUES ('%s', '%s', '%s', '%s', '%s', '%s', '%s')", book.Name, book.Genre, book.Price_of_book, book.Num_of_copies,book.Authors,book.Price_per_day,book.Reg_date)
		fmt.Println(s)
		insert, err1 := db.Query(s)
		if err1 != nil {
			panic(err1)
		}

		fmt.Println(insert)
		defer insert.Close()
		http.Redirect(w,r,"/successful",http.StatusSeeOther)
	}else {
		fmt.Fprintf(w," or data is incorrect")
	}
	http.Redirect(w,r,"/home",http.StatusSeeOther)
}

func AddBook(w http.ResponseWriter, r *http.Request){
	t, err:=template.ParseFiles("template/addbook.html")
	if err!=nil{
		fmt.Fprintf(w,err.Error())
	}
	t.Execute(w,nil)

}