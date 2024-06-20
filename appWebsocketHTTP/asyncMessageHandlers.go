package appWebsocketHTTP

import (
	"Systemge/Application"
	"Systemge/Message"
	"SystemgeSamplePingPong/topics"
)

func (app *AppWebsocketHTTP) GetAsyncMessageHandlers() map[string]Application.AsyncMessageHandler {
	return map[string]Application.AsyncMessageHandler{
		topics.PONG: func(message *Message.Message) error {
			println("PONG")
			app.client.GetWebsocketServer().Broadcast(Message.NewAsync("pong", app.client.GetName(), ""))
			return nil
		},
	}
}
