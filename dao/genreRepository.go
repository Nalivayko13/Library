package dao

import (
	"fmt"
	"log"
)

type Genre struct {
	BookName string `json:"book_name"`
	GenreOfBook int `json:"genre_of_book"`
}

func Save_BookGenre_toDB(IdBook int,genre Genre) {
	db := openDB()
	defer db.Close()
	s := fmt.Sprintf("INSERT INTO `book_genre` (`id_book`, `id_genre`) VALUES (%d, %d);", IdBook, genre.GenreOfBook)
	fmt.Println(s)
	insert, err1 := db.Query(s)
	if err1 != nil {
		log.Println(err1)
	}
	fmt.Println(insert)
	defer insert.Close()

}
