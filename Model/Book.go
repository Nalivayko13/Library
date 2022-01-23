package Model

import (
	"errors"
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

func SaveBook(book *dao.Book) error{

	_,err1:=strconv.Atoi(book.Price_of_book)
	if err1!=nil{
		log.Print(err1)
		return errors.New("incorrect price of book")
	}
	_,err2:=strconv.Atoi(book.Price_per_day)
	if err2!=nil{
		log.Print(err2)
		return errors.New("incorrect price per day")
	}
	_,err3:=strconv.Atoi(book.Num_of_copies)
	if err3!=nil{
		log.Print(err3)
		return errors.New("incorrect number of copies")
	}
	if book.Name=="" || book.Genre==nil || book.Price_of_book=="" || book.Price_per_day=="" || book.Num_of_copies=="" || book.Reg_date=="" || book.Authors=="" || book.Cover_photo==""{
		log.Println("No data")
		return errors.New("No data")
	}
	err4 := DownloadFile(book.Cover_photo,fmt.Sprintf("C:/GO/booksCover/%s.jpg",book.Name))
	if err4 != nil {
		log.Fatal(err4)
	}
    book.IdBook=dao.Save_book_toDB(*book)
	for _, genre := range book.Genre {
		//genre.BookName=book.Name
		dao.Save_BookGenre_toDB(book.IdBook,genre)
	}
	return nil
}

func Home(AllBooks []dao.Book, limit, page int) []dao.Book {
	AllBooks = []dao.Book{}
	dao.Get_booksWithPage_fronDB(&AllBooks,limit, page)
	if limit==0 || page==0 {
	dao.Get_books_fronDB(&AllBooks)
		log.Println("this is all books")
	}

	args := make([]interface{}, len(AllBooks))
	for i, _ := range AllBooks {
		args[i] = AllBooks[i].IdBook
	}
	fmt.Println("это массив айди", args)
	//запрос в бд для получения жанров по всем айдиникам сразу
	var gen []dao.Genre
	gen=dao.Get_allGenre_fromDB(args)
	for i,_:= range args {
		for j, _ := range gen {
		g, _ := strconv.Atoi(gen[j].IdBook)
			if AllBooks[i].IdBook == g {
				AllBooks[i].Genre=append(AllBooks[i].Genre,gen[j])
				fmt.Println(AllBooks[i].Genre)
			}
		}

	}
	fmt.Println(AllBooks)
	return AllBooks
}
