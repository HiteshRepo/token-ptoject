package service

import (
	"context"
	"fmt"
	"github.com/hiteshrepo/token_project/internal/app/model"
	"github.com/hiteshrepo/token_project/internal/pkg/mapper"
	tokenv1 "github.com/hiteshrepo/token_project/internal/pkg/proto"
	"log"
)

type TokenService struct {
	tokenSvcClient tokenv1.TokenServiceClient
	mapper         mapper.Mapper
}

func ProvideTokenServiceClient(tokenSvcClient tokenv1.TokenServiceClient, mapper mapper.Mapper) *TokenService {
	return &TokenService{tokenSvcClient: tokenSvcClient, mapper: mapper}
}

func (ts *TokenService) CreateToken(ctx context.Context, id int64) {
	req := &tokenv1.CreateTokenRequest{Id: id}
	resp, err := ts.tokenSvcClient.CreateToken(ctx, req)
	if err != nil {
		log.Println(fmt.Sprintf("some error occured: %v", err))
		log.Println(fmt.Sprintf("status received: %v", resp.Status))
		return
	}

	log.Println(fmt.Sprintf("token created successfully"))
}

func (ts *TokenService) WriteToken(ctx context.Context, token *model.Token) {
	req := &tokenv1.WriteTokenRequest{Token: ts.mapper.TokenPb(token)}
	resp, err := ts.tokenSvcClient.WriteToken(ctx, req)
	if err != nil {
		log.Println(fmt.Sprintf("some error occured: %v", err))
		log.Println(fmt.Sprintf("status received: %v", resp.Status))
		return
	}

	log.Println(fmt.Sprintf("token updated successfully"))
}

func (ts *TokenService) ReadToken(ctx context.Context, id int64) {
	req := &tokenv1.ReadTokenRequest{Id: id}
	resp, err := ts.tokenSvcClient.ReadToken(ctx, req)
	if err != nil {
		log.Println(fmt.Sprintf("some error occured: %v", err))
		log.Println(fmt.Sprintf("status received: %v", resp.Status))
		return
	}

	log.Println(fmt.Sprintf("token fetched successfully, token details: %v", resp.Token))
}

func (ts *TokenService) DropToken(ctx context.Context, id int64) {
	req := &tokenv1.DropTokenRequest{Id: id}
	resp, err := ts.tokenSvcClient.DropToken(ctx, req)
	if err != nil {
		log.Println(fmt.Sprintf("some error occured: %v", err))
		log.Println(fmt.Sprintf("status received: %v", resp.Status))
		return
	}

	log.Println(fmt.Sprintf("token dropped successfully"))
}
