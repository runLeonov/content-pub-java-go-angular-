package service

import (
	app "contentPublicationBACK"
	"contentPublicationBACK/pkg/repository"
)

type Authorization interface {
	CreateUser(user app.User) (int, error)
	GenerateToken(email, password string) (string, error)
	ParseToken(token string) (int, error)
}

type Titles interface {
	GetTitles() ([]app.Title, error)
	GetTitle(id int) (app.Title, error)
	CreateTitle(title app.Title) (uint, error)
	UpdateTitle(title app.Title, id int) (uint, error)
	DeleteTitle(id int) (uint, error)
	GetTitlesByCategories(categories []app.Category) ([]app.Title, error)
}

type Service struct {
	Titles
	Authorization
}

func NewService(repo *repository.Repository) *Service {
	return &Service{
		Authorization: NewAuthService(repo.Authorization),
		Titles:        NewTitleService(repo.Titles, repo.Categories),
	}
}
