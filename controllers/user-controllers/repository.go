package userControllers

import (
	"net/http"

	"github.com/guivictorr/fast-feet-go/models"
	"gorm.io/gorm"
)

type Repository interface {
	CreateUser(*models.UserEntity) (*models.UserEntity, int)
	ListUsers() ([]models.UserEntity, int)
	FindUser(userId string) (*models.UserEntity, int)
	DeleteUser(userId string) int
	UpdateUser(*models.UserEntity) (*models.UserEntity, int)
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

	if rowsAffected := db.Select("*").Where("cpf=?", input.Cpf).Find(&user).RowsAffected; rowsAffected > 0 {
		return nil, http.StatusConflict
	}

	if err := db.Create(&input).Error; err != nil {
		return nil, http.StatusExpectationFailed
	}

	return input, http.StatusCreated
}

func (repo *repository) ListUsers() ([]models.UserEntity, int) {
	db := repo.db

	var users []models.UserEntity

	if err := db.Find(&users).Error; err != nil {
		return nil, http.StatusNotFound
	}

	return users, http.StatusOK
}

func (repo *repository) FindUser(userId string) (*models.UserEntity, int) {
	db := repo.db

	var user *models.UserEntity

	if rowsAffected := db.Where("id=?", userId).Find(&user).RowsAffected; rowsAffected <= 0 {
		return nil, http.StatusNotFound
	}

	return user, http.StatusOK
}

func (repo *repository) DeleteUser(userId string) int {
	db := repo.db

	var user *models.UserEntity

	if rowsAffected := db.Where("id=?", userId).Find(&user).RowsAffected; rowsAffected <= 0 {
		return http.StatusNotFound
	}

	if err := db.Delete(&user).Error; err != nil {
		return http.StatusInternalServerError
	}

	return http.StatusNoContent
}

func (repo *repository) UpdateUser(input *models.UserEntity) (*models.UserEntity, int) {
	db := repo.db

	var user models.UserEntity

	if rowsAffected := db.Select("*").Where("id=?", input.ID).Find(&user).RowsAffected; rowsAffected <= 0 {
		return nil, http.StatusNotFound
	}

	if err := db.Model(&user).Updates(&input).Error; err != nil {
		return nil, http.StatusExpectationFailed
	}

	return &user, http.StatusOK
}
