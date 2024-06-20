package app

import (
	"Systemge/Application"
	"Systemge/Client"
)

type App struct {
	client        *Client.Client
	pingsReceived int
}

func New(client *Client.Client, args []string) (Application.Application, error) {
	app := &App{
		client: client,
	}
	return app, nil
}

func (app *App) OnStart() error {
	return nil
}

func (app *App) OnStop() error {
	return nil
}
