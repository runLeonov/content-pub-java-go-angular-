package repository

import (
	app "contentPublicationBACK"
	"gorm.io/gorm"
	"time"
)

type TitleRepo struct {
	db *gorm.DB
}

func NewTitleRepo(db *gorm.DB) *TitleRepo {
	return &TitleRepo{db: db}
}

func (r *TitleRepo) GetAllTitles() ([]app.Title, error) {
	var titles []app.Title
	r.db.Preload(serialTableE).
		Preload(categoryTableE).
		Preload(tagTableE).
		Preload(titleContentTableE).
		Preload(imagesTableE).
		Find(&titles)
	return titles, nil
}

func (r *TitleRepo) GetTitleById(id int) (app.Title, error) {
	var title app.Title
	r.db.Debug().Where("id = ?", id).
		Preload(categoryTableE).
		Preload(serialTableE).
		Preload(tagTableE).
		Preload(titleContentTableE).
		Preload(imagesTableE).
		First(&title)
	return title, nil
}

func (r *TitleRepo) SaveNewTitle(title app.Title) (uint, error) {
	title.CreationDate = time.Now()
	r.db.Debug().Create(&title)
	return title.ID, nil
}

func (r *TitleRepo) UpdateTitle(title app.Title, id int) (uint, error) {
	title.ID = uint(id)
	r.db.Where("id = ?", id).Save(&title)
	return title.ID, nil
}

func (r *TitleRepo) GetAllTitlesByNameRegex(name string) ([]app.Title, error) {
	var titles []app.Title
	r.db.Preload(serialTableE).
		Preload(categoryTableE).
		Preload(tagTableE).
		Preload(titleContentTableE).
		Preload(imagesTableE).
		Where("title_name LIKE %?%", name).Find(&titles)
	return titles, nil
}

func (r *TitleRepo) DeleteTitleById(id int) (uint, error) {
	//err := r.db.Debug().Delete(&app.Title{}, id).Association(categoryTableE).Delete()
	_ = r.db.Debug().Delete(&app.Title{}, id).Association("Titles.Serials.Categories.Tags").Delete()
	//if err != nil {
	//	return 0, err
	//}
	return uint(id), nil
}

func (r *TitleRepo) GetTitlesByCategories(categoriesIds []uint) ([]app.Title, error) {
	var titles []app.Title
	r.db.Preload(categoryTableE).
		Preload(serialTableE).
		Preload(tagTableE).
		Preload(titleContentTableE).
		Preload(imagesTableE).
		Joins("inner join titles_categories tc on tc.category_id = titles.id ").
		Where("id IN ?", categoriesIds).Find(&titles)
	return titles, nil
}
