package repository

import (
	app "contentPublicationBACK"
	"gorm.io/gorm"
)

type CategoryRepo struct {
	db *gorm.DB
}

func NewCategoryRepo(db *gorm.DB) *CategoryRepo {
	return &CategoryRepo{db: db}
}

func (r *CategoryRepo) GetAll() ([]app.Category, error) {
	var categories []app.Category
	r.db.Preload(titleTableE).Find(&categories)
	return categories, nil
}

func (r *CategoryRepo) GetByGenres(genres []string) ([]app.Category, error) {
	var categories []app.Category
	r.db.Preload(titleTableE).
		Joins("inner join title_categories tc on tc.category_id = categories.id ").
		Where("genre IN ?", genres).Find(&categories)
	return categories, nil
}
