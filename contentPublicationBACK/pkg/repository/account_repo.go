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
		Preload("Comments", "user_id = ?", id).
		Preload("Likes.TitleContent").
		Preload("Likes.TitleContent.Title").
		First(&user).Error
	return user, err
}

func (r *AccountRepo) SaveUserInfo(user app.User) error {
	return r.db.Debug().Save(&user).Error
}

func (r *AccountRepo) GetUserLikes(userId int) ([]app.Title, error) {
	var titles []app.Title
	r.db.Debug().
		Distinct().
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

func (r *AccountRepo) GetUserLikesLimit(userId int) ([]app.Title, error) {
	var title []app.Title
	r.db.Debug().
		Distinct().
		Preload(serialTableE).
		Preload(categoryTableE).
		Preload(tagTableE).
		Preload(imagesTableE).
		Preload(titleContentTableE).
		Preload(userTableE).
		Joins("inner join title_contents t on t.title_id = titles.id").
		Joins("inner join likes l on l.title_content_id = t.id").
		Joins("inner join users u on u.id = l.user_id ").
		Where("u.id = ?", userId).
		First(&title)
	return title, nil
}

func (r *AccountRepo) GetUserComments(userId int) ([]app.Title, error) {
	var titles []app.Title
	r.db.Debug().
		Distinct().
		Preload(serialTableE).
		Preload(categoryTableE).
		Preload(tagTableE).
		Preload(imagesTableE).
		Preload(titleContentTableE).
		Preload(commentTableE, "user_id = ?", userId, func(db *gorm.DB) *gorm.DB {
			return db.Limit(1).Order("creation_date desc")
		}).
		Preload(commentUserTableE, "id = ?", userId).
		Joins("inner join title_contents t on t.title_id = titles.id").
		Joins("inner join comments l on l.title_content_id = t.id").
		Joins("inner join users u on u.id = l.user_id").
		Where("l.user_id = ?", userId).
		Find(&titles)
	return titles, nil
}

func (r *AccountRepo) GetUserCommentsLimit(userId int) ([]app.Title, error) {
	var title []app.Title
	r.db.Debug().
		Distinct().
		Preload(serialTableE).
		Preload(categoryTableE).
		Preload(tagTableE).
		Preload(imagesTableE).
		Preload(titleContentTableE).
		Preload(commentTableE, "user_id = ?", userId, func(db *gorm.DB) *gorm.DB {
			return db.Limit(1).Order("creation_date desc")
		}).
		Preload(commentUserTableE, "id = ?", userId, func(db *gorm.DB) *gorm.DB {
			return db.Limit(1)
		}).
		Joins("inner join title_contents t on t.title_id = titles.id").
		Joins("inner join comments l on l.title_content_id = t.id").
		Joins("inner join users u on u.id = l.user_id").
		Where("l.user_id = ?", userId).
		First(&title)
	return title, nil
}
