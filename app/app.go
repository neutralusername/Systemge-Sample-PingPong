package app

import (
	"Systemge/Config"
	"Systemge/Node"
	"Systemge/Resolution"
	"Systemge/Utilities"
)

type App struct {
	pingsReceived int
}

func New() Node.Application {
	app := &App{}
	return app
}

func (app *App) OnStart(node *Node.Node) error {
	return nil
}

func (app *App) OnStop(node *Node.Node) error {
	return nil
}

func (app *App) GetApplicationConfig() Config.Application {
	return Config.Application{
		ResolverResolution:         Resolution.New("resolver", "127.0.0.1:60000", "127.0.0.1", Utilities.GetFileContent("MyCertificate.crt")),
		HandleMessagesSequentially: false,
	}
}
