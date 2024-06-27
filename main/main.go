package main

import (
	"Systemge/Module"
	"Systemge/Node"
	"SystemgeSamplePingPong/app"
	"SystemgeSamplePingPong/appWebsocketHTTP"
)

const RESOLVER_ADDRESS = "127.0.0.1:60000"
const RESOLVER_NAME_INDICATION = "127.0.0.1"
const RESOLVER_TLS_CERT_PATH = "MyCertificate.crt"
const WEBSOCKET_PORT = ":8443"
const HTTP_PORT = ":8080"

const ERROR_LOG_FILE_PATH = "error.log"

func main() {
	Module.NewResolverFromConfig("resolver.systemge", ERROR_LOG_FILE_PATH).Start()
	Module.NewBrokerFromConfig("brokerApp.systemge", ERROR_LOG_FILE_PATH).Start()
	Module.NewBrokerFromConfig("brokerWebsocket.systemge", ERROR_LOG_FILE_PATH).Start()

	nodeApp := Module.NewNode(&Node.NodeConfig{
		Name:       "nodeApp",
		LoggerPath: ERROR_LOG_FILE_PATH,
	}, app.New(), nil, nil)
	applicationWebsocketHTTP := appWebsocketHTTP.New()
	nodeWebsocketHTTP := Module.NewNode(&Node.NodeConfig{
		Name:       "nodeWebsocketHTTP",
		LoggerPath: ERROR_LOG_FILE_PATH,
	}, applicationWebsocketHTTP, applicationWebsocketHTTP, applicationWebsocketHTTP)
	Module.StartCommandLineInterface(Module.NewMultiModule(
		nodeApp,
		nodeWebsocketHTTP,
	))
}
