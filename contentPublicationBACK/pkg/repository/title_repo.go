package repository

import (
	app "contentPublicationBACK"
	"gorm.io/gorm"
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
		Find(&titles)
	return titles, nil
}

func (r *TitleRepo) GetAllPossibleContent() (app.AllContent, error) {
	var cats []app.Category
	r.db.Find(&cats)
	var tags []app.Tag
	r.db.Find(&tags)
	var serials []app.Serial
	r.db.Find(&serials)
	var types []app.StaticType
	r.db.Find(&types)

	cot := app.AllContent{Categories: cats, Tags: tags, Serials: serials, Types: types}
	return cot, nil
}

func (r *TitleRepo) GetTitleById(id int) (app.Title, error) {
	var title app.Title
	r.db.Debug().Where("id = ?", id).
		Preload(categoryTableE).
		Preload(serialTableE).
		Preload(tagTableE).
		Preload(titleContentTableE).
		Preload(imagesTableE).
		Preload(likeTableE).
		First(&title)
	return title, nil
}

func (r *TitleRepo) LikeById(like app.Like) error {
	r.db.Exec("UPDATE title_contents SET likes_count = likes_count + 1 WHERE title_id = ?", &like.TitleContentID)
	return r.db.Debug().Create(&like).Error
}

func (r *TitleRepo) UnLikeById(like app.Like) error {
	r.db.Exec("UPDATE title_contents SET likes_count = likes_count - 1 WHERE title_id = ?", &like.TitleContentID)
	return r.db.Where("title_content_id = ?", &like.TitleContentID).Delete(&like).Error
}

func (r *TitleRepo) GetRandom() (app.Title, error) {
	var title app.Title
	r.db.Debug().
		Raw("SELECT * FROM titles ORDER BY RAND() LIMIT 1").
		Preload(categoryTableE).
		Preload(serialTableE).
		Preload(tagTableE).
		Preload(titleContentTableE).
		Preload(imagesTableE).
		Preload(likeTableE).
		First(&title)
	return title, nil
}

func (r *TitleRepo) SaveNewTitle(title app.Title) (uint, error) {
	r.db.Create(&title)
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
		Preload(likeTableE).
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
		Preload(likeTableE).
		Joins("inner join titles_categories tc on tc.category_id = titles.id ").
		Where("id IN ?", categoriesIds).Find(&titles)
	return titles, nil
}
