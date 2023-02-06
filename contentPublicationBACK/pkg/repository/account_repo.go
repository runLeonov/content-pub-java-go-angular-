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

func (r *AuthRepo) GetUserLikes(id int) ([]app.Title, error) {
	var titles []app.Title
	r.db.Raw("select * from titles "+
		"inner join title_contents on title_id = titles.id "+
		"inner join likes on likes.title_content_id = title_contents.id "+
		"inner join users on users.id = likes.user_id "+
		"where users.id = ?", id).Preload(titleContentTableE).Preload(likeTableE).Scan(&titles)
	return titles, nil
}
