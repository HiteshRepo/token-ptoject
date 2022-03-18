package main

import (
	"context"
	"github.com/hiteshrepo/token_project/internal/app/client/di"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	app, err := di.InitializeApp(ctx, cancel)
	check(err)

	app.Start(ctx)

	app.Shutdown()
}

func check(err error) {
	if err != nil {
		panic(err)
	}
}
