package main

import (
	"Systemge/Module"
	"SystemgeSamplePingPong/app"
	"SystemgeSamplePingPong/appWebsocketHTTP"
)

const RESOLVER_ADDRESS = "127.0.0.1:60000"
const WEBSOCKET_PORT = ":8443"
const HTTP_PORT = ":8080"

const ERROR_LOG_FILE_PATH = "error.log"

func main() {
	Module.NewResolverFromConfig("resolver.systemge", ERROR_LOG_FILE_PATH).Start()
	Module.NewBrokerFromConfig("brokerApp.systemge", ERROR_LOG_FILE_PATH).Start()
	Module.NewBrokerFromConfig("brokerWebsocket.systemge", ERROR_LOG_FILE_PATH).Start()

	clientApp := Module.NewClient(&Module.ClientConfig{
		Name:            "clientApp",
		ResolverAddress: RESOLVER_ADDRESS,
		LoggerPath:      ERROR_LOG_FILE_PATH,
	}, app.New, nil)
	clientWebsocketHTTP := Module.NewCompositeClientWebsocketHTTP(&Module.ClientConfig{
		Name:             "clientWebsocketHTTP",
		ResolverAddress:  RESOLVER_ADDRESS,
		LoggerPath:       ERROR_LOG_FILE_PATH,
		WebsocketPattern: "/ws",
		WebsocketPort:    WEBSOCKET_PORT,
		WebsocketCert:    "",
		WebsocketKey:     "",
		HTTPPort:         HTTP_PORT,
		HTTPCert:         "",
		HTTPKey:          "",
	}, appWebsocketHTTP.New, nil)
	Module.StartCommandLineInterface(Module.NewMultiModule(
		clientApp,
		clientWebsocketHTTP,
	), clientApp.GetApplication().GetCustomCommandHandlers())
}
