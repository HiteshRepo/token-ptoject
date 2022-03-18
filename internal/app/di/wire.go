//go:build wireinject
// +build wireinject

//go:generate wire

package di

import (
	"github.com/hiteshrepo/token_project/internal/app"
	"github.com/hiteshrepo/token_project/internal/app/handler"
	"github.com/hiteshrepo/token_project/internal/app/mapper"
	"github.com/hiteshrepo/token_project/internal/app/repository"
	"github.com/hiteshrepo/token_project/internal/app/service"
	"github.com/hiteshrepo/token_project/internal/pkg/config"
	"github.com/hiteshrepo/token_project/internal/pkg/grpc"
	"github.com/google/wire"
)

var configSet = wire.NewSet(
	config.ProvideAppConfig,
	wire.FieldsOf(new(config.AppConfig), "ServerConfig"),
)

var grpcSet = wire.NewSet(
	grpc.ProvideListener,
	grpc.ProvideGrpcServerOptions,
	grpc.ProvideServer,
)

var repoSet = wire.NewSet(
	repository.ProvideTokenRepository,
)

var serviceSet = wire.NewSet(
	service.ProvideTokenService,
)

var handlerSet = wire.NewSet(
	handler.ProvideTokenHandler,
)

func InitializeApp() (*app.App, error) {
	wire.Build(
		configSet,
		grpcSet,
		repoSet,
		serviceSet,
		handlerSet,
		mapper.ProvideMapper,
		wire.Struct(new(app.App), "*"),
	)

	return &app.App{}, nil
}
