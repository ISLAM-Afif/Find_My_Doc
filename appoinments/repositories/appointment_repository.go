package repositories

import (
	"findMyDoc/internal/entities"
	"gorm.io/gorm"
)

type AppointmentRepository interface {
	BookAppointment(appointment *entities.Appointment) error
	GetPendingAppointments(doctorID int) ([]entities.Appointment, error)
	UpdateAppointmentStatus(appointmentID int, status string) error
	GetAppointmentsByStatus(doctorID int, status string) ([]entities.Appointment, error)
}

type appointmentRepository struct {
	db *gorm.DB
}

func NewAppointmentRepository(db *gorm.DB) AppointmentRepository {
	return &appointmentRepository{db: db}
}

func (r *appointmentRepository) BookAppointment(appointment *entities.Appointment) error {
	return r.db.Create(appointment).Error
}

func (r *appointmentRepository) GetPendingAppointments(doctorID int) ([]entities.Appointment, error) {
	var appointments []entities.Appointment
	err := r.db.Where("doctor_id = ? AND status = ?", doctorID, "pending").Find(&appointments).Error
	return appointments, err
}

func (r *appointmentRepository) UpdateAppointmentStatus(appointmentID int, status string) error {
	return r.db.Model(&entities.Appointment{}).
		Where("id = ?", appointmentID).
		Update("status", status).Error
}

func (r *appointmentRepository) GetAppointmentsByStatus(doctorID int, status string) ([]entities.Appointment, error) {
	var appointments []entities.Appointment
	err := r.db.
		Preload("Doctor").
		Preload("Patient").
		Where("doctor_id = ? AND status = ?", doctorID, status).
		Find(&appointments).Error
	return appointments, err
}