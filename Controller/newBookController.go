package Controller

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
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
	Books,errr:= Model.Home(Books,limit,page)
	if errr!=nil{
		b, _ := json.Marshal(fmt.Sprintf("Error: %s",errr))
		//json.NewEncoder(w).Encode(b)
		w.Write(b)
	}
	err := json.NewEncoder(w).Encode(Books)
	if err != nil {
		log.Println(w, http.StatusBadRequest)
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
	err1:=Model.SaveBook(&book)
	if err1 != nil {
		b, _ := json.Marshal(fmt.Sprintf("Error: %s",err1))
		//json.NewEncoder(w).Encode(b)
		w.Write(b)
		log.Println(w, http.StatusBadRequest)
	}
}

func NewSaveCoverController(w http.ResponseWriter,r *http.Request) {
	if r.Body == nil {
		w.WriteHeader(http.StatusBadRequest)
	}
	defer r.Body.Close()

	cover, err := ioutil.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusUnprocessableEntity)
	}
	Model.SaveCoverFile(cover)
}

func NewGetCoverController(w http.ResponseWriter,r *http.Request) {
	name:= r.URL.Query().Get("name")
	fContent, err := ioutil.ReadFile(fmt.Sprintf("C:/GO/booksCover/%s.jpg",name))
	if err != nil {
		log.Println(err)
	}
    w.Write(fContent)
}

func NewGetCoverController1(w http.ResponseWriter,r *http.Request) {
	id:= r.URL.Query().Get("id")
	fContent, err:=Model.GetCoverById(id)
	if err != nil {
		json.NewEncoder(w).Encode(fmt.Sprintf("Error: %s",err))
	}
	json.NewEncoder(w).Encode(fContent)
}