//go:build wireinject
// +build wireinject

//go:generate wire

package di

import (
	"github.com/google/wire"
	"github.com/hiteshrepo/token_project/internal/app/server/app"
	"github.com/hiteshrepo/token_project/internal/app/server/grpc"
	"github.com/hiteshrepo/token_project/internal/app/server/handler"
	"github.com/hiteshrepo/token_project/internal/app/server/repository"
	"github.com/hiteshrepo/token_project/internal/app/server/service"
	"github.com/hiteshrepo/token_project/internal/pkg/config"
	"github.com/hiteshrepo/token_project/internal/pkg/mapper"
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
	service.ProvideHashService,
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
