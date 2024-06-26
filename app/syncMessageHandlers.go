package app

import (
	"Systemge/Message"
	"Systemge/Node"
	"SystemgeSamplePingPong/topics"
)

func (app *App) GetSyncMessageHandlers() map[string]Node.SyncMessageHandler {
	return map[string]Node.SyncMessageHandler{
		topics.PINGPONG: func(client *Node.Node, message *Message.Message) (string, error) {
			app.pingsReceived++
			return "pong", nil
		},
	}
}
