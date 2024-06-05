package appWebsocket

import (
	"Systemge/Error"
	"Systemge/Message"
	"Systemge/MessageBrokerClient"
	"Systemge/Utilities"
	"SystemgeFramework/topics"
)

type WebsocketApp struct {
	logger              *Utilities.Logger
	messageBrokerClient *MessageBrokerClient.Client
}

func New(logger *Utilities.Logger, messageBrokerClient *MessageBrokerClient.Client) MessageBrokerClient.WebsocketApplication {
	return &WebsocketApp{
		logger:              logger,
		messageBrokerClient: messageBrokerClient,
	}
}

func (app *WebsocketApp) GetAsyncMessageHandlers() map[string]MessageBrokerClient.AsyncMessageHandler {
	return map[string]MessageBrokerClient.AsyncMessageHandler{
		topics.PONG: func(message *Message.Message) error {
			println("PONG")
			return nil
		},
	}
}

func (app *WebsocketApp) GetSyncMessageHandlers() map[string]MessageBrokerClient.SyncMessageHandler {
	return map[string]MessageBrokerClient.SyncMessageHandler{}
}

func (app *WebsocketApp) GetCustomCommandHandlers() map[string]MessageBrokerClient.CustomCommandHandler {
	return map[string]MessageBrokerClient.CustomCommandHandler{}
}

func (app *WebsocketApp) GetWebsocketMessageHandlers() map[string]MessageBrokerClient.WebsocketMessageHandler {
	return map[string]MessageBrokerClient.WebsocketMessageHandler{}
}

func (app *WebsocketApp) OnConnectHandler(connection *MessageBrokerClient.WebsocketClient) error {
	reponse, err := app.messageBrokerClient.SyncMessage(Message.NewSync(topics.PINGPONG_SYNC, connection.GetId(), "ping"))
	if err != nil {
		return Error.New("error sending pingPongSync message", err)
	}
	if reponse.Payload != "pong" {
		return Error.New("expected pong, got "+reponse.Payload, nil)
	}
	err = app.messageBrokerClient.AsyncMessage(Message.NewAsync(topics.PING, connection.GetId(), "ping"))
	if err != nil {
		return Error.New("error sending ping message", err)
	}
	return nil
}

func (app *WebsocketApp) OnDisconnectHandler(connection *MessageBrokerClient.WebsocketClient) error {
	return nil
}
