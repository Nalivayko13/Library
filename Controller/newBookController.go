package Controller

import (
	"encoding/json"
	"library/Model"
	"library/dao"
	"log"
	"net/http"
	"strconv"
	"time"
)

func NewHomeController(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	page,_:= strconv.Atoi(r.URL.Query().Get("page"))
	limit,_:= strconv.Atoi(r.URL.Query().Get("limit"))
	Books = Model.Home(Books,limit,page)
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
