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
	err := r.db.Create(&user).Error
	return user.ID, err
}

func (r *AuthRepo) GetUser(email, password string) (app.User, error) {
	var user app.User
	err := r.db.Debug().Where("email = ? AND password = ?", email, password).First(&user).Error
	return user, err
}

func (r *AuthRepo) CheckExist(email string) bool {
	var user app.User
	r.db.Where("email = ?", email).First(&user)
	return user.ID != 0
}

func (r *AuthRepo) GetUserById(id int) (app.User, error) {
	var user app.User
	err := r.db.Where("id = ?", id).First(&user).Error
	return user, err
}
