package app

import (
	"Systemge/Application"
	"Systemge/Client"
	"Systemge/Message"
	"Systemge/Utilities"
	"SystemgeSamplePingPong/topics"
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

func (app *App) GetAsyncMessageHandlers() map[string]Application.AsyncMessageHandler {
	return map[string]Application.AsyncMessageHandler{
		topics.PING: func(message *Message.Message) error {
			app.pingsReceived++
			err := app.client.AsyncMessage("pong", app.client.GetName(), "pong")
			if err != nil {
				return Utilities.NewError("error sending pong message", err)
			}
			return nil
		},
	}
}

func (app *App) GetSyncMessageHandlers() map[string]Application.SyncMessageHandler {
	return map[string]Application.SyncMessageHandler{
		topics.PINGPONG: func(message *Message.Message) (string, error) {
			app.pingsReceived++
			return "pong", nil
		},
	}
}

func (app *App) GetCustomCommandHandlers() map[string]Application.CustomCommandHandler {
	return map[string]Application.CustomCommandHandler{
		"pingsReceived": func(args []string) error {
			println(app.pingsReceived)
			return nil
		},
	}
}
