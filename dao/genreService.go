package dao

import (
	"fmt"
	"log"
)

type Genre struct {
	BookName string `json:"book_name"`
	GenreOfBook string `json:"genre_of_book"`
}

func Save_BookGenre_toDB(genre Genre) {
	db := openDB()
	defer db.Close()
	s := fmt.Sprintf("INSERT INTO `genres` (`book_name`, `genre_of_book`) VALUES ('%s','%s');", genre.BookName, genre.GenreOfBook)
	fmt.Println(s)
	insert, err1 := db.Query(s)
	if err1 != nil {
		log.Println(err1)
	}
	fmt.Println(insert)
	defer insert.Close()

}
