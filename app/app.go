package app

import (
	"github.com/neutralusername/Systemge/Helpers"
	"github.com/neutralusername/Systemge/Node"
)

type App struct {
	pingsReceived int
}

func New() Node.Application {
	app := &App{}
	return app
}

func (app *App) GetCommandHandlers() map[string]Node.CommandHandler {
	return map[string]Node.CommandHandler{
		"pingsReceived": func(node *Node.Node, args []string) (string, error) {
			return Helpers.IntToString(app.pingsReceived), nil
		},
	}
}
