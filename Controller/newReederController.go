package Controller

import (
	"encoding/json"
	"fmt"
	"library/Model"
	"library/dao"
	"log"
	"net/http"
)

func NewSaveReederController(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	var t dao.Reeder
	err := decoder.Decode(&t)
	if err != nil {
		log.Println(err)
	}
	err1:=Model.SaveReeder(&t)
	if err1 != nil {
		b, _ := json.Marshal(fmt.Sprintf("Error: %s",err1))
		//json.NewEncoder(w).Encode(b)
		w.Write(b)
	}
}

func NewGetReedersController(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	AllReeders = Model.GetReeders(AllReeders)
	err := json.NewEncoder(w).Encode(AllReeders)
	if err != nil {
		log.Println(err)
	}
}
