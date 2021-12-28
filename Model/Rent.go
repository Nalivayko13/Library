package Model
type Rent struct {
	Id_rent string `json:"id_rent"`
	Id_book string `json:"id_book"`
	Id_reeder string `json:"id_reeder"`
	First_date string `json:"first_date"`
	Last_date string `json:"last_date"`
	Fine float64 `json:"fine"`
}
