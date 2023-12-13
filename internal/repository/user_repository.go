package repository

import (
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
	"main/internal/entity"
)

type UserRepository struct {
	Repository[entity.User]
	Log *logrus.Logger
}

func (r *UserRepository) FindByToken(db *gorm, user *entity.User, token string) error {
	return db.Where("token = ?", token).First(user).Error
}
