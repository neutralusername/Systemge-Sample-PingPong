package appWebsocketHTTP

import (
	"Systemge/Config"
	"Systemge/Helpers"
	"Systemge/Message"
	"Systemge/Node"
	"SystemgeSamplePingPong/topics"
)

func (app *AppWebsocketHTTP) GetSystemgeComponentConfig() Config.Systemge {
	return Config.Systemge{
		HandleMessagesSequentially: false,

		BrokerSubscribeDelayMs:    1000,
		TopicResolutionLifetimeMs: 10000,
		SyncResponseTimeoutMs:     10000,
		TcpTimeoutMs:              5000,

		ResolverEndpoint: Config.TcpEndpoint{
			Address: "127.0.0.1:60000",
			Domain:  "example.com",
			TlsCert: Helpers.GetFileContent("MyCertificate.crt"),
		},
	}
}

func (app *AppWebsocketHTTP) GetAsyncMessageHandlers() map[string]Node.AsyncMessageHandler {
	return map[string]Node.AsyncMessageHandler{
		topics.PONG: func(node *Node.Node, message *Message.Message) error {
			println("PONG")
			node.WebsocketBroadcast(Message.NewAsync("pong", node.GetName(), ""))
			return nil
		},
	}
}

func (app *AppWebsocketHTTP) GetSyncMessageHandlers() map[string]Node.SyncMessageHandler {
	return map[string]Node.SyncMessageHandler{}
}
