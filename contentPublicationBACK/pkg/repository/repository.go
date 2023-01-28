package repository

import (
	app "contentPublicationBACK"
	"gorm.io/gorm"
)

type Authorization interface {
	CreateUser(user app.User) (int, error)
	GetUser(email, password string) (app.User, error)
}

type Titles interface {
	GetAllTitles() ([]app.Title, error)
	GetTitleById(id int) (app.Title, error)
	GetRandom() (app.Title, error)
	SaveNewTitle(title app.Title) (uint, error)
	UpdateTitle(title app.Title, id int) (uint, error)
	DeleteTitleById(id int) (uint, error)
	GetTitlesByCategories(ids []uint) ([]app.Title, error)
	GetAllTitlesByNameRegex(name string) ([]app.Title, error)
}

type Categories interface {
	GetAll() ([]app.Category, error)
	GetByGenres(genres []string) ([]app.Category, error)
}

type Repository struct {
	Authorization
	Titles
	Categories
}

func NewRepository(db *gorm.DB) *Repository {
	return &Repository{
		Authorization: NewAuthRepo(db),
		Titles:        NewTitleRepo(db),
		Categories:    NewCategoryRepo(db),
	}
}
