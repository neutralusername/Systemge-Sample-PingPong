package app

import (
	"Systemge/Application"
	"Systemge/Message"
	"SystemgeSamplePingPong/topics"
)

func (app *App) GetSyncMessageHandlers() map[string]Application.SyncMessageHandler {
	return map[string]Application.SyncMessageHandler{
		topics.PINGPONG: func(message *Message.Message) (string, error) {
			app.pingsReceived++
			return "pong", nil
		},
	}
}
