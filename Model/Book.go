package Model

import (
	_ "github.com/go-sql-driver/mysql"
	_ "github.com/gorilla/mux"
)

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
