package usecases

import (
	"errors"
	"findMyDoc/internal/entities"
	"findMyDoc/users/repositories"
	"findMyDoc/users/requests"
	"findMyDoc/pkg/auth"

	"golang.org/x/crypto/bcrypt"
)

type UserUsecase interface {
	LoginUser(email, password string) (string, error)
	RegisterUser(req requests.RegisterRequest) (string, error)
}

type userUsecase struct {
	repo repositories.UserRepository
}

func NewUserUsecase(repo repositories.UserRepository) UserUsecase {
	return &userUsecase{repo: repo}
}

func (uc *userUsecase) RegisterUser(req requests.RegisterRequest) (string, error) {
	// Hash the password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		return "",errors.New("failed to hash password")
	}

	// Create user entity
	user := entities.User{
		Email:    req.Email,
		Password: string(hashedPassword),
		Role:     req.Role,
	}

	// Save user to the database
	if err := uc.repo.CreateUser(&user); err != nil {
		return "",errors.New("failed to create user")
	}

	// Create doctor or patient profile
	if req.Role == "doctor" {
		doctor := entities.Doctor{
			ID:     int(user.ID),
			Name:       req.Name,
			Speciality: req.Speciality,
			Latitude:   req.Latitude,
			Longitude:  req.Longitude,
			Ratings:    0.0,
		}
		if err := uc.repo.CreateDoctor(&doctor); err != nil {
			return "",errors.New("failed to create doctor profile")
		}
	} else if req.Role == "patient" {
		patient := entities.Patient{
			ID:  int(user.ID),
			Name:    req.Name,
			Ratings: 0.0,
		}
		if err := uc.repo.CreatePatient(&patient); err != nil {
			return "",errors.New("failed to create patient profile")
		}
	} else {
		return "", errors.New("invalid role")
	}

	return auth.GenerateToken(int(user.ID), req.Role)
}

// LoginUser verifies the password and returns a JWT token
func (u *userUsecase) LoginUser(email, password string) (string, error) {
	user, err := u.repo.GetUserByEmail(email)
	if err != nil {
		return "", errors.New("invalid email or password")
	}

	// Verify password
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
		return "", errors.New("invalid email or password")
	}

	// Generate JWT token
	return auth.GenerateToken(int(user.ID), user.Role)
}
