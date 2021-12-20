package main

import (
"database/sql"
	"time"

	//"database/sql"
"fmt"
_ "github.com/go-sql-driver/mysql"
"github.com/gorilla/mux"
"html/template"
"net/http"
)

func main(){
	dt := time.Now()
	fmt.Println("Current date and time is: ", dt.String())
	r := mux.NewRouter()
	r.HandleFunc("/home", Home)
	r.HandleFunc("/addbook", AddBook)
	r.HandleFunc("/save_book",SaveBook)
	r.HandleFunc("/successful",Successful_book_reg)
	r.HandleFunc("/addreeder",AddReeder)
	r.HandleFunc("/reeders",GetReeders)
	r.HandleFunc("/save_reeder",SaveReeder)

	//http.Handle("/", r)
	http.ListenAndServe("127.0.0.1:8080", r)

}


func Home(w http.ResponseWriter,r *http.Request){
	t, err:=template.ParseFiles("template/home.html")
	if err!=nil{
		fmt.Fprintf(w,err.Error())
	}

	db, err :=sql.Open("mysql","root:@tcp(127.0.0.1:3306)/library")
	if err!=nil {
		panic(err)
	}
	defer db.Close()

	res,err:=db.Query("SELECT * FROM `books`")
	if err!=nil{
		panic(err)
	}
	AllBooks = []Book{}
	for res.Next(){
		var book Book
		err = res.Scan(&book.IdBook, &book.Name,&book.Genre, &book.Price_of_book, &book.Num_of_copies, &book.Authors, &book.Cover_photo, &book.Price_of_book,&book.Reg_date)
		if err!=nil{
			panic(err)
		}
	 	AllBooks = append(AllBooks, book)
	}

	t.Execute(w,AllBooks)

}


