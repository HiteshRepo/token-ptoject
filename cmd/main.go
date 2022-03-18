package main

import (
	"github.com/hiteshrepo/token_project/internal/app/di"
	"gopkg.in/DataDog/dd-trace-go.v1/ddtrace/tracer"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	app, err := di.InitializeApp()
	check(err)

	tracer.Start()
	defer tracer.Stop()

	app.Start(check)

	<-interrupt()
	app.Shutdown()
}

func interrupt() chan os.Signal {
	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, syscall.SIGINT, syscall.SIGTERM)
	return interrupt
}

func check(err error) {
	if err != nil {
		panic(err)
	}
}
