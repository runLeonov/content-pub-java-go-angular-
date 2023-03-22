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
	err := r.db.Debug().Where("id = ?", id).
		Preload("Likes", func(db *gorm.DB) *gorm.DB {
			return db.Limit(1)
		}).
		Preload("Likes.TitleContent").
		Preload("Likes.TitleContent.Title").
		First(&user).Error
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
		Preload(userTableE).
		Joins("inner join title_contents t on t.title_id = titles.id").
		Joins("inner join likes l on l.title_content_id = t.id").
		Joins("inner join users u on u.id = l.user_id ").
		Where("u.id = ?", userId).
		Find(&titles)
	return titles, nil
}
