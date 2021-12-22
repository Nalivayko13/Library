package Model

import (
	_ "database/sql"
	_ "fmt"
	_ "github.com/go-sql-driver/mysql"
	_ "net/http"
)
type Reeder struct {
	IdReeder int `json:"id_reeder"`
	Name string `json:"name"`
	Surname string `json:"surname"`
	DateOfBirth string `json:"date_of_birth"`
	Address string `json:"address"`
	Email string `json:"email"`
}


