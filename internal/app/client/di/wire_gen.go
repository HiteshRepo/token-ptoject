// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package di

import (
	"context"
	"github.com/google/wire"
	"github.com/hiteshrepo/token_project/internal/app/client/app"
	"github.com/hiteshrepo/token_project/internal/app/client/grpc"
	"github.com/hiteshrepo/token_project/internal/app/client/service"
	"github.com/hiteshrepo/token_project/internal/pkg/config"
	"github.com/hiteshrepo/token_project/internal/pkg/mapper"
	"github.com/hiteshrepo/token_project/internal/pkg/proto"
)

// Injectors from wire.go:

func InitializeApp(ctx context.Context, cancel context.CancelFunc) (*app.App, error) {
	grpcConn, err := grpc.ProvideGrpcConn(ctx)
	if err != nil {
		return nil, err
	}
	appConfig, err := config.ProvideAppConfig()
	if err != nil {
		return nil, err
	}
	serverConfig := appConfig.ServerConfig
	v := grpc.ProvideGrpcClientOptions()
	clientConn, err := grpc.ProvideClient(grpcConn, serverConfig, v...)
	if err != nil {
		return nil, err
	}
	tokenServiceClient := tokenv1.NewTokenServiceClient(clientConn)
	mapperMapper := mapper.ProvideMapper()
	tokenService := service.ProvideTokenServiceClient(tokenServiceClient, mapperMapper)
	executeService := service.ProvideExecuteService(tokenService)
	parserService := service.ProvideParserService()
	appApp := &app.App{
		ExecuteSvc: executeService,
		ParserSvc:  parserService,
	}
	return appApp, nil
}

// wire.go:

var configSet = wire.NewSet(config.ProvideAppConfig, wire.FieldsOf(new(config.AppConfig), "ServerConfig"))

var grpcSet = wire.NewSet(grpc.ProvideGrpcConn, grpc.ProvideGrpcClientOptions, grpc.ProvideClient)

var serviceSet = wire.NewSet(tokenv1.NewTokenServiceClient, service.ProvideExecuteService, service.ProvideParserService, service.ProvideTokenServiceClient)