package service

import (
	app "contentPublicationBACK"
	"contentPublicationBACK/pkg/repository"
)

type AccountService struct {
	repoAuth repository.Authorization
	repoAcc  repository.Account
}

func NewAccountService(repo repository.Authorization, repoAcc repository.Account) *AccountService {
	return &AccountService{repoAuth: repo, repoAcc: repoAcc}
}

func (s *AccountService) GetUserInfo(id int) (app.User, error) {
	user, err := s.repoAcc.GetUserInfo(id)
	if err != nil {
		return app.User{}, err
	}

	return user, err
}

func (s *AccountService) GetLikedTitlesByUserId(id int) ([]app.Title, error) {
	return s.repoAcc.GetUserLikes(id)
}
