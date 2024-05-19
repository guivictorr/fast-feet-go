package authControllers

import (
	"net/http"

	"github.com/guivictorr/fast-feet-go/models"
	"github.com/guivictorr/fast-feet-go/utils"
	"gorm.io/gorm"
)

type Repository interface {
	CreateSession(input *models.UserEntity) (*models.UserEntity, int)
}

type repository struct {
	db *gorm.DB
}

func NewAuthRepository(db *gorm.DB) *repository {
	return &repository{db: db}
}

func (r *repository) CreateSession(input *models.UserEntity) (*models.UserEntity, int) {
	var users models.UserEntity
	db := r.db.Model(&users)

	if rowsAffected := db.Where("cpf=?", input.Cpf).Find(&users).RowsAffected; rowsAffected == 0 {
		return nil, http.StatusNotFound
	}

	err := utils.ComparePassword(users.Password, input.Password)
	if err != nil {
		return nil, http.StatusUnauthorized
	}

	return &users, http.StatusAccepted
}
