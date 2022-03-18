package service

import (
	"context"
	"github.com/hiteshrepo/token_project/internal/app/model"
)

type ExecuteService struct {
	TokenSvc  *TokenService
}

func ProvideExecuteService(tokenSvc  *TokenService) *ExecuteService {
	return &ExecuteService{TokenSvc: tokenSvc}
}

func (es *ExecuteService) TriggerAction(ctx context.Context, action string, token *model.Token) {
	if action == actionNameCreate {
		es.TokenSvc.CreateToken(ctx, token.Id)
	}

	if action == actionNameWrite {
		es.TokenSvc.WriteToken(ctx, token)
	}

	if action == actionNameRead {
		es.TokenSvc.ReadToken(ctx, token.Id)
	}

	if action == actionNameDrop {
		es.TokenSvc.DropToken(ctx, token.Id)
	}
}
