package appWebsocketHTTP

import (
	"Systemge/Error"
	"Systemge/Node"
	"SystemgeSamplePingPong/topics"
)

func (app *AppWebsocketHTTP) GetWebsocketMessageHandlers() map[string]Node.WebsocketMessageHandler {
	return map[string]Node.WebsocketMessageHandler{}
}

func (app *AppWebsocketHTTP) OnConnectHandler(client *Node.Node, websocketClient *Node.WebsocketClient) {
	reponse, err := client.SyncMessage(topics.PINGPONG, websocketClient.GetId(), "ping")
	if err != nil {
		client.GetLogger().Log(Error.New("error sending pingPongSync message", err).Error())
	}
	if reponse.GetPayload() != "pong" {
		client.GetLogger().Log(Error.New("expected pong, got "+reponse.GetPayload(), nil).Error())
	}
	err = client.AsyncMessage(topics.PING, websocketClient.GetId(), "ping")
	if err != nil {
		client.GetLogger().Log(Error.New("error sending ping message", err).Error())
	}
}

func (app *AppWebsocketHTTP) OnDisconnectHandler(client *Node.Node, websocketClient *Node.WebsocketClient) {
	println("websocket client disconnected")
}
