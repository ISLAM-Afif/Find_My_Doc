package repositories

import (
	"findMyDoc/internal/entities"

	"gorm.io/gorm"
)

type PatientRepository interface {
	GetPatientById(id int) (entities.Patient, error)
}

type patientRepository struct {
	db *gorm.DB
}

func NewPatientRepository(db *gorm.DB) PatientRepository {
	return &patientRepository{
		db: db,
	}
}

func (r *patientRepository) GetPatientById(id int) (entities.Patient, error) {
	var patient entities.Patient
	if err := r.db.Where("id = ?", id).Find(&patient).Error; err != nil {
		return patient, err
	}
	return patient, nil
}
