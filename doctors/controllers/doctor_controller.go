package controllers

import (
	"encoding/json"
	"findMyDoc/doctors/usecases"
	"net/http"
	"strconv"
)

type DoctorController struct {
	usecase usecases.DoctorUsecase
}

func NewDoctorController(uc usecases.DoctorUsecase) *DoctorController {
	return &DoctorController{
		usecase: uc,
	}
}

func (c *DoctorController) SearchDoctors(w http.ResponseWriter, r *http.Request) {
	speciality := r.URL.Query().Get("speciality")
	latStr := r.URL.Query().Get("latitude")
	longStr := r.URL.Query().Get("longitude")

	lat, err1 := strconv.ParseFloat(latStr, 64)
	long, err2 := strconv.ParseFloat(longStr, 64)
	if err1 != nil || err2 != nil {
		http.Error(w, "Invalid latitude or longitude", http.StatusBadRequest)
		return
	}

	doctors, err := c.usecase.SearchDoctors(speciality, lat, long)
	if err != nil {
		http.Error(w, "Error retrieving doctors", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(doctors)
}
