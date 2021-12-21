package main

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
	"library/Controller"
	"library/Model"
	"net/http"
	"time"
)

func main(){
	dt := time.Now()
	fmt.Println("Current date and time is: ", dt.String())
	r := mux.NewRouter()
	r.HandleFunc("/home", Model.Home)
	r.HandleFunc("/addbook", Model.AddBook)
	r.HandleFunc("/save_book", Model.SaveBook)
	r.HandleFunc("/successful", Controller.Successful_book_reg)
	r.HandleFunc("/addreeder", Controller.AddReeder)
	r.HandleFunc("/reeders", Controller.GetReeders)      //!!!!
	r.HandleFunc("/save_reeder", Controller.SaveReeder)  //!!!!

	http.ListenAndServe("127.0.0.1:8080", r)

}


