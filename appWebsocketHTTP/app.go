package appWebsocketHTTP

import (
	"Systemge/Application"
	"Systemge/Client"
	"Systemge/Utilities"
	"SystemgeSamplePingPong/topics"
)

type AppWebsocketHTTP struct {
	client *Client.Client
}

func New(client *Client.Client, args []string) (Application.CompositeApplicationWebsocketHTTP, error) {
	return &AppWebsocketHTTP{
		client: client,
	}, nil
}

func (app *AppWebsocketHTTP) OnStart() error {
	err := app.client.AsyncMessage(topics.PING, app.client.GetName(), "ping")
	if err != nil {
		app.client.GetLogger().Log(Utilities.NewError("error sending ping message", err).Error())
	}
	return nil
}

func (app *AppWebsocketHTTP) OnStop() error {
	err := app.client.AsyncMessage(topics.PING, app.client.GetName(), "ping")
	if err != nil {
		app.client.GetLogger().Log(Utilities.NewError("error sending ping message", err).Error())
	}
	println("successfully sent ping message to broker but clientApp already stopped due to multi-module stop order.")
	return nil
}
