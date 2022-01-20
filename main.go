package main

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
	"library/Controller"
	"library/Model"
	"log"
	"net/http"
	"time"
)
func RoutersJson() {
	r := mux.NewRouter()
	r.HandleFunc("/book", Controller.NewHomeController).Methods("GET")//1
	r.HandleFunc("/book", Controller.NewSaveBookController).Methods("POST")//2
	r.HandleFunc("/reeder", Controller.NewSaveReederController).Methods("POST")//3
	r.HandleFunc("/reeders", Controller.NewGetReedersController).Methods("GET")//4
	r.HandleFunc("/rent", Controller.NewSaveRentController).Methods("POST")//5
	r.HandleFunc("/rentcomplete", Controller.NewCompleteRentController).Methods("PUT")
	r.HandleFunc("/rent", Controller.NewGetRentController).Methods("GET")
	r.HandleFunc("/cover", Controller.NewSaveCoverController).Methods("POST")
	r.HandleFunc("/cover", Controller.NewGetCoverController).Methods("GET")
	r.HandleFunc("/coverId", Controller.NewGetCoverController1).Methods("GET")
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
	r.HandleFunc("/complete_rent", Controller.CompleteRentController)//6+
	http.ListenAndServe("127.0.0.1:8080", r)
}

func main(){
	dt := time.Now()
	log.Printf("Current date and time is: ", dt.String())
	log.Printf("Start")
	/*err := Model.DownloadFile("https://kot-pes.com/wp-content/uploads/2019/03/post_5b48c1cfca497.jpg",fmt.Sprintf("C:/GO/booksCover/%s.jpg","cat"))
	if err != nil {
		log.Fatal(err)
	}  */
	err1 := Model.CallAt(18, 41, 0, Model.FindDebtorAndSendEmail)
	if err1 != nil {
		log.Println(err1)
	}

	RoutersJson()
	//RoutersTemplate()


}


