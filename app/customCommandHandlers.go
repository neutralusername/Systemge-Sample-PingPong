package app

import (
	"Systemge/Node"
)

func (app *App) GetCustomCommandHandlers() map[string]Node.CustomCommandHandler {
	return map[string]Node.CustomCommandHandler{
		"pingsReceived": func(node *Node.Node, args []string) error {
			println(app.pingsReceived)
			return nil
		},
	}
}
