package entities

import "time"

type Appointment struct {
	ID           int       `json:"id" gorm:"primaryKey"`
	DoctorID     int       `json:"doctor_id"`
	Doctor       Doctor    `json:"doctor" gorm:"foreignKey:DoctorID;references:ID"`
	PatientID    int       `json:"patient_id"`
	Patient      Patient   `json:"patient" gorm:"foreignKey:PatientID;references:ID"`
	Status       string    `json:"status"`        // e.g., "pending", "accepted"
	UrgencyLevel int       `json:"urgency_level"` // 1-5 scale
	Latitude     float64   `json:"latitude"`
	Longitude    float64   `json:"longitude"`
	CreatedAt    time.Time `json:"created_at" gorm:"autoCreateTime"`
}
