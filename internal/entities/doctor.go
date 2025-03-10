package entities

type Doctor struct {
	ID         int     `json:"id" gorm:"primaryKey"`
	Name       string  `json:"name"`
	Speciality string  `json:"speciality"`
	Latitude   float64 `json:"latitude"`
	Longitude  float64 `json:"longitude"`
	Ratings    float64 `json:"ratings"`
}
