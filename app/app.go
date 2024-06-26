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

func (app *App) OnStart(client *Node.Node) error {
	return nil
}

func (app *App) OnStop(client *Node.Node) error {
	return nil
}
