package appWebsocket

import (
	"Systemge/Application"
	"Systemge/Client"
	"Systemge/Error"
	"Systemge/Message"
	"Systemge/WebsocketClient"
	"SystemgeSamplePingPong/topics"
)

type WebsocketApp struct {
	client *Client.Client
}

func New(client *Client.Client, args []string) Application.WebsocketApplication {
	return &WebsocketApp{
		client: client,
	}
}

func (app *WebsocketApp) OnStart() error {
	err := app.client.AsyncMessage(topics.PING, app.client.GetName(), "ping")
	if err != nil {
		app.client.GetLogger().Log(Error.New("error sending ping message", err).Error())
	}
	return nil
}

func (app *WebsocketApp) OnStop() error {
	err := app.client.AsyncMessage(topics.PING, app.client.GetName(), "ping")
	if err != nil {
		app.client.GetLogger().Log(Error.New("error sending ping message", err).Error())
	}
	println("successfully sent ping message to broker but clientApp already stopped due to multi-module stop order.")
	return nil
}

func (app *WebsocketApp) GetAsyncMessageHandlers() map[string]Application.AsyncMessageHandler {
	return map[string]Application.AsyncMessageHandler{
		topics.PONG: func(message *Message.Message) error {
			println("PONG")
			app.client.GetWebsocketServer().Broadcast(Message.NewAsync("pong", app.client.GetName(), ""))
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
	reponse, err := app.client.SyncMessage(topics.PINGPONG, connection.GetId(), "ping")
	if err != nil {
		app.client.GetLogger().Log(Error.New("error sending pingPongSync message", err).Error())
	}
	if reponse.GetPayload() != "pong" {
		app.client.GetLogger().Log(Error.New("expected pong, got "+reponse.GetPayload(), nil).Error())
	}
	err = app.client.AsyncMessage(topics.PING, connection.GetId(), "ping")
	if err != nil {
		app.client.GetLogger().Log(Error.New("error sending ping message", err).Error())
	}
}

func (app *WebsocketApp) OnDisconnectHandler(connection *WebsocketClient.Client) {
	println("websocket client disconnected")
}
