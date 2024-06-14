package main

import (
	"Systemge/Module"
	"SystemgeSamplePingPong/app"
	"SystemgeSamplePingPong/appWebsocket"
)

const TOPICRESOLUTIONSERVER_ADDRESS = "127.0.0.1:60000"
const WEBSOCKET_PORT = ":8443"

const ERROR_LOG_FILE_PATH = "error.log"

func main() {
	Module.NewResolverServerFromConfig("resolver.systemge", ERROR_LOG_FILE_PATH).Start()
	Module.NewBrokerServerFromConfig("brokerApp.systemge", ERROR_LOG_FILE_PATH).Start()
	Module.NewBrokerServerFromConfig("brokerWebsocket.systemge", ERROR_LOG_FILE_PATH).Start()

	clientApp := Module.NewClient("clientApp", TOPICRESOLUTIONSERVER_ADDRESS, ERROR_LOG_FILE_PATH, app.New)
	Module.StartCommandLineInterface(Module.NewMultiModule(
		clientApp,
		Module.NewWebsocketClient("clientWebsocket", TOPICRESOLUTIONSERVER_ADDRESS, ERROR_LOG_FILE_PATH, "/ws", WEBSOCKET_PORT, "", "", appWebsocket.New),
		Module.NewHTTPServerFromConfig("httpServe.systemge", ERROR_LOG_FILE_PATH),
	), clientApp.GetApplication().GetCustomCommandHandlers())
}
