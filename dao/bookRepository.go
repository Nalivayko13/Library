package dao

import (
	_ "database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	_ "github.com/gorilla/mux"
	_ "github.com/jmoiron/sqlx"
	"log"
	"strconv"
)

type Book struct {
	IdBook int `json:"id_book" db:"id_book"`
	Name string `json:"name" db:"name" `
	Genre []Genre `json:"genre" db:"genre"`
	Price_of_book string `json:"price_of_book" db:"price_of_book"`
	Num_of_copies string `json:"num_of_copies" db:"num_of_copies"`
	Authors string `json:"authors"`
	Cover_photo string `json:"cover_photo"`
	Price_per_day string `json:"price_per_day"`
	Reg_date string `json:"reg_date"`
}
func Save_book_toDB(Mbook Book) int  {
	var book Book
	book=Mbook
	db := openDB()
	defer db.Close()
/*	s := fmt.Sprintf("INSERT INTO `books` (`name`,`price_of_book`,`num_of_copies`,`authors`, `price_per_day`,`reg_date`,`cover_photo`) VALUES ('%s', '%s', '%s', '%s', '%s', '%s','%s' )", book.Name, book.Price_of_book, book.Num_of_copies,book.Authors,book.Price_per_day,book.Reg_date,book.Cover_photo)
	fmt.Println(s)
	insert, err1 := db.Query(s)
	if err1 != nil {
		panic(err1)
	}

	fmt.Println(insert)
	defer insert.Close()    */
	stmt, err := db.Prepare("INSERT INTO `books` SET `name`=?,`price_of_book`=?,`num_of_copies`=?,`authors`=?, `price_per_day`=?,`reg_date`=?,`cover_photo`=?")
	if err != nil {
		return -1
	}
	defer stmt.Close()

	res, err := stmt.Exec(book.Name,book.Price_of_book, book.Num_of_copies, book.Authors, book.Price_per_day, book.Reg_date, book.Cover_photo)
	if err != nil {
		return -1
	}
	id,_:=res.LastInsertId()
	return int(id)

}

func Get_booksWithPage_fronDB(AllBooks *[]Book,limit, page int){
	db := openDB()
	defer db.Close()
	//TODO change limit(20)
	res,err:=db.Query(fmt.Sprintf(  "SELECT `id_book`,`name`,`price_of_book`,`num_of_copies`,`authors`,`cover_photo`,`price_of_book`,`reg_date` FROM `books` ORDER BY `id_book` LIMIT %d OFFSET %d",
		limit,limit*(page-1)))
	if err!=nil{
		panic(err)
	}
	for res.Next(){
		var book Book
		err = res.Scan(&book.IdBook, &book.Name, &book.Price_of_book, &book.Num_of_copies, &book.Authors, &book.Cover_photo, &book.Price_of_book,&book.Reg_date)
		if err!=nil{
			panic(err)
		}
		*AllBooks = append (*AllBooks, book)
	}

}







func Get_books_fronDB(AllBooks *[]Book){
	db := openDB()
	defer db.Close()
	res,err:=db.Query("SELECT `id_book`,`name`,`price_of_book`,`num_of_copies`,`authors`,`cover_photo`,`price_of_book`,`reg_date` FROM `books`")
	if err!=nil{
		panic(err)
	}
	for res.Next(){
		var book Book
		err = res.Scan(&book.IdBook, &book.Name, &book.Price_of_book, &book.Num_of_copies, &book.Authors, &book.Cover_photo, &book.Price_of_book,&book.Reg_date)
		if err!=nil{
			panic(err)
		}
		*AllBooks = append (*AllBooks, book)
	}
}

func Get_priceOfbook_byId(idBook string) string {
	db:=openDB()
	defer db.Close()
	s := fmt.Sprintf("SELECT `price_per_day` FROM `books` WHERE `id_book` = '%s'", idBook)
	res,err:=db.Query(s)
	if err!=nil{
		panic(err)
	}
	pricePerDay:=""
	for res.Next(){
		err = res.Scan(&pricePerDay)
		if err!=nil{
			panic(err)
		}
	}
	fmt.Println("price per day is ", pricePerDay)
	return pricePerDay
}

func Get_numOfCopies_fromDB(idBook string) int {
	db:=openDB()
	defer db.Close()
	s := fmt.Sprintf("SELECT `num_of_copies` FROM `books` WHERE `id_book` = '%s'", idBook)
	res,err:=db.Query(s)
	if err!=nil{
		panic(err)
	}
	numofcopies:=""
	for res.Next(){
		err = res.Scan(&numofcopies)
		if err!=nil{
			panic(err)
		}
	}
	fmt.Println("price per day is ", numofcopies)
	num,_:=strconv.Atoi(numofcopies)
	return num
}
func Set_numOfCopies_fromDB(numOfCopies string,idBook string){
	db := openDB()
	defer db.Close()
	s := fmt.Sprintf("UPDATE `books` SET `num_of_copies` = '%s' WHERE `id_book` = '%s'",numOfCopies,idBook)
	log.Println(s)
	insert, err1 := db.Query(s)
	if err1 != nil {
		panic(err1)
	}
	defer insert.Close()
}



func Get_AllbookID_fromDB(idBooks *[]int) {
	db := openDB()
	defer db.Close()
	res,err:=db.Query("SELECT `id_book` FROM `books`")
	if err!=nil{
		panic(err)
	}
	for res.Next(){
		var bookID int
		err = res.Scan(&bookID)
		*idBooks = append (*idBooks, bookID)
	}
}

func Get_CoverPhoto_byId(idBook string) string {
	db:=openDB()
	defer db.Close()
	res,err:=db.Query( fmt.Sprintf("SELECT `cover_photo` FROM `books` WHERE `id_book` = '%s'", idBook))
	if err!=nil{
		panic(err)
	}
	cover:=""
	for res.Next(){
		err = res.Scan(&cover)
		if err!=nil{
			panic(err)
		}
	}
	return cover
}