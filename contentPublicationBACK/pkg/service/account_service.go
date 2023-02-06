package service

import (
	app "contentPublicationBACK"
	"contentPublicationBACK/pkg/repository"
)

type AccountService struct {
	repoAuth repository.Authorization
	repo     repository.Account
}

func NewAccountService(repo repository.Authorization, repoAcc repository.Account) *AccountService {
	return &AccountService{repoAuth: repo, repo: repoAcc}
}

func (s *AccountService) GetUserInfo(id int) (app.User, error) {
	user, err := s.repo.GetUserInfo(id)
	if err != nil {
		return app.User{}, err
	}

	return user, err
}
