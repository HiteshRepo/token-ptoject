package service

import (
	"github.com/hiteshrepo/token_project/internal/app/model"
	"github.com/hiteshrepo/token_project/internal/app/repository"
)

type TokenService struct {
	tokenRepo *repository.TokenRepository
}

func ProvideTokenService(tokenRepo *repository.TokenRepository) *TokenService {
	return &TokenService{tokenRepo: tokenRepo}
}

func (ts TokenService) Create(id int64) {
	ts.tokenRepo.Create(id)
}

func (ts TokenService) Write(token *model.Token) error {
	err := ts.tokenRepo.Write(token)
	if err != nil {
		return err
	}
	return nil
}

func (ts TokenService) Read(id int64) (*model.Token, error) {
	return ts.tokenRepo.Read(id)
}

func (ts TokenService) Drop(id int64) error {
	return ts.tokenRepo.Drop(id)
}
