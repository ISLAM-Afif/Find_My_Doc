package usecases

import (
	"findMyDoc/doctors/repositories"
	"findMyDoc/internal/entities"
)

type DoctorUsecase interface {
	SearchDoctors(speciality string, latitude, longitude float64) ([]entities.Doctor, error)
}

type doctorUsecase struct {
	doctorRepo repositories.DoctorRepository
}

func NewDoctorUsecase(repo repositories.DoctorRepository) DoctorUsecase {
	return &doctorUsecase{
		doctorRepo: repo,
	}
}

func (u *doctorUsecase) SearchDoctors(speciality string, latitude, longitude float64) ([]entities.Doctor, error) {
	return u.doctorRepo.SearchDoctors(speciality, latitude, longitude)
}

func (u *doctorUsecase) GetDoctorById(id int) (entities.Doctor, error) {
	return u.doctorRepo.GetDoctorById(id)
}
