package userControllers

import (
	"net/http"

	"github.com/guivictorr/fast-feet-go/models"
	"gorm.io/gorm"
)

type Repository interface {
	CreateUser(*models.UserEntity) (*models.UserEntity, int)
}

type repository struct {
	db *gorm.DB
}

func NewUserController(db *gorm.DB) *repository {
	return &repository{db: db}
}

func (repo *repository) CreateUser(input *models.UserEntity) (*models.UserEntity, int) {
	db := repo.db

	var user models.UserEntity

	// CHECK IF THIS IS CORRECT
	checkIfUserExists := db.Select("*").Where("cpf=?", input.Cpf).Find(&user)

	if checkIfUserExists.RowsAffected > 0 {
		return nil, http.StatusConflict
	}

	createUser := db.Create(&user)

	if createUser.Error != nil {
		return nil, http.StatusExpectationFailed
	}

	return input, http.StatusCreated
}
