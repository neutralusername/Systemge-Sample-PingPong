package app

import (
	"Systemge/Error"
	"Systemge/Message"
	"Systemge/MessageBrokerClient"
	"Systemge/Utilities"
	"SystemgeFramework/topics"
)

type App struct {
	logger              *Utilities.Logger
	messageBrokerClient *MessageBrokerClient.Client
	pingsReceived       int
}

func New(logger *Utilities.Logger, messageBrokerClient *MessageBrokerClient.Client) MessageBrokerClient.Application {
	app := &App{
		logger:              logger,
		messageBrokerClient: messageBrokerClient,
	}
	return app
}

func (app *App) GetAsyncMessageHandlers() map[string]MessageBrokerClient.AsyncMessageHandler {
	return map[string]MessageBrokerClient.AsyncMessageHandler{
		topics.PING: func(message *Message.Message) error {
			app.pingsReceived++
			err := app.messageBrokerClient.AsyncMessage(Message.NewAsync("pong", app.messageBrokerClient.GetName(), "pong"))
			if err != nil {
				return Error.New("error sending pong message", err)
			}
			return nil
		},
	}
}

func (app *App) GetSyncMessageHandlers() map[string]MessageBrokerClient.SyncMessageHandler {
	return map[string]MessageBrokerClient.SyncMessageHandler{
		topics.PINGPONG_SYNC: func(message *Message.Message) (string, error) {
			app.pingsReceived++
			return "pong", nil
		},
	}
}

func (app *App) GetCustomCommandHandlers() map[string]MessageBrokerClient.CustomCommandHandler {
	return map[string]MessageBrokerClient.CustomCommandHandler{
		"pingsReceived": func(args []string) error {
			println(app.pingsReceived)
			return nil
		},
	}
}
