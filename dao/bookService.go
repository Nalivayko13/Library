package dao

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	_ "github.com/gorilla/mux"
	"log"
)

func openDB() *sql.DB {
	db, err :=sql.Open("mysql","root:@tcp(127.0.0.1:3306)/library")
	if err!=nil {
		log.Fatal(err)
	}
	if err = db.Ping(); err != nil {
		return nil
	}
	return db
}
type Book struct {
	IdBook int `json:"id_book"`
	Name string `json:"name"`
	Genre string `json:"genre"`
	Price_of_book string `json:"price_of_book"`
	Num_of_copies string `json:"num_of_copies"`
	Authors string `json:"authors"`
	Cover_photo string `json:"cover_photo"`
	Price_per_day string `json:"price_per_day"`
	Reg_date string `json:"reg_date"`
}
func Save_book_toDB(Mbook Book){
	var book Book
	book=Mbook
	db := openDB()
	defer db.Close()
	s := fmt.Sprintf("INSERT INTO `books` (`name`,`genre`,`price_of_book`,`num_of_copies`,`authors`, `price_per_day`,`reg_date`) VALUES ('%s', '%s', '%s', '%s', '%s', '%s', '%s')", book.Name, book.Genre, book.Price_of_book, book.Num_of_copies,book.Authors,book.Price_per_day,book.Reg_date)
	fmt.Println(s)
	insert, err1 := db.Query(s)
	if err1 != nil {
		panic(err1)
	}

	fmt.Println(insert)
	defer insert.Close()
}

func Get_books_fronDB(AllBooks *[]Book){
	db := openDB()
	defer db.Close()
	res,err:=db.Query("SELECT * FROM `books`")
	if err!=nil{
		panic(err)
	}
	for res.Next(){
		var book Book
		err = res.Scan(&book.IdBook, &book.Name,&book.Genre, &book.Price_of_book, &book.Num_of_copies, &book.Authors, &book.Cover_photo, &book.Price_of_book,&book.Reg_date)
		if err!=nil{
			panic(err)
		}
		*AllBooks = append (*AllBooks, book)
	}
}
