package services

import (
	"golang-user-api/models"
	"golang-user-api/repositories"
)

type UserService struct {
	Repo *repositories.UserRepository
}

func NewUserService(
	repo *repositories.UserRepository,
) *UserService {
	return &UserService{Repo: repo}
}

func (s *UserService) GetAll() ([]models.User, error) {
	return s.Repo.FindAll()
}

func (s *UserService) GetByID(id uint) (models.User, error) {
	return s.Repo.FindByID(id)
}

func (s *UserService) Create(
	req models.CreateUserRequest,
) (models.User, error) {

	user := models.User{
		Name:  req.Name,
		Email: req.Email,
		Age:   req.Age,
	}

	err := s.Repo.Create(&user)

	return user, err
}

func (s *UserService) Update(
	id uint,
	req models.UpdateUserRequest,
) (models.User, error) {

	user, err := s.Repo.FindByID(id)

	if err != nil {
		return user, err
	}

	user.Name = req.Name
	user.Email = req.Email
	user.Age = req.Age

	err = s.Repo.Update(&user)

	return user, err
}

func (s *UserService) Delete(id uint) error {
	user, err := s.Repo.FindByID(id)

	if err != nil {
		return err
	}

	return s.Repo.Delete(&user)
}
