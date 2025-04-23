package handler

import (
	"github.com/mallet-zn/gopportunities/config"
	"gorm.io/gorm"
)

var (
	logger *config.Logger
	db     *gorm.DB
)

func InitilizeHandler() {
	logger = config.GetLogger("handler")
	db = config.GetSLite()
}
