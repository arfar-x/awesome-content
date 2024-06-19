package repository

import (
	"gorm.io/gorm"
)

type BaseRepositoryAdapter struct {
	DB *gorm.DB
}
