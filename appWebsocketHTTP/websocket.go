package appWebsocketHTTP

import (
	"Systemge/Client"
	"Systemge/Utilities"
	"SystemgeSamplePingPong/topics"
)

func (app *AppWebsocketHTTP) GetWebsocketMessageHandlers() map[string]Client.WebsocketMessageHandler {
	return map[string]Client.WebsocketMessageHandler{}
}

func (app *AppWebsocketHTTP) OnConnectHandler(client *Client.Client, websocketClient *Client.WebsocketClient) {
	reponse, err := client.SyncMessage(topics.PINGPONG, websocketClient.GetId(), "ping")
	if err != nil {
		client.GetLogger().Log(Utilities.NewError("error sending pingPongSync message", err).Error())
	}
	if reponse.GetPayload() != "pong" {
		client.GetLogger().Log(Utilities.NewError("expected pong, got "+reponse.GetPayload(), nil).Error())
	}
	err = client.AsyncMessage(topics.PING, websocketClient.GetId(), "ping")
	if err != nil {
		client.GetLogger().Log(Utilities.NewError("error sending ping message", err).Error())
	}
}

func (app *AppWebsocketHTTP) OnDisconnectHandler(client *Client.Client, websocketClient *Client.WebsocketClient) {
	println("websocket client disconnected")
}
