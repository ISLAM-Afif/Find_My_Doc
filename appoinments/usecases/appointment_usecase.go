package usecases

import (
	"findMyDoc/appoinments/repositories"
	"findMyDoc/internal/entities"
)

type AppointmentUsecase interface {
	BookAppointment(appointment *entities.Appointment) error
	GetPendingAppointments(doctorID int) ([]entities.Appointment, error)
	AcceptAppointment(appointmentID int) error
	GetAcceptedAppointments(doctorID int) ([]entities.Appointment, error)
}

type appointmentUsecase struct {
	appointmentRepo repositories.AppointmentRepository
}

func NewAppointmentUsecase(repo repositories.AppointmentRepository) AppointmentUsecase {
	return &appointmentUsecase{
		appointmentRepo: repo,
	}
}

func (u *appointmentUsecase) BookAppointment(appointment *entities.Appointment) error {
	appointment.Status = "pending" // Default status
	return u.appointmentRepo.BookAppointment(appointment)
}

func (u *appointmentUsecase) GetPendingAppointments(doctorID int) ([]entities.Appointment, error) {
	return u.appointmentRepo.GetPendingAppointments(doctorID)
}

func (u *appointmentUsecase) AcceptAppointment(appointmentID int) error {
	return u.appointmentRepo.UpdateAppointmentStatus(appointmentID, "accepted")
}

func (u *appointmentUsecase) GetAcceptedAppointments(doctorID int) ([]entities.Appointment, error) {
	return u.appointmentRepo.GetAppointmentsByStatus(doctorID, "accepted")
}