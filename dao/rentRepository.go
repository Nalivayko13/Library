package dao

import (
	"fmt"
	"log"
)

type Rent struct {
	Id_rent string `json:"id_rent" db:"id_rent"`
	Id_book string `json:"id_book" db:"id_book"`
	Id_reeder string `json:"id_reeder" db:"id_reeder"`
	First_date string `json:"first_date"`
	Last_date string `json:"last_date"`
	Fine string `json:"fine"`
	Comlete string `json:"comlete"`
}

func Save_rent_toDB(rent Rent){
	fmt.Println(rent)
	db := openDB()
	defer db.Close()

	s := fmt.Sprintf("INSERT INTO `rent` (`id_book`,`id_reeder`,`first_date`, `last_date`, `fine`) VALUES ('%s', '%s', '%s', '%s', %v)", rent.Id_book, rent.Id_reeder,rent.First_date,rent.Last_date,rent.Fine)
	fmt.Println(s)
	insert, err1 := db.Query(s)
	if err1 != nil {
		panic(err1)
	}

	fmt.Println(insert)
	defer insert.Close()

}

func Complete_rent_toDB(rent Rent){
	db := openDB()
	defer db.Close()

	s := fmt.Sprintf("UPDATE `rent` SET `completed` = 'completed',`fine` = '%s' WHERE `id_rent` = '%s'",rent.Fine,rent.Id_rent)
	log.Println(s)
	insert, err1 := db.Query(s)
	if err1 != nil {
		panic(err1)
	}
	defer insert.Close()
}
func Get_rent_byId(rent *Rent){
	db:=openDB()
	defer db.Close()
	s := fmt.Sprintf("SELECT `id_book`,`id_reeder`,`first_date`,`last_date` FROM `rent` WHERE `id_rent` = '%s'", rent.Id_rent)
	res,err:=db.Query(s)
	if err!=nil{
		panic(err)
	}

	for res.Next(){
		err = res.Scan(&rent.Id_book,&rent.Id_reeder,&rent.First_date,&rent.Last_date)
		if err!=nil{
			panic(err)
		}
	}
}

func Get_AllRent_fromDB(rent *[]Rent){
	db := openDB()
	defer db.Close()
	res,err:=db.Query("SELECT * FROM `rent`")
	if err!=nil{
		panic(err)
	}
	for res.Next(){
		var r Rent
		err = res.Scan(&r.Id_rent, &r.Id_book, &r.Id_reeder, &r.First_date, &r.Last_date,&r.Fine, &r.Comlete)
		if err!=nil{
			panic(err)
		}
		*rent = append(*rent, r)
	}
}

func Get_AllRentWithPage_fromDB(rent *[]Rent, limit,page int){
	db := openDB()
	defer db.Close()
	res,err:=db.Query(fmt.Sprintf("SELECT * FROM `rent` ORDER BY `id_rent` LIMIT %d OFFSET %d",limit,limit*(page-1)))
	if err!=nil{
		panic(err)
	}
	for res.Next(){
		var r Rent
		err = res.Scan(&r.Id_rent, &r.Id_book, &r.Id_reeder, &r.First_date, &r.Last_date,&r.Fine, &r.Comlete)
		if err!=nil{
			panic(err)
		}
		*rent = append(*rent, r)
	}
}