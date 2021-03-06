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
	if book.IdBook==-1{
		return errors.New("Something wrong")
	}
	for i, _ := range book.Genre {
		dao.Save_BookGenre_toDB(book.IdBook,book.Genre[i])
	}

	return nil
}

func Home(AllBooks []dao.Book, limit, page int) ([]dao.Book,error) {
	AllBooks = []dao.Book{}
	var Total = []dao.Book{}
	dao.Get_booksWithPage_fronDB(&AllBooks,limit, page)
	dao.Get_books_fronDB(&Total)
	totalCount:=len(Total)
	LimitOfPages:=(totalCount/limit)+1

	if LimitOfPages<page{
		return nil,errors.New("No more pages")
	}
	if limit==0 || page==0 {
	//dao.Get_books_fronDB(&AllBooks)
		//log.Println("this is all books")
		return nil,errors.New("no page or limit")
	}

	args := make([]interface{}, len(AllBooks))
	for i, _ := range AllBooks {
		args[i] = AllBooks[i].IdBook
	}
	fmt.Println("?????? ???????????? ????????", args)
	//???????????? ?? ???? ?????? ?????????????????? ???????????? ???? ???????? ?????????????????? ??????????
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
	return AllBooks, nil
}
