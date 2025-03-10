package entities

type Patient struct {
	ID         int     `json:"id" gorm:"primaryKey"`
	Name       string  `json:"name"`
	Ratings    float64 `json:"ratings"`
}
