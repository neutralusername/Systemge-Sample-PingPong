package appWebsocketHTTP

import (
	"Systemge/Config"
	"Systemge/Error"
	"Systemge/Node"
	"Systemge/TcpServer"
	"SystemgeSamplePingPong/topics"
)

func (app *AppWebsocketHTTP) GetWebsocketMessageHandlers() map[string]Node.WebsocketMessageHandler {
	return map[string]Node.WebsocketMessageHandler{}
}

func (app *AppWebsocketHTTP) OnConnectHandler(node *Node.Node, websocketClient *Node.WebsocketClient) {
	reponse, err := node.SyncMessage(topics.PINGPONG, websocketClient.GetId(), "ping")
	if err != nil {
		node.GetLogger().Error(Error.New("error sending pingPongSync message", err).Error())
	}
	if reponse.GetPayload() != "pong" {
		node.GetLogger().Error(Error.New("expected pong, got "+reponse.GetPayload(), nil).Error())
	}
	err = node.AsyncMessage(topics.PING, websocketClient.GetId(), "ping")
	if err != nil {
		node.GetLogger().Error(Error.New("error sending ping message", err).Error())
	}
}

func (app *AppWebsocketHTTP) OnDisconnectHandler(node *Node.Node, websocketClient *Node.WebsocketClient) {
	println("websocket client disconnected")
}

func (app *AppWebsocketHTTP) GetWebsocketComponentConfig() Config.Websocket {
	return Config.Websocket{
		Pattern:                          "/ws",
		Server:                           TcpServer.New(8443, "", ""),
		HandleClientMessagesSequentially: false,
		ClientMessageCooldownMs:          0,
		ClientWatchdogTimeoutMs:          20000,
	}
}
