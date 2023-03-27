package repository

import (
	app "contentPublicationBACK"
	"gorm.io/gorm"
)

type Authorization interface {
	CreateUser(user app.User) (int, error)
	GetUser(email, password string) (app.User, error)
	CheckExist(email string) bool
	GetUserById(id int) (app.User, error)
}

type Account interface {
	GetUserInfo(id int) (app.User, error)
	GetUserLikes(id int) ([]app.Title, error)
	SaveUserInfo(user app.User) error
	GetUserLikesLimit(id int) ([]app.Title, error)
	GetUserComments(userId int) ([]app.Title, error)
	GetUserCommentsLimit(userId int) ([]app.Title, error)
}

type Titles interface {
	GetAllTitles() ([]app.Title, error)
	GetAllPossibleContent() (app.AllContent, error)
	GetTitleById(id int) (app.Title, error)
	LikeById(likeObj app.Like) error
	CommentById(comm app.Comment) error
	UnLikeById(likeObj app.Like) error
	GetRandom() (app.Title, error)
	SaveNewTitle(title app.Title) (uint, error)
	UpdateTitle(title app.Title, id int) (uint, error)
	DeleteTitleById(id int) (uint, error)
	GetTitlesByCategories(ids []uint) ([]app.Title, error)
	GetTitlesBySerials(ids []uint) ([]app.Title, error)
	GetTitlesByTags(ids []uint) ([]app.Title, error)
	GetAllTitlesByNameRegex(name string) ([]app.Title, error)
	AddView(id int) error
	GetByRowFilter(row string, ids []int) ([]app.Title, error)
	GetImagesByTitleId(id int) ([]app.Image, error)
}

type Categories interface {
	GetAll() ([]app.Category, error)
	GetByGenres(genres []string) ([]app.Category, error)
}

type Repository struct {
	Authorization
	Titles
	Categories
	Account
}

func NewRepository(db *gorm.DB) *Repository {
	return &Repository{
		Authorization: NewAuthRepo(db),
		Titles:        NewTitleRepo(db),
		Categories:    NewCategoryRepo(db),
		Account:       NewAccountRepo(db),
	}
}
