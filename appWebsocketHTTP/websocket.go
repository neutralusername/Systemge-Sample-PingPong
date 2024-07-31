package appWebsocketHTTP

import (
	"SystemgeSamplePingPong/topics"

	"github.com/neutralusername/Systemge/Error"
	"github.com/neutralusername/Systemge/Node"
)

func (app *AppWebsocketHTTP) GetWebsocketMessageHandlers() map[string]Node.WebsocketMessageHandler {
	return map[string]Node.WebsocketMessageHandler{}
}

func (app *AppWebsocketHTTP) OnConnectHandler(node *Node.Node, websocketClient *Node.WebsocketClient) {
	reponseChannel, err := node.SyncMessage(topics.PINGPONG, "ping")
	if err != nil {
		if errorLogger := node.GetErrorLogger(); errorLogger != nil {
			errorLogger.Log(Error.New("error sending pingPongSync message", err).Error())
		}
	}
	response, err := reponseChannel.ReceiveResponse()
	if err != nil {
		panic(err)

	}
	if response.GetMessage().GetPayload() != "pong" {
		panic("unexpected response")
	}
	err = node.AsyncMessage(topics.PING, "ping")
	if err != nil {
		if errorLogger := node.GetErrorLogger(); errorLogger != nil {
			errorLogger.Log(Error.New("error sending ping message", err).Error())
		}
	}
}

func (app *AppWebsocketHTTP) OnDisconnectHandler(node *Node.Node, websocketClient *Node.WebsocketClient) {
	println("websocket client disconnected")
}
