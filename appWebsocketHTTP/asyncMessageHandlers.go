package appWebsocketHTTP

import (
	"Systemge/Message"
	"Systemge/Node"
	"SystemgeSamplePingPong/topics"
)

func (app *AppWebsocketHTTP) GetAsyncMessageHandlers() map[string]Node.AsyncMessageHandler {
	return map[string]Node.AsyncMessageHandler{
		topics.PONG: func(client *Node.Node, message *Message.Message) error {
			println("PONG")
			client.WebsocketBroadcast(Message.NewAsync("pong", client.GetName(), ""))
			return nil
		},
	}
}
