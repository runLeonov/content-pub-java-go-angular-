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

func (s *AccountService) ChangeUserInfo(newUser app.User, id int) (app.User, error) {
	user, err := s.repoAcc.GetUserInfo(id)

	user.LastName = newUser.LastName
	user.FirstName = newUser.FirstName
	user.Name = newUser.Name

	err = s.repoAcc.SaveUserInfo(user)

	if err != nil {
		return app.User{}, err
	}

	return user, err
}

func (s *AccountService) GetLikedTitlesByUserId(id int) ([]app.Title, error) {
	return s.repoAcc.GetUserLikes(id)
}

func (s *AccountService) ChangeRole(role string, id int) error {
	user, err := s.repoAcc.GetUserInfo(id)

	user.Role = role
	err = s.repoAcc.SaveUserInfo(user)

	if err != nil {
		return err
	}

	return err
}

func (s *AccountService) GetLikedTitlesByUserIdLimit(id int) ([]app.Title, error) {
	return s.repoAcc.GetUserLikesLimit(id)
}

func (s *AccountService) GetCommentedTitlesByUserIdLimit(id int) ([]app.Title, error) {
	return s.repoAcc.GetUserCommentsLimit(id)
}

func (s *AccountService) GetCommentedTitlesByUserId(id int) ([]app.Title, error) {
	return s.repoAcc.GetUserComments(id)
}
