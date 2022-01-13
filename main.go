package main

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
	"library/Controller"
	"log"
	"net/http"
	"time"
)
func RoutersJson() {
	r := mux.NewRouter()
	r.HandleFunc("/home", Controller.NewHomeController)//1
	//r.HandleFunc("/addbook", Controller.NewSaveBookController)//2
	r.HandleFunc("/save_reeder", Controller.NewSaveReederController)//3
	r.HandleFunc("/reeders", Controller.NewGetReedersController)//4
	r.HandleFunc("/save_rent", Controller.NewSaveRentController)//5
	r.HandleFunc("/complete_rent", Controller.NewCompleteRentController)//6

	http.ListenAndServe("127.0.0.1:8080", r)
}
func RoutersTemplate() {
	r := mux.NewRouter()
	r.HandleFunc("/home", Controller.HomeController)//1+
	r.HandleFunc("/addbook", Controller.AddBookController)//no
	r.HandleFunc("/save_book", Controller.SaveBookController)//2-
	r.HandleFunc("/successful", Controller.Successful_book_reg)//no
	r.HandleFunc("/addreeder", Controller.AddReederController)//no
	r.HandleFunc("/reeders", Controller.GetReedersController)//3+
	r.HandleFunc("/save_reeder", Controller.SaveReederController)//4+
	r.HandleFunc("/give_book", Controller.GiveBook)//no
	r.HandleFunc("/save_rent", Controller.SaveRentController)//5+
	r.HandleFunc("/returnbook", Controller.ReturnBookController)//no
	r.HandleFunc("/complete_rent", Controller.CompleteRentController)//6-
	http.ListenAndServe("127.0.0.1:8080", r)
}

func main(){
	dt := time.Now()
	log.Printf("Current date and time is: ", dt.String())
	log.Printf("Start")

	RoutersJson()
	//RoutersTemplate()



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


