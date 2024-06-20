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

	clientApp := Module.NewClient("clientApp", RESOLVER_ADDRESS, ERROR_LOG_FILE_PATH, app.New, nil)
	Module.StartCommandLineInterface(Module.NewMultiModule(
		clientApp,
		Module.NewCompositeClientWebsocketHTTP("clientWebsocketHTTP", RESOLVER_ADDRESS, ERROR_LOG_FILE_PATH, "/ws", WEBSOCKET_PORT, "", "", HTTP_PORT, "", "", appWebsocketHTTP.New, nil),
	), clientApp.GetApplication().GetCustomCommandHandlers())
}
