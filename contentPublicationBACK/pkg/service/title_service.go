package service

import (
	app "contentPublicationBACK"
	"contentPublicationBACK/pkg/repository"
)

type TitleService struct {
	repoTitles     repository.Titles
	repoCategories repository.Categories
}

func NewTitleService(repo repository.Titles, repoCate repository.Categories) *TitleService {
	return &TitleService{repoTitles: repo, repoCategories: repoCate}
}

func (s *TitleService) GetTitles() ([]app.Title, error) {
	return s.repoTitles.GetAllTitles()
}

func (s *TitleService) GetAllPossibleContent() (app.AllContent, error) {
	return s.repoTitles.GetAllPossibleContent()
}

func (s *TitleService) GetTitle(id int) (app.Title, error) {
	return s.repoTitles.GetTitleById(id)
}

func (s *TitleService) GetRandom() (app.Title, error) {
	return s.repoTitles.GetRandom()
}

func (s *TitleService) CreateTitle(title app.Title) (uint, error) {
	return s.repoTitles.SaveNewTitle(title)
}

func (s *TitleService) UpdateTitle(title app.Title, id int) (uint, error) {
	return s.repoTitles.UpdateTitle(title, id)
}

func (s *TitleService) DeleteTitle(id int) (uint, error) {
	return s.repoTitles.DeleteTitleById(id)
}

func (s *TitleService) GetTitlesByCategories(categories []app.Category) ([]app.Title, error) {
	var arrInt []uint
	for _, category := range categories {
		arrInt = append(arrInt, category.ID)
	}
	return s.repoTitles.GetTitlesByCategories(arrInt)
}
