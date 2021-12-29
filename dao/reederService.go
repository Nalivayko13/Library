package dao

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	_ "github.com/gorilla/mux"
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
