package repositories

import (
	"findMyDoc/internal/entities"

	"gorm.io/gorm"
)

type UserRepository interface {
	CreateUser(user *entities.User) error
	GetUserByEmail(email string) (*entities.User, error)
	CreateDoctor(doctor *entities.Doctor) error
	CreatePatient(patient *entities.Patient) error
	GetUserRoleByUserId(id int) string
}

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &userRepository{db: db}
}

func (r *userRepository) CreateUser(user *entities.User) error {
	return r.db.Create(user).Error
}

func (r *userRepository) GetUserByEmail(email string) (*entities.User, error) {
	var user entities.User
	err := r.db.Where("email = ?", email).First(&user).Error
	return &user, err
}

func (r *userRepository) GetUserRoleByUserId(id int) string {
	var user entities.User
	r.db.Select("role").Where("id = ?", id).First(&user)
	return user.Role
}

func (r *userRepository) CreateDoctor(doctor *entities.Doctor) error {
	return r.db.Create(doctor).Error
}

func (r *userRepository) CreatePatient(patient *entities.Patient) error {
	return r.db.Create(patient).Error
}
