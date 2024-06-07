package mysql

import (
	"time"

	"gorm.io/gorm"
)

type BaseRepositoryAdapter struct {
	DB *gorm.DB
}

type BaseModel struct {
	ID        uint `gorm:"primaryKey"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}
