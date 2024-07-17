package app

import (
	"Systemge/Node"
)

type App struct {
	pingsReceived int
}

func New() Node.Application {
	app := &App{}
	return app
}

func (app *App) GetCustomCommandHandlers() map[string]Node.CustomCommandHandler {
	return map[string]Node.CustomCommandHandler{
		"pingsReceived": func(node *Node.Node, args []string) error {
			println(app.pingsReceived)
			return nil
		},
	}
}
