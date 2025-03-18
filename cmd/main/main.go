package main

import (
	"log"
	"net/http"


	"findMyDoc/pkg/db"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"

	"findMyDoc/middlewares"
)

import doctorsControllers "findMyDoc/doctors/controllers"
import doctorsUsecases "findMyDoc/doctors/usecases"
import doctorsRepositories "findMyDoc/doctors/repositories"

import patientsRepositories "findMyDoc/patients/repositories"

import appointmentsControllers "findMyDoc/appoinments/controllers"
import appointmentsUsecases "findMyDoc/appoinments/usecases"
import appointmentsRepositories "findMyDoc/appoinments/repositories"

import usersControllers "findMyDoc/users/controllers"
import usersUsecases "findMyDoc/users/usecases"
import usersRepositories "findMyDoc/users/repositories"

func main() {
	// Setup database connection
	connStr := "user=postgres password=1696 dbname=find_my_doc sslmode=disable"
	database, err := db.NewPostgresDB(connStr)
	if err != nil {
		log.Fatal(err)
	}

	// Setup doctor-related components
	doctorRepo := doctorsRepositories.NewDoctorRepository(database)
	doctorUsecase := doctorsUsecases.NewDoctorUsecase(doctorRepo)
	doctorController := doctorsControllers.NewDoctorController(doctorUsecase)

	patientRepo := patientsRepositories.NewPatientRepository(database)

	// Setup appointment-related components
	appointmentRepo := appointmentsRepositories.NewAppointmentRepository(database)
	appointmentUsecase := appointmentsUsecases.NewAppointmentUsecase(appointmentRepo)
	appointmentController := appointmentsControllers.NewAppointmentController(appointmentUsecase)

	// User authentication setup
	userRepo := usersRepositories.NewUserRepository(database)
	userUsecase := usersUsecases.NewUserUsecase(userRepo)
	userController := usersControllers.NewUserController(userUsecase , doctorRepo, patientRepo)


	// Setup router
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Post("/register", userController.RegisterHandler)
	r.Post("/login", userController.LoginHandler)

	r.Route("/api", func(r chi.Router) {
		r.Use(middlewares.JWTMiddleware)
		r.Get("/doctors", doctorController.SearchDoctors) // search doctor
		r.Post("/appointments", appointmentController.BookAppointmentHandler) // book an appoinment
		r.Get("/doctors/appointments/pending", appointmentController.GetPendingAppointmentsHandler) // pending appoinment list
		r.Patch("/appointments/{id}/accept", appointmentController.AcceptAppointmentHandler) // accept appoinment
		r.Get("/doctors/appointments/accepted", appointmentController.GetAcceptedAppointmentsHandler) // accepted appoinment list
		r.Get("/profile", userController.GetProfile) // get profile
	})

	// Start server
	log.Println("Server running on port 3001")
	http.ListenAndServe(":3001", r)
}
