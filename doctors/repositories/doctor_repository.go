package repositories

import (
	"findMyDoc/internal/entities"
	"findMyDoc/internal/utils"

	"gorm.io/gorm"
)

type DoctorRepository interface {
	SearchDoctors(speciality string, latitude, longitude float64) ([]entities.Doctor, error)
	GetDoctorById(id int) (entities.Doctor, error)
}

type doctorRepository struct {
	db *gorm.DB
}

func NewDoctorRepository(db *gorm.DB) DoctorRepository {
	return &doctorRepository{
		db: db,
	}
}

func (r *doctorRepository) SearchDoctors(speciality string, latitude, longitude float64) ([]entities.Doctor, error) {
	var doctors []entities.Doctor

	if speciality == "" {
		if err := r.db.Find(&doctors).Error; err != nil {
			return nil, err
		}
	} else {
		if err := r.db.Where("speciality = ?", speciality).Find(&doctors).Error; err != nil {
			return nil, err
		}
	}

	// Filter doctors by location (within 10km)
	var filteredDoctors []entities.Doctor
	for _, doc := range doctors {
		distance := utils.CalculateDistance(latitude, longitude, doc.Latitude, doc.Longitude)
		if distance <= 10.0 {
			filteredDoctors = append(filteredDoctors, doc)
		}
	}

	return filteredDoctors, nil
}

func (r *doctorRepository) GetDoctorById(id int) (entities.Doctor, error) {
	var doctor entities.Doctor
	if err := r.db.Where("id = ?", id).Find(&doctor).Error; err != nil {
		return doctor, err
	}
	return doctor, nil
}
