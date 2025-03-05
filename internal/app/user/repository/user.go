package repository

import (
	"mycareerapp/internal/domain/dto"
	"mycareerapp/internal/domain/entity"

	"gorm.io/gorm"
)

type UserMySQLItf interface {
    Create(user *entity.User) error
    Get(user *entity.User, userParam dto.UserParam) error
}

type UserMySQL struct {
    db *gorm.DB
}

func NewUserMySQL(db *gorm.DB) UserMySQLItf {
    return &UserMySQL{db}
}

func (r *UserMySQL) Create(user *entity.User) error {
    return r.db.Create(user).Error
}

func (r *UserMySQL) Get(user *entity.User, userParam dto.UserParam) error {
    return r.db.First(&user, userParam).Error
}
