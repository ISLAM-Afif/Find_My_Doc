package controllers

import (
	"encoding/json"
	"findMyDoc/appoinments/usecases"
	"findMyDoc/internal/entities"
	"findMyDoc/internal/utils"
	"net/http"

	"strconv"

	"github.com/go-chi/chi/v5"
)

type AppointmentController struct {
	usecase usecases.AppointmentUsecase
}

func NewAppointmentController(uc usecases.AppointmentUsecase) *AppointmentController {
	return &AppointmentController{
		usecase: uc,
	}
}

func (c *AppointmentController) BookAppointmentHandler(w http.ResponseWriter, r *http.Request) {

	authHeader := r.Header.Get("Authorization")
	userID := utils.ExtractUserIDFromToken(authHeader)

	var appointment entities.Appointment
	if err := json.NewDecoder(r.Body).Decode(&appointment); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	appointment.PatientID = userID

	err := c.usecase.BookAppointment(&appointment)
	if err != nil {
		http.Error(w, "Failed to book appointment", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(appointment)
}

func (c *AppointmentController) GetPendingAppointmentsHandler(w http.ResponseWriter, r *http.Request) {
	
	authHeader := r.Header.Get("Authorization")
	doctorID := utils.ExtractUserIDFromToken(authHeader)

	appointments, err := c.usecase.GetPendingAppointments(doctorID)
	if err != nil {
		http.Error(w, "Failed to fetch pending appointments", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(appointments)
}

func (c *AppointmentController) AcceptAppointmentHandler(w http.ResponseWriter, r *http.Request) {
	appointmentIDStr := chi.URLParam(r, "id")
	appointmentID, err := strconv.Atoi(appointmentIDStr)
	if err != nil {
		http.Error(w, "Invalid appointment ID", http.StatusBadRequest)
		return
	}

	err = c.usecase.AcceptAppointment(appointmentID)
	if err != nil {
		http.Error(w, "Failed to accept appointment", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Appointment accepted successfully"))
}

func (c *AppointmentController) GetAcceptedAppointmentsHandler(w http.ResponseWriter, r *http.Request) {
	authHeader := r.Header.Get("Authorization")
	doctorID := utils.ExtractUserIDFromToken(authHeader)

	appointments, err := c.usecase.GetAcceptedAppointments(doctorID)
	if err != nil {
		http.Error(w, "Failed to fetch accepted appointments", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(appointments)
}
