package controllers

import (
	"encoding/json"
	"findMyDoc/users/usecases"
	"findMyDoc/users/requests"
	"net/http"
	"findMyDoc/internal/utils"
)

import doctorsRepositories "findMyDoc/doctors/repositories"
import patientsRepositories "findMyDoc/patients/repositories"

type UserController struct {
	usecase usecases.UserUsecase
	doctorsRepositories doctorsRepositories.DoctorRepository
	patientsRepositories patientsRepositories.PatientRepository
}

func NewUserController(uc usecases.UserUsecase ,dc doctorsRepositories.DoctorRepository, pc patientsRepositories.PatientRepository) *UserController {
	return &UserController{usecase: uc , doctorsRepositories: dc , patientsRepositories: pc}
}

func (c *UserController) LoginHandler(w http.ResponseWriter, r *http.Request) {
	var req struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request", http.StatusBadRequest)
		return
	}

	token, err := c.usecase.LoginUser(req.Email, req.Password)
	if err != nil {
		http.Error(w, "Invalid credentials", http.StatusUnauthorized)
		return
	}

	json.NewEncoder(w).Encode(map[string]string{"token": token})
}

func (c *UserController) RegisterHandler(w http.ResponseWriter, r *http.Request) {
	var req requests.RegisterRequest

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request", http.StatusBadRequest)
		return
	}

	token, err := c.usecase.RegisterUser(req);
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]string{
		"message": "User registered successfully",
		"token": token,
	})
}

func (c *UserController) GetProfile(w http.ResponseWriter, r *http.Request) {
	authHeader := r.Header.Get("Authorization")
	userID := utils.ExtractUserIDFromToken(authHeader)
	role := utils.ExtractRoleFromToken(authHeader)

	if role == "doctor" {
		doctors, err := c.doctorsRepositories.GetDoctorById(userID)
		if err != nil {
			http.Error(w, "Error retrieving doctors", http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(doctors)
	} else {
		patients, err := c.patientsRepositories.GetPatientById(userID)
		if err != nil {
			http.Error(w, "Error retrieving doctors", http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(patients)
	}
	
}
