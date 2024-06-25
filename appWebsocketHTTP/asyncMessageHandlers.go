package appWebsocketHTTP

import (
	"Systemge/Client"
	"Systemge/Message"
	"SystemgeSamplePingPong/topics"
)

func (app *AppWebsocketHTTP) GetAsyncMessageHandlers() map[string]Client.AsyncMessageHandler {
	return map[string]Client.AsyncMessageHandler{
		topics.PONG: func(client *Client.Client, message *Message.Message) error {
			println("PONG")
			client.WebsocketBroadcast(Message.NewAsync("pong", client.GetName(), ""))
			return nil
		},
	}
}
