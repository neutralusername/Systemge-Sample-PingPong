package app

import (
	"Systemge/Client"
)

func (app *App) GetCustomCommandHandlers() map[string]Client.CustomCommandHandler {
	return map[string]Client.CustomCommandHandler{
		"pingsReceived": func(client *Client.Client, args []string) error {
			println(app.pingsReceived)
			return nil
		},
	}
}
