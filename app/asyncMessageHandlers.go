package app

import (
	"Systemge/Client"
	"Systemge/Message"
	"Systemge/Utilities"
	"SystemgeSamplePingPong/topics"
)

func (app *App) GetAsyncMessageHandlers() map[string]Client.AsyncMessageHandler {
	return map[string]Client.AsyncMessageHandler{
		topics.PING: func(client *Client.Client, message *Message.Message) error {
			app.pingsReceived++
			err := client.AsyncMessage("pong", client.GetName(), "pong")
			if err != nil {
				return Utilities.NewError("error sending pong message", err)
			}
			return nil
		},
	}
}
