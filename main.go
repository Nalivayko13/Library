package main

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
	"library/Controller"
	"net/http"
	"time"
)

func main(){
	dt := time.Now()
	fmt.Println("Current date and time is: ", dt.String())
	r := mux.NewRouter()
	r.HandleFunc("/home", Controller.Home)
	r.HandleFunc("/addbook", Controller.AddBook)
	r.HandleFunc("/save_book", Controller.SaveBook)
	r.HandleFunc("/successful", Controller.Successful_book_reg)
	r.HandleFunc("/addreeder", Controller.AddReeder)
	r.HandleFunc("/reeders", Controller.GetReeders)
	r.HandleFunc("/save_reeder", Controller.SaveReeder)
	r.HandleFunc("/give_book", Controller.GiveBook)
	r.HandleFunc("/save_rent", Controller.SaveRent)


	http.ListenAndServe("127.0.0.1:8080", r)

}


