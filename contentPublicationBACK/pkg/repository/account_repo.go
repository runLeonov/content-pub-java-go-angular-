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

func (r *AccountRepo) GetUserInfo(id int) (app.User, error) {
	var user app.User
	err := r.db.Where("id = ?", id).Preload("Likes").First(&user).Error
	return user, err
}

func (r *AccountRepo) GetUserLikes(userId int) ([]app.Title, error) {
	var titles []app.Title
	r.db.Debug().
		Preload(serialTableE).
		Preload(categoryTableE).
		Preload(tagTableE).
		Preload(imagesTableE).
		Preload(titleContentTableE).
		Preload(likeTableE).
		Preload("Content.Likes.User").
		Find(&titles, app.Title{Content: app.TitleContent{Likes: []app.Like{{UserID: uint(userId)}}}})
	return titles, nil
}
