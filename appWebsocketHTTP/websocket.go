package appWebsocketHTTP

import (
	"SystemgeSamplePingPong/topics"
	"time"

	"github.com/neutralusername/Systemge/Error"
	"github.com/neutralusername/Systemge/Node"
)

func (app *AppWebsocketHTTP) GetWebsocketMessageHandlers() map[string]Node.WebsocketMessageHandler {
	return map[string]Node.WebsocketMessageHandler{}
}

func (app *AppWebsocketHTTP) OnConnectHandler(node *Node.Node, websocketClient *Node.WebsocketClient) {
	reponseChannel, err := node.SyncMessage(topics.PINGPONG, "ping")
	if err != nil {
		panic(err)
	}
	println("sent ping-sync")
	response, err := reponseChannel.ReceiveResponse()
	if err != nil {
		panic(err)

	}
	if response.GetMessage().GetPayload() != "pong" {
		panic("unexpected response")
	}
	println("received pong-sync")

	startedAt := time.Now()
	for i := 0; i < 100000; i++ {
		go func() {
			err := node.AsyncMessage(topics.PING, "")
			if err != nil {
				if errorLogger := node.GetErrorLogger(); errorLogger != nil {
					errorLogger.Log(Error.New("error sending ping message", err).Error())
				}
			}
		}()
	}
	println("100000 pings sent in " + time.Since(startedAt).String())
}

func (app *AppWebsocketHTTP) OnDisconnectHandler(node *Node.Node, websocketClient *Node.WebsocketClient) {
	println("websocket client disconnected")
}
