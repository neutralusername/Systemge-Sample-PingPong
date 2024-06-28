package app

import (
	"Systemge/Config"
	"Systemge/Node"
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
		ResolverAddress:            "127.0.0.1:60000",
		ResolverNameIndication:     "127.0.0.1",
		ResolverTLSCert:            Utilities.GetFileContent("MyCertificate.crt"),
		HandleMessagesSequentially: false,
	}
}
