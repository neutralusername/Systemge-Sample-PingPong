package appWebsocketHTTP

import (
	"Systemge/Message"
	"Systemge/Node"
	"SystemgeSamplePingPong/topics"
)

func (app *AppWebsocketHTTP) GetAsyncMessageHandlers() map[string]Node.AsyncMessageHandler {
	return map[string]Node.AsyncMessageHandler{
		topics.PONG: func(node *Node.Node, message *Message.Message) error {
			println("PONG")
			node.WebsocketBroadcast(Message.NewAsync("pong", node.GetName(), ""))
			return nil
		},
	}
}
