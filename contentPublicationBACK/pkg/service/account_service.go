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

func (s *AccountService) CreateSubscription(subs app.Subscription) (uint, error) {
	id, err := s.repoAcc.CreateUserSub(subs)
	if err != nil {
		return id, err
	}

	return id, err
}

func (s *AccountService) DeleteSubscription(subs app.Subscription, user app.User) (uint, error) {
	id, err := s.repoAcc.DeleteUserSub(subs, user)
	if err != nil {
		return id, err
	}

	return id, err
}

func (s *AccountService) GetAuthorSubscribers(authorId uint) ([]app.Subscription, error) {
	return s.repoAcc.GetAuthorSubscribers(authorId)
}

func (s *AccountService) GetUserSubscriptions(userId uint) ([]app.Subscription, error) {
	return s.repoAcc.GetUserSubscriptions(userId)
}

func (s *AccountService) CheckUserSubForAuthor(authorId uint, user app.User) (bool, error) {
	subs, err := s.repoAcc.GetAuthorSubscribers(authorId)

	if err != nil {
		return false, err
	}

	for _, sub := range subs {
		if uint(user.ID) == sub.SubscriberID {
			return true, nil
		}

	}

	return false, err
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

func (s *AccountService) GetUserPublishedByUserId(userId int) ([]app.Title, error) {
	return s.repoAcc.GetUserPublished(userId)
}
