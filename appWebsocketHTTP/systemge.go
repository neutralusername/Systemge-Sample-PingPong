package appWebsocketHTTP

import (
	"SystemgeSamplePingPong/topics"

	"github.com/neutralusername/Systemge/Message"
	"github.com/neutralusername/Systemge/Node"
)

func (app *AppWebsocketHTTP) GetAsyncMessageHandlers() map[string]Node.AsyncMessageHandler {
	return map[string]Node.AsyncMessageHandler{
		topics.PONG: func(node *Node.Node, message *Message.Message) error {
			println("PONG")
			node.WebsocketBroadcast(Message.NewAsync("pong", ""))
			return nil
		},
	}
}

func (app *AppWebsocketHTTP) GetSyncMessageHandlers() map[string]Node.SyncMessageHandler {
	return map[string]Node.SyncMessageHandler{}
}
