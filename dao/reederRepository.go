package dao

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	_ "github.com/gorilla/mux"
	"log"
)
type Reeder struct {
	IdReeder int `json:"id_reeder"`
	Name string `json:"name"`
	Surname string `json:"surname"`
	DateOfBirth string `json:"date_of_birth"`
	Address string `json:"address"`
	Email string `json:"email"`
}
func Save_reeder_toDB(reeder Reeder){
	db := openDB()
	defer db.Close()

	s := fmt.Sprintf("INSERT INTO `reeders` (`name`,`surname`,`date_of_birth`, `address`, `email`) VALUES ('%s', '%s', '%s', '%s', '%s')", reeder.Name, reeder.Surname,reeder.DateOfBirth,reeder.Address,reeder.Email)
	fmt.Println(s)
	insert, err1 := db.Query(s)
	if err1 != nil {
		panic(err1)
	}

	fmt.Println(insert)
	defer insert.Close()

}
func Get_reeders_fronDB(AllReeders *[]Reeder){
	db := openDB()
	defer db.Close()
	res,err:=db.Query("SELECT * FROM `reeders`")
	if err!=nil{
		panic(err)
	}
	for res.Next(){
		var reeder Reeder
		err = res.Scan(&reeder.IdReeder,&reeder.Name,&reeder.Surname, &reeder.DateOfBirth, &reeder.Address, &reeder.Email)
		if err!=nil{
			panic(err)
		}
		*AllReeders = append(*AllReeders, reeder)
	}
}
func Get_email_byId(idReeder string) string {
	db:=openDB()
	defer db.Close()
	s := fmt.Sprintf("SELECT `email` FROM `reeders` WHERE `id_reeder` = '%s'", idReeder)
	res,err:=db.Query(s)
	if err!=nil{
		panic(err)
	}
	email:=""
	for res.Next(){
		err = res.Scan(&email)
		if err!=nil{
			panic(err)
		}
	}
	log.Println("email is  ", email)
	return email
}

func Get_AllreedersID_fromDB(idReeders *[]int) {
	db := openDB()
	defer db.Close()
	res,err:=db.Query("SELECT `id_reeder` FROM `reeders`")
	if err!=nil{
		panic(err)
	}
	for res.Next(){
		var reederID int
		err = res.Scan(&reederID)
		*idReeders = append (*idReeders, reederID)
	}
}

func Get_reedersWithPage_fronDB(AllReeders *[]Reeder,limit, page int){
	db := openDB()
	defer db.Close()
	//TODO change limit(20)
	res,err:=db.Query(fmt.Sprintf(  "SELECT * FROM `reeders` ORDER BY `name` LIMIT %d OFFSET %d",
		limit,limit*(page-1)))
	if err!=nil{
		panic(err)
	}
	for res.Next(){
		var reeder Reeder
		err = res.Scan(&reeder.IdReeder,&reeder.Name,&reeder.Surname, &reeder.DateOfBirth, &reeder.Address, &reeder.Email)
		if err!=nil{
			panic(err)
		}
		*AllReeders = append(*AllReeders, reeder)
	}

}

func Get_Allemails_fromDB(emails *[]string) {
	db := openDB()
	defer db.Close()
	res,err:=db.Query("SELECT `email` FROM `reeders`")
	if err!=nil{
		panic(err)
	}
	for res.Next(){
		var email string
		err = res.Scan(&email)
		*emails = append (*emails, email)
	}
}
