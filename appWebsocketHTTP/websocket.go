package appWebsocketHTTP

import (
	"Systemge/Error"
	"Systemge/Node"
	"SystemgeSamplePingPong/topics"
)

func (app *AppWebsocketHTTP) GetWebsocketMessageHandlers() map[string]Node.WebsocketMessageHandler {
	return map[string]Node.WebsocketMessageHandler{}
}

func (app *AppWebsocketHTTP) OnConnectHandler(node *Node.Node, websocketClient *Node.WebsocketClient) {
	reponse, err := node.SyncMessage(topics.PINGPONG, websocketClient.GetId(), "ping")
	if err != nil {
		node.GetLogger().Log(Error.New("error sending pingPongSync message", err).Error())
	}
	if reponse.GetPayload() != "pong" {
		node.GetLogger().Log(Error.New("expected pong, got "+reponse.GetPayload(), nil).Error())
	}
	err = node.AsyncMessage(topics.PING, websocketClient.GetId(), "ping")
	if err != nil {
		node.GetLogger().Log(Error.New("error sending ping message", err).Error())
	}
}

func (app *AppWebsocketHTTP) OnDisconnectHandler(node *Node.Node, websocketClient *Node.WebsocketClient) {
	println("websocket client disconnected")
}

func (app *AppWebsocketHTTP) GetWebsocketComponentConfig() Node.WebsocketComponentConfig {
	return Node.WebsocketComponentConfig{
		Pattern:     "/ws",
		Port:        ":8443",
		TlsCertPath: "",
		TlsKeyPath:  "",
	}
}
