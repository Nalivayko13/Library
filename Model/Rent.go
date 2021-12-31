package Model

import (
	"fmt"
	"library/dao"
	"log"
	"strconv"
	"time"
)

type Rent struct {
	Id_rent string `json:"id_rent"`
	Id_book string `json:"id_book"`
	Id_reeder string `json:"id_reeder"`
	First_date string `json:"first_date"`
	Last_date string `json:"last_date"`
	Fine float64 `json:"fine"`
}
func SaveRent(rent dao.Rent){

	if rent.Id_reeder=="" || rent.Id_book=="" {
		fmt.Println("enter data")
	}
	dao.Save_rent_toDB(rent)
}

func CompleteRent(rent dao.Rent){
	dao.Get_rent_byId(&rent)
	CountFine(&rent)
	fmt.Println(rent.Fine)
	dao.Complete_rent_toDB(rent)
}

func CountFine(rent *dao.Rent) {
	t1,_:=time.Parse("01-02-2006", rent.Last_date)
	t2:=time.Now()
	log.Println(t1)
	log.Println(t2)
	if t2.After(t1) {
		price:=dao.Get_priceOfbook_byId(rent.Id_book)
		p1,_:=strconv.Atoi(price)
		p2:=float64(p1)
		day:=(t2.Sub(t1).Hours())/24
		fine:= p2*(0.01)*day
		rent.Fine=fine
	}else { rent.Fine=0}
}