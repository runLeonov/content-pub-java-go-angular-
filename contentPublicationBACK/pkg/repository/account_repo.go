package repository

import (
	app "contentPublicationBACK"
	"gorm.io/gorm"
)

type AccountRepo struct {
	db *gorm.DB
}

func NewAccountRepo(db *gorm.DB) *AccountRepo {
	return &AccountRepo{db: db}
}

func (r *AuthRepo) GetUserInfo(id int) (app.User, error) {
	var user app.User
	err := r.db.Where("id = ?", id).Preload("Likes").First(&user).Error
	return user, err
}
