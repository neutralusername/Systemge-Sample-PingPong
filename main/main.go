package main

import (
	"Systemge/Module"
	"SystemgeFramework/app"
	"SystemgeFramework/appWebsocket"
)

const TOPICRESOLUTIONSERVER_ADDRESS = ":60000"
const HTTP_DEV_PORT = ":8080"
const WEBSOCKET_PORT = ":8443"

const ERROR_LOG_FILE_PATH = "error.log"

func main() {
	clientApp := Module.NewClient("clientApp", TOPICRESOLUTIONSERVER_ADDRESS, ERROR_LOG_FILE_PATH, app.New)
	Module.StartCommandLineInterface(Module.NewMultiModule(
		Module.NewResolverServerFromConfig("resolver.systemge", ERROR_LOG_FILE_PATH),
		Module.NewBrokerServerFromConfig("brokerApp.systemge", ERROR_LOG_FILE_PATH),
		Module.NewBrokerServerFromConfig("brokerWebsocket.systemge", ERROR_LOG_FILE_PATH),
		clientApp,
		Module.NewWebsocketClient("clientWebsocket", TOPICRESOLUTIONSERVER_ADDRESS, ERROR_LOG_FILE_PATH, "/ws", WEBSOCKET_PORT, "", "", appWebsocket.New),
		Module.NewHTTPServerFromConfig("httpServe.systemge", ERROR_LOG_FILE_PATH),
	), clientApp.GetApplication().GetCustomCommandHandlers())
}
