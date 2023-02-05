package service

import (
	app "contentPublicationBACK"
	"contentPublicationBACK/pkg/repository"
	"crypto/sha1"
	"errors"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"time"
)

const (
	salt       = "asdhnbgalkgdhkgafskjl642"
	signingKey = "aeigh%UY#!*R(THQFG()oagi"
)

type AuthService struct {
	repo repository.Authorization
}

type tokenClaims struct {
	jwt.StandardClaims
	UserId int `json:"user_id"`
}

func NewAuthService(repo repository.Authorization) *AuthService {
	return &AuthService{repo: repo}
}

func (s *AuthService) CreateUser(user app.User) (int, error) {
	user.Password = generatePasswordHash(user.Password)

	return s.repo.CreateUser(user)
}

func (s *AuthService) GenerateToken(email, password string) (string, error) {
	user, err := s.repo.GetUser(email, generatePasswordHash(password))
	if err != nil {
		return "", err
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &tokenClaims{
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(12 * time.Hour).Unix(),
			IssuedAt:  time.Now().Unix(),
		},
		user.ID},
	)

	return token.SignedString([]byte(signingKey))
}

func (s *AuthService) GetUser(id int) (app.User, error) {
	user, err := s.repo.GetUserById(id)
	if err != nil {
		return app.User{}, err
	}

	return user, err
}

func (s *AuthService) CheckUserExist(email string) bool {
	exist := s.repo.CheckExist(email)
	return exist
}

func (s *AuthService) ParseToken(token string) (int, error) {
	tok, err := jwt.ParseWithClaims(token, &tokenClaims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("invalid singing method")
		}

		return []byte(signingKey), nil
	})

	if err != nil {
		return 0, err
	}
	claims, ok := tok.Claims.(*tokenClaims)
	if !ok {
		return 0, errors.New("token claims are not of type *tokenClaims")
	}

	return claims.UserId, nil
}

func generatePasswordHash(password string) string {
	hash := sha1.New()
	hash.Write([]byte(password))

	return fmt.Sprintf("%x", hash.Sum([]byte(salt)))
}
