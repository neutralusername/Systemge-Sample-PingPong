package appWebsocketHTTP

import (
	"Systemge/Error"
	"Systemge/Node"
	"SystemgeSamplePingPong/topics"
)

type AppWebsocketHTTP struct {
}

func New() Node.WebsocketHTTPApplication {
	return &AppWebsocketHTTP{}
}

func (app *AppWebsocketHTTP) OnStart(client *Node.Node) error {
	err := client.AsyncMessage(topics.PING, client.GetName(), "ping")
	if err != nil {
		client.GetLogger().Log(Error.New("error sending ping message", err).Error())
	}
	return nil
}

func (app *AppWebsocketHTTP) OnStop(client *Node.Node) error {
	err := client.AsyncMessage(topics.PING, client.GetName(), "ping")
	if err != nil {
		client.GetLogger().Log(Error.New("error sending ping message", err).Error())
	}
	println("successfully sent ping message to broker but clientApp already stopped due to multi-module stop order.")
	return nil
}
