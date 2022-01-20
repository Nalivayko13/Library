package dao

import (
	"fmt"
	"log"
)
type Genre_to_Book struct {
	IdGenre string `json:"id_genre"`
	IdBook string `json:"id_book"`
}
type Genre struct {
	IdGenre string `json:"id_genre"`
	Name string `json:"name,omitempty"`
}

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
func Get_Genre_fromDB(IdBook int) []Genre{
	db := openDB()
	defer db.Close()
	res1,err:=db.Query(fmt.Sprintf("SELECT `id_genre` FROM `book_genre` WHERE `id_book`=%d",IdBook))
	var Gen []Genre
	if err!=nil{
		panic(err)
	}
	for res1.Next(){
		var g Genre
		err = res1.Scan(&g.IdGenre)
		if err!=nil{
			panic(err)
		}
		Gen = append(Gen, g)
	}
	return Gen
}

func Get_GenreName_fromDB(idGenre int) []Genre{
	db := openDB()
	defer db.Close()
	res1,err:=db.Query(fmt.Sprintf("SELECT `name` FROM `genres2` WHERE `id_genre`=%d",idGenre))
	var Gen []Genre
	if err!=nil{
		panic(err)
	}
	for res1.Next(){
		var g Genre
		err = res1.Scan(&g.Name)
		if err!=nil{
			panic(err)
		}
		Gen = append(Gen, g)
	}
	return Gen
}
