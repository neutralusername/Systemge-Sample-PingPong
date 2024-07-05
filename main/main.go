package main

import (
	"Systemge/Broker"
	"Systemge/Config"
	"Systemge/Module"
	"Systemge/Node"
	"Systemge/Resolution"
	"Systemge/Resolver"
	"Systemge/Utilities"
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
	err := Resolver.New(Config.ParseResolverConfigFromFile("resolver.systemge")).Start()
	if err != nil {
		panic(err)
	}
	err = Broker.New(Config.ParseBrokerConfigFromFile("brokerApp.systemge")).Start()
	if err != nil {
		panic(err)
	}
	err = Broker.New(Config.ParseBrokerConfigFromFile("brokerWebsocketHTTP.systemge")).Start()
	if err != nil {
		panic(err)
	}
	applicationWebsocketHTTP := appWebsocketHTTP.New()
	Module.StartCommandLineInterface(Module.NewMultiModule(
		Node.New(Config.Node{
			Name:               "nodeApp",
			LoggerPath:         ERROR_LOG_FILE_PATH,
			ResolverResolution: Resolution.New("resolver", "127.0.0.1:60000", "127.0.0.1", Utilities.GetFileContent("MyCertificate.crt")),
		}, app.New()),
		Node.New(Config.Node{
			Name:               "nodeWebsocketHTTP",
			LoggerPath:         ERROR_LOG_FILE_PATH,
			ResolverResolution: Resolution.New("resolver", "127.0.0.1:60000", "127.0.0.1", Utilities.GetFileContent("MyCertificate.crt")),
		}, applicationWebsocketHTTP),
	))
}
