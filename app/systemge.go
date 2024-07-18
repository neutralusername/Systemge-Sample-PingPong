package app

import (
	"Systemge/Config"
	"Systemge/Error"
	"Systemge/Helpers"
	"Systemge/Message"
	"Systemge/Node"
	"Systemge/Tcp"
	"SystemgeSamplePingPong/topics"
)

func (app *App) GetSystemgeComponentConfig() Config.Systemge {
	return Config.Systemge{
		HandleMessagesSequentially: false,

		BrokerSubscribeDelayMs:    1000,
		TopicResolutionLifetimeMs: 10000,
		SyncResponseTimeoutMs:     10000,
		TcpTimeoutMs:              5000,

		ResolverEndpoint: Tcp.NewEndpoint("127.0.0.1:60000", "example.com", Helpers.GetFileContent("MyCertificate.crt")),
	}
}

func (app *App) GetSyncMessageHandlers() map[string]Node.SyncMessageHandler {
	return map[string]Node.SyncMessageHandler{
		topics.PINGPONG: func(node *Node.Node, message *Message.Message) (string, error) {
			app.pingsReceived++
			return "pong", nil
		},
	}
}

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
