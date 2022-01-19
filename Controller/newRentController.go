package Controller

import (
	"encoding/json"
	"library/Model"
	"library/dao"
	"log"
	"net/http"
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
	Model.SaveRent(rent)
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
	Rent = Model.GetRent(Rent)
	err := json.NewEncoder(w).Encode(Rent)
	if err != nil {
		log.Println(err)
	}
}