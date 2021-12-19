package main

import (
	_ "database/sql"
	_ "fmt"
	_ "github.com/go-sql-driver/mysql"
	_ "html/template"
	"net/http"
	_ "net/http"
)
type User struct {
	Id int `json:"id"`
	Name string `json:"name"`
	Surname string `json:"surname"`
	DateOfBirth string `json:"date_of_birth"`
	Address string `json:"address"`
	Email string `json:"email"`
}

func AddReeder(w http.ResponseWriter,r *http.Request){

}

func GetReeders(w http.ResponseWriter,r *http.Request){

}

