package Model

import (
	"errors"
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
func SaveRent(rent dao.Rent) error{

	if rent.Id_reeder=="" || rent.Id_book=="" {
		fmt.Println("enter data")
		simpError:=errors.New("No data")
		return simpError
		
	}
	correct:=0
	var idBook []int
	dao.Get_AllbookID_fromDB(&idBook)
	var idReeder []int
	dao.Get_AllreedersID_fromDB(&idReeder)
	for _,idB := range idBook{
		id:=strconv.Itoa(idB)
		if id==rent.Id_book{
			correct++
		}
	}
	for _,idR := range idReeder{
		id:=strconv.Itoa(idR)
		if id==rent.Id_book{
			correct++
		}
	}
	if correct==0 {
	log.Println("no such id exists")
		return errors.New("no such id exists")
	}

	numOfCopies:=dao.Get_numOfCopies_fromDB(rent.Id_book)-1
	fmt.Println(numOfCopies)
	if numOfCopies<0 {
		log.Println("out of copies of the book...")
		simpError:=errors.New("out of copies of the book...")
		return simpError
	}


	str:=strconv.Itoa(numOfCopies)
	dao.Set_numOfCopies_fromDB(str, rent.Id_book)
	dao.Save_rent_toDB(rent)
	return nil
}

func CompleteRent(rent dao.Rent){
	dao.Get_rent_byId(&rent)
	numOfCopies:=dao.Get_numOfCopies_fromDB(rent.Id_book)+1
	str:=strconv.Itoa(numOfCopies)
	dao.Set_numOfCopies_fromDB(str,rent.Id_book)
	CountFine(&rent)
	fmt.Println(rent.Fine)
	fmt.Println(rent.Id_book)
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
		s := fmt.Sprintf("%f", fine)
		rent.Fine=s
	}else { rent.Fine=strconv.Itoa(0)}
}

func GetRent(rent []dao.Rent) []dao.Rent{
	rent = []dao.Rent{}
	dao.Get_AllRent_fromDB(&rent)
	return rent
}