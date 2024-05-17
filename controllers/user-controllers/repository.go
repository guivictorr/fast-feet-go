package userControllers

import (
	"net/http"

	"github.com/guivictorr/fast-feet-go/models"
	"gorm.io/gorm"
)

type Repository interface {
	CreateUser(*models.UserEntity) (*models.UserEntity, int)
	ListUsers() ([]models.UserEntity, int)
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

	checkIfUserExists := db.Select("*").Where("cpf=?", input.Cpf).Find(&user)

	if checkIfUserExists.RowsAffected > 0 {
		return nil, http.StatusConflict
	}

	createUser := db.Create(&input)

	if createUser.Error != nil {
		return nil, http.StatusExpectationFailed
	}

	return input, http.StatusCreated
}

func (repo *repository) ListUsers() ([]models.UserEntity, int) {
	db := repo.db

	var users []models.UserEntity

	checkIfUsersExists := db.Find(&users)

	if checkIfUsersExists.Error != nil {
		return nil, http.StatusNotFound
	}

	return users, http.StatusOK
}
