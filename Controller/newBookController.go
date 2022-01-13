package Controller

import (
	"encoding/json"
	"library/Model"
	"library/dao"
	"log"
	"net/http"
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
	w.Header().Set("Content-Type", "application/json")
	var book dao.Book
	err:=json.NewDecoder(r.Body).Decode(&book)
	if err!=nil{
		log.Fatal(err)
	}
	Model.SaveBook(&book)
}
