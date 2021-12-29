package Model

import (
	"fmt"
	"library/dao"
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
	dao.Complete_rent_toDB(rent)
}