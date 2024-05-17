package userControllers

import (
	"github.com/guivictorr/fast-feet-go/models"
)

type Service interface {
	CreateUser(*UserInput) (*models.UserEntity, int)
	ListUsers() ([]models.UserEntity, int)
	FindUser(userId string) (*models.UserEntity, int)
	DeleteUser(userId string) int
}

type service struct {
	repository Repository
}

func NewUserService(repo *repository) *service {
	return &service{repository: repo}
}

func (s *service) CreateUser(input *UserInput) (*models.UserEntity, int) {
	user := models.UserEntity{
		Name:     input.Name,
		Cpf:      input.Cpf,
		Password: input.Password,
		Role:     input.Role,
	}
	return s.repository.CreateUser(&user)
}

func (s *service) ListUsers() ([]models.UserEntity, int) {
	return s.repository.ListUsers()
}

func (s *service) FindUser(userId string) (*models.UserEntity, int) {
	return s.repository.FindUser(userId)
}

func (s *service) DeleteUser(userId string) int {
	return s.repository.DeleteUser(userId)
}
