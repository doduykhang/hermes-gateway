package service

import (
	"github.com/google/uuid"
)

type Token interface {
	CheckToken(token string) (string, error)		
	CreateToken(userID string) (string, error)		
}

type token struct {
	cache Cache
}

func NewToken(cache Cache) Token {
	return &token {
		cache: cache,
	}
}

func (a *token) CheckToken(token string) (string, error) {
	return a.cache.Get(token) 	
}
func (a *token) CreateToken(userID string) (string, error) {
	token := uuid.New().String()
	err := a.cache.Set(token, userID)	
	if err != nil {
		return "", err
	}
	return token, nil
}


