package service

import (
	"errors"
	"github.com/HeadGardener/link-shortener/internal/app/models"
	"github.com/HeadGardener/link-shortener/internal/app/repository"
	"github.com/golang-jwt/jwt"
	"golang.org/x/crypto/bcrypt"
	"time"
)

const (
	tokenTTL   = 12 * time.Hour
	signingKey = "wmvimowembinwwinbuewn"
)

type AuthService struct {
	repos repository.Authorization
}

func NewAuthService(repos repository.Authorization) *AuthService {
	return &AuthService{repos: repos}
}

type tokenClaims struct {
	jwt.StandardClaims
	UserID string `json:"user_id"`
}

func (s *AuthService) CreateUser(user models.User) (interface{}, error) {
	var err error
	user.Password, err = generatePasswordHash(user.Password)
	if err != nil {
		return 0, err
	}

	return s.repos.CreateUser(user)
}

func generatePasswordHash(password string) (string, error) {
	hashed, err := bcrypt.GenerateFromPassword([]byte(password), 0)

	return string(hashed), err
}

func (s *AuthService) GenerateToken(userInput models.UserInput) (string, error) {
	var err error
	userInput.Password, err = generatePasswordHash(userInput.Password)
	userID, err := s.repos.GetUser(userInput)
	if err != nil {
		return "", err
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &tokenClaims{
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(tokenTTL).Unix(),
			IssuedAt:  time.Now().Unix(),
		},
		userID,
	})

	return token.SignedString([]byte(signingKey))
}

func (s *AuthService) ParseToken(accessToken string) (string, error) {
	token, err := jwt.ParseWithClaims(accessToken, &tokenClaims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("invalid signing method")
		}

		return []byte(signingKey), nil
	})

	if err != nil {
		return "", err
	}

	claims, ok := token.Claims.(*tokenClaims)
	if !ok {
		return "", errors.New("token claims are not of type *tokenClaims")
	}

	return claims.UserID, nil
}
