package app

import (
	"Systemge/Client"
	"Systemge/Error"
	"Systemge/Message"
	"SystemgeSamplePingPong/topics"
)

func (app *App) GetAsyncMessageHandlers() map[string]Client.AsyncMessageHandler {
	return map[string]Client.AsyncMessageHandler{
		topics.PING: func(client *Client.Client, message *Message.Message) error {
			app.pingsReceived++
			err := client.AsyncMessage("pong", client.GetName(), "pong")
			if err != nil {
				return Error.New("error sending pong message", err)
			}
			return nil
		},
	}
}
