package app

import (
	"bufio"
	"context"
	"fmt"
	"github.com/hiteshrepo/token_project/internal/app/client/service"
	"log"
	"os"
	"strings"
)

type App struct {
	ExecuteSvc *service.ExecuteService
	ParserSvc  *service.ParserService
}

func (a *App) Start(ctx context.Context) {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter text: ")
	text, _ := reader.ReadString('\n')

	parts := strings.Split(text, " ")
	var action string
	if action = a.ParserSvc.GetAction(parts); action == "" {
		return
	}

	_, model, err := a.ParserSvc.GetInput(parts, action)
	checkErr(err)

	a.ExecuteSvc.TriggerAction(ctx, action, model)
}

func (a *App) Shutdown() {}

func checkErr(err error) {
	if err != nil {
		log.Fatalf(err.Error())
	}
}
