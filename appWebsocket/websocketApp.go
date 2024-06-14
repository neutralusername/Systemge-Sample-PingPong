package appWebsocket

import (
	"Systemge/Application"
	"Systemge/Error"
	"Systemge/Message"
	"Systemge/MessageBrokerClient"
	"Systemge/Utilities"
	"Systemge/WebsocketClient"
	"SystemgeSamplePingPong/topics"
)

type WebsocketApp struct {
	logger              *Utilities.Logger
	messageBrokerClient *MessageBrokerClient.Client
}

func New(logger *Utilities.Logger, messageBrokerClient *MessageBrokerClient.Client) Application.WebsocketApplication {
	return &WebsocketApp{
		logger:              logger,
		messageBrokerClient: messageBrokerClient,
	}
}

func (app *WebsocketApp) OnStart() error {
	err := app.messageBrokerClient.AsyncMessage(topics.PING, app.messageBrokerClient.GetName(), "ping")
	if err != nil {
		app.logger.Log(Error.New("error sending ping message", err).Error())
	}
	return nil
}

func (app *WebsocketApp) OnStop() error {
	err := app.messageBrokerClient.AsyncMessage(topics.PING, app.messageBrokerClient.GetName(), "ping")
	if err != nil {
		app.logger.Log(Error.New("error sending ping message", err).Error())
	}
	println("successfully sent ping message to broker but clientApp already stopped due to multi-module stop order.")
	return nil
}

func (app *WebsocketApp) GetAsyncMessageHandlers() map[string]Application.AsyncMessageHandler {
	return map[string]Application.AsyncMessageHandler{
		topics.PONG: func(message *Message.Message) error {
			println("PONG")
			app.messageBrokerClient.GetWebsocketServer().Broadcast(Message.NewAsync("pong", app.messageBrokerClient.GetName(), ""))
			return nil
		},
	}
}

func (app *WebsocketApp) GetSyncMessageHandlers() map[string]Application.SyncMessageHandler {
	return map[string]Application.SyncMessageHandler{}
}

func (app *WebsocketApp) GetCustomCommandHandlers() map[string]Application.CustomCommandHandler {
	return map[string]Application.CustomCommandHandler{}
}

func (app *WebsocketApp) GetWebsocketMessageHandlers() map[string]Application.WebsocketMessageHandler {
	return map[string]Application.WebsocketMessageHandler{}
}

func (app *WebsocketApp) OnConnectHandler(connection *WebsocketClient.Client) {
	reponse, err := app.messageBrokerClient.SyncMessage(topics.PINGPONG, connection.GetId(), "ping")
	if err != nil {
		app.logger.Log(Error.New("error sending pingPongSync message", err).Error())
	}
	if reponse.GetPayload() != "pong" {
		app.logger.Log(Error.New("expected pong, got "+reponse.GetPayload(), nil).Error())
	}
	err = app.messageBrokerClient.AsyncMessage(topics.PING, connection.GetId(), "ping")
	if err != nil {
		app.logger.Log(Error.New("error sending ping message", err).Error())
	}
}

func (app *WebsocketApp) OnDisconnectHandler(connection *WebsocketClient.Client) {
	println("websocket client disconnected")
}
