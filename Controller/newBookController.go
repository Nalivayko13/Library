package Controller

import (
	"encoding/json"
	"library/Model"
	"library/dao"
	"log"
	"net/http"
	"time"
)

func NewHomeController(w http.ResponseWriter, r *http.Request) {
	Books = Model.Home(Books)
	err := json.NewEncoder(w).Encode(Books)
	if err != nil {
		log.Println(err)
	}
}
var Books = []dao.Book{}
func NewSaveBookController(w http.ResponseWriter, r *http.Request) {
	var book dao.Book
	book.Reg_date=time.Now().Format("01-02-2006")
	err := json.NewDecoder(r.Body).Decode(&book)
	if err != nil {
		log.Println(err)
	}
	Model.SaveBook(&book)
}
