package service

import (
	"Internship_backend_avito/internal/entity"
	"Internship_backend_avito/internal/repository/postgresdb"
	"Internship_backend_avito/pkg/hasher"
	"github.com/dgrijalva/jwt-go"
	"time"
)

const (
	signingKey = "opofpajdskvisvieorfd"
	tokenTTL   = 12 * time.Hour
)

type tokenClaims struct {
	jwt.StandardClaims
	UserId int `json:"user_id"`
}

type AuthServices struct {
	repo postgresdb.Authorization
}

func NewAuthServices(repo postgresdb.Authorization) *AuthServices {
	return &AuthServices{repo: repo}
}

func (s *AuthServices) CreateUser(input AuthCreateUserInput) (int, error) {
	user := entity.User{
		Username: input.Username,
		Password: hasher.GeneratePasswordHash(input.Password),
	}
	return s.repo.CreateUser(user)
}

func (s *AuthServices) GenerateToken(input AuthGenerateTokenInput) (string, error) {
	user, err := s.repo.GetUser(input.Username, hasher.GeneratePasswordHash(input.Password))
	if err != nil {
		return "", err
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &tokenClaims{
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(tokenTTL).Unix(),
			IssuedAt:  time.Now().Unix(),
		},
		user.Id,
	})
	return token.SignedString([]byte(signingKey))
}
