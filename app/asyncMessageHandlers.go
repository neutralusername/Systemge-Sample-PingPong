package app

import (
	"Systemge/Error"
	"Systemge/Message"
	"Systemge/Node"
	"SystemgeSamplePingPong/topics"
)

func (app *App) GetAsyncMessageHandlers() map[string]Node.AsyncMessageHandler {
	return map[string]Node.AsyncMessageHandler{
		topics.PING: func(node *Node.Node, message *Message.Message) error {
			app.pingsReceived++
			err := node.AsyncMessage("pong", node.GetName(), "pong")
			if err != nil {
				return Error.New("error sending pong message", err)
			}
			return nil
		},
	}
}
