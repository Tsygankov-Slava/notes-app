package service

import (
	"crypto/sha1"
	"errors"
	"fmt"
	"github.com/Tsygankov-Slava/notes-app"
	"github.com/Tsygankov-Slava/notes-app/pkg/repository"
	"github.com/dgrijalva/jwt-go"
	"time"
)

type AuthService struct {
	repo repository.Authorization
}

func NewAuthService(repo repository.Authorization) *AuthService {
	return &AuthService{repo: repo}
}

func (s *AuthService) CreateUser(user notes.User) (int, error) {
	user.Password = generatePasswordHash(user.Password)
	return s.repo.CreateUser(user)
}

const salt = "sdfhlw4r232fsdlj23rlsdf3" // const for hashing

func generatePasswordHash(password string) string {
	/* Using sha1 algorithm for hashing */
	hash := sha1.New()
	hash.Write([]byte(password))

	return fmt.Sprintf("%x", hash.Sum([]byte(salt)))
}

const (
	tokenTTL     = 12 * time.Hour
	signatureKey = "sdjk123ADfh23F4rsdfjlASF23asd12"
)

type tokenClaims struct {
	jwt.StandardClaims
	UserId int `json:"user_id"`
}

func (s *AuthService) GenerateToken(username, password string) (string, error) {
	user, err := s.repo.GetUser(username, generatePasswordHash(password))
	if err != nil {
		return "", err
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, tokenClaims{
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(tokenTTL).Unix(), // The token is valid for 12 hours
			IssuedAt:  time.Now().Unix(),               // Token generation time
		}, user.Id,
	})
	return token.SignedString([]byte(signatureKey)) // Signing the token
}

func (s *AuthService) ParseToken(accessToken string) (int, error) {
	token, err := jwt.ParseWithClaims(accessToken, &tokenClaims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok { // Checking the token signature
			return nil, errors.New("invalid signature method")
		}
		return []byte(signatureKey), nil
	})
	if err != nil {
		return 0, err
	}
	claims, ok := token.Claims.(*tokenClaims)
	if !ok {
		return 0, errors.New("token claims are not of the type *tokenClaims")
	}
	return claims.UserId, nil
}
