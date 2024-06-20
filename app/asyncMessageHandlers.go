package app

import (
	"Systemge/Application"
	"Systemge/Message"
	"Systemge/Utilities"
	"SystemgeSamplePingPong/topics"
)

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
