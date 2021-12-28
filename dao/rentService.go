package dao

import (
	"database/sql"
	"fmt"
)

type Rent struct {
	Id_rent string
	Id_book string
	Id_reeder string
	First_date string
	Last_date string
	Fine float64
}

func Save_rent_toDB(rent Rent){
	db, err :=sql.Open("mysql","root:@tcp(127.0.0.1:3306)/library")
	if err!=nil {
		panic(err)
	}
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
	db, err :=sql.Open("mysql","root:@tcp(127.0.0.1:3306)/library")
	if err!=nil {
		panic(err)
	}
	defer db.Close()

	s := fmt.Sprintf("UPDATE `rent` SET `completed` = 'completed' WHERE `id_rent` = '%s'",rent.Id_rent)
	fmt.Println(s)
	insert, err1 := db.Query(s)
	if err1 != nil {
		panic(err1)
	}

	fmt.Println(insert)
	defer insert.Close()
}