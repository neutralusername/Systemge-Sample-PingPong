package app

import (
	"Systemge/Client"
	"Systemge/Message"
	"SystemgeSamplePingPong/topics"
)

func (app *App) GetSyncMessageHandlers() map[string]Client.SyncMessageHandler {
	return map[string]Client.SyncMessageHandler{
		topics.PINGPONG: func(client *Client.Client, message *Message.Message) (string, error) {
			app.pingsReceived++
			return "pong", nil
		},
	}
}
