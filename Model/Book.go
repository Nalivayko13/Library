package Model

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	_ "github.com/gorilla/mux"
	"library/dao"
	"log"
	"strconv"
)

type Book struct {
	IdBook int `json:"id_book"`
	Name string `json:"name"`
	Genre []Genre `json:"genre"`
	Price_of_book string `json:"price_of_book"`
	Num_of_copies string `json:"num_of_copies"`
	Authors string `json:"authors"`
	Cover_photo string `json:"cover_photo"`
	Price_per_day string `json:"price_per_day"`
	Reg_date string `json:"reg_date"`
}

func SaveBook(book *dao.Book){

	_,err1:=strconv.Atoi(book.Price_of_book)
	if err1!=nil{
		log.Print(err1)
	}
	_,err2:=strconv.Atoi(book.Price_per_day)
	if err2!=nil{
		log.Print(err2)
	}
	_,err3:=strconv.Atoi(book.Num_of_copies)
	if err3!=nil{
		log.Print(err3)
	}

	if book.Name=="" || book.Price_of_book=="" || book.Price_per_day=="" || book.Num_of_copies=="" || book.Reg_date=="" || book.Authors=="" || book.Cover_photo==""{
		log.Println("No data")
	}
	err4 := DownloadFile(book.Cover_photo,fmt.Sprintf("C:/GO/booksCover/%s.jpg",book.Name))
	if err4 != nil {
		log.Fatal(err4)
	}
	dao.Save_book_toDB(*book)
	for _, genre := range book.Genre {
		genre.BookName=book.Name
		dao.Save_BookGenre_toDB(genre)
	}
}

func Home(AllBooks []dao.Book) []dao.Book {
	AllBooks = []dao.Book{}
	dao.Get_books_fronDB(&AllBooks)
	return AllBooks
}