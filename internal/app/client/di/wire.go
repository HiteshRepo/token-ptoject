//go:build wireinject
// +build wireinject

//go:generate wire

package di

import (
	"context"
	"github.com/google/wire"
	"github.com/hiteshrepo/token_project/internal/app/client/app"
	"github.com/hiteshrepo/token_project/internal/app/client/grpc"
	"github.com/hiteshrepo/token_project/internal/app/client/service"
	"github.com/hiteshrepo/token_project/internal/pkg/config"
	"github.com/hiteshrepo/token_project/internal/pkg/mapper"
	tokenv1 "github.com/hiteshrepo/token_project/internal/pkg/proto"
)

var configSet = wire.NewSet(
	config.ProvideAppConfig,
	wire.FieldsOf(new(config.AppConfig), "ServerConfig"),
)

var grpcSet = wire.NewSet(
	grpc.ProvideGrpcConn,
	grpc.ProvideGrpcClientOptions,
	grpc.ProvideClient,
)

var serviceSet = wire.NewSet(
	tokenv1.NewTokenServiceClient,
	service.ProvideExecuteService,
	service.ProvideParserService,
	service.ProvideTokenServiceClient,
)

func InitializeApp(ctx context.Context, cancel context.CancelFunc) (*app.App, error) {
	wire.Build(
		configSet,
		grpcSet,
		serviceSet,
		mapper.ProvideMapper,
		wire.Struct(new(app.App), "*"),
	)

	return &app.App{}, nil
}
