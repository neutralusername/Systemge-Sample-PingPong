package appWebsocketHTTP

import (
	"Systemge/Application"
	"Systemge/Utilities"
	"Systemge/WebsocketClient"
	"SystemgeSamplePingPong/topics"
)

func (app *AppWebsocketHTTP) GetWebsocketMessageHandlers() map[string]Application.WebsocketMessageHandler {
	return map[string]Application.WebsocketMessageHandler{}
}

func (app *AppWebsocketHTTP) OnConnectHandler(connection *WebsocketClient.Client) {
	reponse, err := app.client.SyncMessage(topics.PINGPONG, connection.GetId(), "ping")
	if err != nil {
		app.client.GetLogger().Log(Utilities.NewError("error sending pingPongSync message", err).Error())
	}
	if reponse.GetPayload() != "pong" {
		app.client.GetLogger().Log(Utilities.NewError("expected pong, got "+reponse.GetPayload(), nil).Error())
	}
	err = app.client.AsyncMessage(topics.PING, connection.GetId(), "ping")
	if err != nil {
		app.client.GetLogger().Log(Utilities.NewError("error sending ping message", err).Error())
	}
}

func (app *AppWebsocketHTTP) OnDisconnectHandler(connection *WebsocketClient.Client) {
	println("websocket client disconnected")
}
