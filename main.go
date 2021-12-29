package main

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
	"library/Controller"
	"log"
	"net/http"
	"time"
)

func main(){
	dt := time.Now()
	log.Printf("Current date and time is: ", dt.String())
	log.Printf("Start")
	r := mux.NewRouter()
	r.HandleFunc("/home", Controller.HomeController)
	r.HandleFunc("/addbook", Controller.AddBookController)
	r.HandleFunc("/save_book", Controller.SaveBookController)
	r.HandleFunc("/successful", Controller.Successful_book_reg)
	r.HandleFunc("/addreeder", Controller.AddReederController)
	r.HandleFunc("/reeders", Controller.GetReedersController)
	r.HandleFunc("/save_reeder", Controller.SaveReederController)
	r.HandleFunc("/give_book", Controller.GiveBook)
	r.HandleFunc("/save_rent", Controller.SaveRentController)
	r.HandleFunc("/returnbook", Controller.ReturnBookController)
	r.HandleFunc("/complete_rent", Controller.CompleteRentController)


	http.ListenAndServe("127.0.0.1:8080", r)

	/*
	srv := &http.Server{
		Addr:     *addr,
		ErrorLog: errorLog,
		Handler:  mux,
	}

	infoLog.Printf("Запуск сервера на %s", *addr)
	// Вызываем метод ListenAndServe() от нашей новой структуры http.Server
	err := srv.ListenAndServe()
	errorLog.Fatal(err)
	 */

}


