package Model

type Genre struct {
	IdGenre string `json:"id_genre"`
	Name string `json:"name"`
}

type Genre_to_Book struct {
IdGenre string `json:"id_genre"`
IdBook string `json:"id_book"`
}
//BookName string `json:"book_name"`
//GenreOfBook string `json:"genre_of_book"`