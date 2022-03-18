package grpc

import (
	"context"
	"fmt"
	"github.com/hiteshrepo/token_project/internal/pkg/config"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"time"
)

type GrpcConn struct {
	options []grpc.DialOption
	ctx     context.Context
}

const dialTimeout = 10 * time.Second

func ProvideGrpcConn(ctx context.Context) (*GrpcConn, error) {
	grpcConn := &GrpcConn{
		ctx: ctx,
	}

	return grpcConn, nil
}

func ProvideClient(conn *GrpcConn, config config.ServerConfig, opts ...grpc.DialOption) (*grpc.ClientConn, error) {
	url := fmt.Sprintf("%s:%d", config.Host, config.Port)
	conn.options = opts
	ctx, cancel := getContext(conn.ctx)
	cc, err := grpc.DialContext(ctx, url, conn.options...)
	if err != nil {
		cancel()
		return nil, fmt.Errorf("unable to connect to client. address = %s error = %+v", url, err)
	}

	return cc, nil
}

func ProvideGrpcClientOptions() []grpc.DialOption {
	var opts []grpc.DialOption
	opts = append(opts, grpc.WithBlock(), grpc.WithTransportCredentials(insecure.NewCredentials()))
	return opts
}

func getContext(parent context.Context) (context.Context, context.CancelFunc) {
	ctx := parent
	if ctx == nil {
		ctx = context.Background()
	}

	return context.WithTimeout(ctx, dialTimeout)
}