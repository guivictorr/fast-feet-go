package handler

import (
	"github.com/guivictorr/fast-feet-go/config"
	"gorm.io/gorm"
)

var (
	db     *gorm.DB
	logger *config.Logger
)

func InitializeHandler() {
	logger = config.GetLogger("handler")
	db = config.GetSQLite()
}
