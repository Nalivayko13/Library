package dao

import (
	"fmt"
	"log"
	"strings"
)
type Genre_to_Book struct {
	IdGenre int `json:"id_genre"`
	IdBook int `json:"id_book"`
}
type Genre struct {
	IdBook string `json:"id_book,omitempty"`
	IdGenre int `json:"id_genre"`
	Name string `json:"name,omitempty"`
}
type Genres []Genre

func Save_BookGenre_toDB(IdBook int,genre Genre) {
	db := openDB()
	defer db.Close()
	s := fmt.Sprintf("INSERT INTO `book_genre` (`id_book`, `id_genre`) VALUES (%d, %d);", IdBook, genre.IdGenre)
	fmt.Println(s)
	insert, err1 := db.Query(s)
	if err1 != nil {
		log.Println(err1)
	}
	fmt.Println(insert)
	defer insert.Close()

}

func Save_AllBookGenre_toDB(gen []interface{}) { //можно ли сделать одним запросом?
	db := openDB()
	defer db.Close()
	stmt :=`INSERT INTO book_genre (id_book, id_genre) VALUES (?` + strings.Repeat(",?", len(gen)-1) + `)`
	res1, err := db.Query(stmt,gen...)

	if err != nil {
		log.Println(err)
	}
	fmt.Println(res1)
	defer res1.Close()

}

func Get_Genre_fromDB(IdBook int) []Genre{
	db := openDB()
	defer db.Close()
	res1,err:=db.Query(fmt.Sprintf("SELECT book_genre.id_genre, genres2.name FROM genres2 JOIN book_genre ON genres2.id_genre=book_genre.id_genre AND book_genre.id_book=%d",IdBook))
	var Gen []Genre
	if err!=nil{
		panic(err)
	}
	for res1.Next(){
		var g Genre
		err = res1.Scan(&g.IdGenre, &g.Name)
		if err!=nil{
			panic(err)
		}
		Gen = append(Gen, g)
	}
	return Gen
}

func Get_allGenre_fromDB(IdBook []interface{}) []Genre{
	db := openDB()
	defer db.Close()
	stmt := `SELECT book_genre.id_book, book_genre.id_genre, genres2.name FROM genres2 JOIN book_genre ON genres2.id_genre=book_genre.id_genre AND book_genre.id_book IN (?` + strings.Repeat(",?", len(IdBook)-1) + `)`
	res1, err := db.Query(stmt, IdBook...)
	if err!=nil{
		panic(err)
	}

	var Gen []Genre
	for res1.Next(){
			var g Genre
			err = res1.Scan(&g.IdBook,&g.IdGenre, &g.Name)
			if err != nil {
				panic(err)
			}
			Gen = append(Gen, g)

	}
	fmt.Println("ЭТО МАССИВ ЖАНРОВ",Gen)
	return Gen
}

