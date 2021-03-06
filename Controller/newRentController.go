package Controller

import (
	"encoding/json"
	"fmt"
	"library/Model"
	"library/dao"
	"log"
	"net/http"
	"strconv"
	"time"
)

func NewSaveRentController(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	var rent dao.Rent
	t:=time.Now()
	rent.First_date=t.Format("01-02-2006")
	rent.Last_date=t.Add(720*time.Hour).Format("01-02-2006")
	rent.Fine="0"
	err := decoder.Decode(&rent)
	if err != nil {
		log.Println(err)
	}
	err1:=Model.SaveRent(&rent)
	if err1 != nil {
		b, _ := json.Marshal(fmt.Sprintf("Error: %s",err1))
		//json.NewEncoder(w).Encode(b)
		w.Write(b)
	}
}

func NewCompleteRentController(w http.ResponseWriter,r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	decoder := json.NewDecoder(r.Body)
	var rent dao.Rent
	err := decoder.Decode(&rent)
	if err != nil {
		log.Println(err)
	}
	Model.CompleteRent(rent)
}
var Rent = []dao.Rent{}
func NewGetRentController(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	page,_:= strconv.Atoi(r.URL.Query().Get("page"))
	limit,_:= strconv.Atoi(r.URL.Query().Get("limit"))
	Rent,errr := Model.GetRent(Rent,limit,page)
	if errr!=nil{
		b, _ := json.Marshal(fmt.Sprintf("Error: %s",errr))
		//json.NewEncoder(w).Encode(b)
		w.Write(b)
	}
	err := json.NewEncoder(w).Encode(Rent)
	if err != nil {
		log.Println(err)
	}
}