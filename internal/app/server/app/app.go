package app

import (
	"fmt"
	"github.com/hiteshrepo/token_project/internal/app/server/handler"
	tokenv1 "github.com/hiteshrepo/token_project/internal/pkg/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"log"
	"net"
)

type App struct {
	Server      *grpc.Server
	Listener    net.Listener
	TickHandler *handler.TokenHandler
}

func (app *App) Start(checkErr func(err error)) {
	app.registerServers()
	go func() {
		log.Println(fmt.Sprintf("GRPC Server started at %s", app.Listener.Addr()))
		err := app.Server.Serve(app.Listener)
		checkErr(err)
	}()
}

func (app *App) registerServers() {
	reflection.Register(app.Server)
	tokenv1.RegisterTokenServiceServer(app.Server, app.TickHandler)
}

func (app *App) Shutdown() {
	app.Server.GracefulStop()
}