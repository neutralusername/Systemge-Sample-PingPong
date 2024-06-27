package appWebsocketHTTP

import (
	"Systemge/Error"
	"Systemge/Node"
	"Systemge/Utilities"
	"SystemgeSamplePingPong/topics"
)

type AppWebsocketHTTP struct {
}

func New() Node.WebsocketHTTPApplication {
	return &AppWebsocketHTTP{}
}

func (app *AppWebsocketHTTP) OnStart(node *Node.Node) error {
	err := node.AsyncMessage(topics.PING, node.GetName(), "ping")
	if err != nil {
		node.GetLogger().Log(Error.New("error sending ping message", err).Error())
	}
	return nil
}

func (app *AppWebsocketHTTP) OnStop(node *Node.Node) error {
	err := node.AsyncMessage(topics.PING, node.GetName(), "ping")
	if err != nil {
		node.GetLogger().Log(Error.New("error sending ping message", err).Error())
	}
	println("successfully sent ping message to broker but app's node already stopped due to multi-module stop order.")
	return nil
}

func (app *AppWebsocketHTTP) GetApplicationConfig() Node.ApplicationConfig {
	return Node.ApplicationConfig{
		ResolverAddress:            "127.0.0.1:60000",
		ResolverNameIndication:     "127.0.0.1",
		ResolverTLSCert:            Utilities.GetFileContent("MyCertificate.crt"),
		HandleMessagesSequentially: false,
	}
}
