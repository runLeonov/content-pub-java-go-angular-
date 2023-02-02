package repository

import (
	app "contentPublicationBACK"
	"gorm.io/gorm"
)

type AuthRepo struct {
	db *gorm.DB
}

func NewAuthRepo(db *gorm.DB) *AuthRepo {
	return &AuthRepo{db: db}
}

func (r *AuthRepo) CreateUser(user app.User) (int, error) {
	r.db.Debug().Create(&user)
	return user.ID, nil
}

func (r *AuthRepo) GetUser(email, password string) (app.User, error) {
	var user app.User
	err := r.db.Where("email = ? AND password = ?", email, password).First(&user).Error
	return user, err
}
