package authControllers

import (
	"github.com/guivictorr/fast-feet-go/models"
)

type Service interface {
	CreateSession(input *models.UserEntity) (*models.UserEntity, int)
}

type service struct {
	repository Repository
}

func NewAuthService(repository *repository) *service {
	return &service{repository: repository}
}

func (r *service) CreateSession(input *SessionInput) (*models.UserEntity, int) {
	user := models.UserEntity{
		Cpf:      input.Cpf,
		Password: input.Password,
	}
	return r.repository.CreateSession(&user)
}
