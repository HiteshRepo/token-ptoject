package main

import (
	"github.com/hiteshrepo/token_project/internal/app/server/di"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	app, err := di.InitializeApp()
	check(err)

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
