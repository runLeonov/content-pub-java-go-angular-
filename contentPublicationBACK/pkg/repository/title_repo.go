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
	r.db.Debug().
		Raw("SELECT * FROM titles INNER JOIN title_contents tc on titles.id = tc.title_id WHERE released = true ORDER BY tc.likes_count DESC").
		Preload(serialTableE).
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
		Preload("User").
		Preload(commentTableE, func(db *gorm.DB) *gorm.DB {
			return db.Order("creation_date desc")
		}).
		Preload(commentUserTableE).
		Where("released = true").
		First(&title)
	return title, nil
}

func (r *TitleRepo) GetImagesByTitleId(id int) ([]app.Image, error) {
	var title app.Title
	r.db.Debug().Where("id = ?", id).
		Preload(categoryTableE).
		Preload(serialTableE).
		Preload(tagTableE).
		Preload(titleContentTableE).
		Preload(imagesTableE).
		Preload(likeTableE).
		Where("released = true").
		First(&title)
	return title.Content.Images, nil
}

func (r *TitleRepo) LikeById(like app.Like) error {
	r.db.Exec("UPDATE title_contents SET likes_count = likes_count + 1 WHERE title_id = ?", &like.TitleContentID)
	return r.db.Debug().Create(&like).Error
}

func (r *TitleRepo) CommentById(comment app.Comment) error {
	return r.db.Debug().Create(&comment).Error
}

func (r *TitleRepo) GetByRowFilter(row string, ids []int) ([]app.Title, error) {
	var titles []app.Title
	r.db.Debug().
		Raw("SELECT * FROM titles INNER JOIN title_contents tc on titles.id = tc.title_id ORDER BY tc.likes_count DESC").
		Preload(serialTableE).
		Preload(categoryTableE).
		Preload(tagTableE).
		Preload(titleContentTableE).
		Find(&titles)
	return titles, nil
}

func (r *TitleRepo) AddView(id int) error {
	return r.db.Exec("UPDATE title_contents SET views = views + 1 WHERE title_id = ?", id).Error
}

func (r *TitleRepo) UnLikeById(like app.Like) error {
	r.db.Exec("UPDATE title_contents SET likes_count = likes_count - 1 WHERE title_id = ?", &like.TitleContentID)
	return r.db.Where("title_content_id = ?", &like.TitleContentID).Delete(&like).Error
}

func (r *TitleRepo) GetRandom() (app.Title, error) {
	var title app.Title
	r.db.Debug().
		Raw("SELECT * FROM titles WHERE released = true ORDER BY RAND() LIMIT 1 ").
		Preload(categoryTableE).
		Preload(serialTableE).
		Preload(tagTableE).
		Preload(titleContentTableE).
		Preload(imagesTableE).
		Preload(likeTableE).
		Preload("User").
		Preload(commentTableE, func(db *gorm.DB) *gorm.DB {
			return db.Order("creation_date desc")
		}).
		Preload(commentUserTableE).
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
		Where("released = true").
		Where("title_name LIKE %?%", name).Find(&titles)
	return titles, nil
}

func (r *TitleRepo) DeleteTitleById(id int) (uint, error) {
	_ = r.db.Debug().Delete(&app.Title{}, id).Association("Titles.Serials.Categories.Tags").Delete()
	return uint(id), nil
}

func (r *TitleRepo) GetTitlesByCategories(ids []uint) ([]app.Title, error) {
	var titles []app.Title
	r.db.Debug().
		Preload(categoryTableE).
		Preload(serialTableE).
		Preload(tagTableE).
		Preload(titleContentTableE).
		Where("released = true").
		Where("titles.id IN (SELECT title_id FROM titles_categories JOIN categories t on titles_categories.category_id = t.id WHERE t.id IN ?)", ids).
		Find(&titles)
	return titles, nil
}

func (r *TitleRepo) GetTitlesByTags(ids []uint) ([]app.Title, error) {
	var titles []app.Title
	r.db.Debug().
		Preload(categoryTableE).
		Preload(serialTableE).
		Preload(tagTableE).
		Preload(titleContentTableE).
		Where("released = true").
		Where("titles.id IN (SELECT title_id FROM titles_tags JOIN tags t on titles_tags.tag_id = t.id WHERE t.id IN ?)", ids).
		Find(&titles)

	return titles, nil
}

func (r *TitleRepo) GetTitlesBySerials(ids []uint) ([]app.Title, error) {
	var titles []app.Title
	r.db.Debug().
		Preload(categoryTableE).
		Preload(serialTableE).
		Preload(tagTableE).
		Preload(titleContentTableE).
		Where("released = true").
		Where("titles.id IN (SELECT title_id FROM titles_serials JOIN serials t on titles_serials.serial_id = t.id WHERE t.id IN ?)", ids).
		Find(&titles)
	return titles, nil
}
