package app

import "Systemge/Application"

func (app *App) GetCustomCommandHandlers() map[string]Application.CustomCommandHandler {
	return map[string]Application.CustomCommandHandler{
		"pingsReceived": func(args []string) error {
			println(app.pingsReceived)
			return nil
		},
	}
}
