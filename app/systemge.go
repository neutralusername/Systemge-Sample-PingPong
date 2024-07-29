package app

import (
	"SystemgeSamplePingPong/topics"

	"github.com/neutralusername/Systemge/Config"
	"github.com/neutralusername/Systemge/Helpers"
	"github.com/neutralusername/Systemge/Message"
	"github.com/neutralusername/Systemge/Node"
)

func (app *App) GetSystemgeComponentConfig() *Config.Systemge {
	return &Config.Systemge{
		HandleMessagesSequentially: false,

		BrokerSubscribeDelayMs:    1000,
		TopicResolutionLifetimeMs: 10000,
		SyncResponseTimeoutMs:     10000,
		TcpTimeoutMs:              5000,

		ResolverEndpoints: []*Config.TcpEndpoint{
			{
				Address: "127.0.0.1:60000",
				Domain:  "example.com",
				TlsCert: Helpers.GetFileContent("MyCertificate.crt"),
			},
		},
	}
}

func (app *App) GetSyncMessageHandlers() map[string]Node.SyncMessageHandler {
	return map[string]Node.SyncMessageHandler{}
}

func (app *App) GetAsyncMessageHandlers() map[string]Node.AsyncMessageHandler {
	return map[string]Node.AsyncMessageHandler{
		topics.PING: func(node *Node.Node, message *Message.Message) error {
			app.pingsReceived++
			println(app.pingsReceived)
			return nil
		},
	}
}
