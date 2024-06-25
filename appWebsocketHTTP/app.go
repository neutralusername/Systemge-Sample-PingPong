package appWebsocketHTTP

import (
	"Systemge/Client"
	"Systemge/Utilities"
	"SystemgeSamplePingPong/topics"
)

type AppWebsocketHTTP struct {
}

func New() Client.CompositeApplicationWebsocketHTTP {
	return &AppWebsocketHTTP{}
}

func (app *AppWebsocketHTTP) OnStart(client *Client.Client) error {
	err := client.AsyncMessage(topics.PING, client.GetName(), "ping")
	if err != nil {
		client.GetLogger().Log(Utilities.NewError("error sending ping message", err).Error())
	}
	return nil
}

func (app *AppWebsocketHTTP) OnStop(client *Client.Client) error {
	err := client.AsyncMessage(topics.PING, client.GetName(), "ping")
	if err != nil {
		client.GetLogger().Log(Utilities.NewError("error sending ping message", err).Error())
	}
	println("successfully sent ping message to broker but clientApp already stopped due to multi-module stop order.")
	return nil
}
