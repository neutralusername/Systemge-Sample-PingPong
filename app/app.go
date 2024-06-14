package app

import (
	"Systemge/Application"
	"Systemge/Error"
	"Systemge/Message"
	"Systemge/MessageBrokerClient"
	"Systemge/Utilities"
	"SystemgeSamplePingPong/topics"
)

type App struct {
	logger              *Utilities.Logger
	messageBrokerClient *MessageBrokerClient.Client
	pingsReceived       int
}

func New(logger *Utilities.Logger, messageBrokerClient *MessageBrokerClient.Client) Application.Application {
	app := &App{
		logger:              logger,
		messageBrokerClient: messageBrokerClient,
	}
	return app
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
			err := app.messageBrokerClient.AsyncMessage("pong", app.messageBrokerClient.GetName(), "pong")
			if err != nil {
				return Error.New("error sending pong message", err)
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
