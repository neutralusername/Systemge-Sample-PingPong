package main

import (
	"Systemge/Broker"
	"Systemge/Config"
	"Systemge/Module"
	"Systemge/Node"
	"Systemge/Resolver"
	"Systemge/TcpEndpoint"
	"Systemge/TcpServer"
	"Systemge/Utilities"
	"SystemgeSamplePingPong/app"
	"SystemgeSamplePingPong/appWebsocketHTTP"
	"SystemgeSamplePingPong/topics"
)

const ERROR_LOG_FILE_PATH = "error.log"

func main() {
	err := Node.New(Config.Node{
		Name:   "nodeResolver",
		Logger: Utilities.NewLogger(ERROR_LOG_FILE_PATH, ERROR_LOG_FILE_PATH, ERROR_LOG_FILE_PATH, ERROR_LOG_FILE_PATH, nil),
	}, Resolver.New(Config.Resolver{
		Server:       TcpServer.New(60000, "MyCertificate.crt", "MyKey.key"),
		ConfigServer: TcpServer.New(60001, "MyCertificate.crt", "MyKey.key"),

		TcpTimeoutMs: 5000,
	})).Start()
	if err != nil {
		panic(err)
	}

	err = Node.New(Config.Node{
		Name:   "nodeBrokerApp",
		Logger: Utilities.NewLogger(ERROR_LOG_FILE_PATH, ERROR_LOG_FILE_PATH, ERROR_LOG_FILE_PATH, ERROR_LOG_FILE_PATH, nil),
	}, Broker.New(Config.Broker{
		Server:       TcpServer.New(60002, "MyCertificate.crt", "MyKey.key"),
		Endpoint:     TcpEndpoint.New("127.0.0.1:60002", "example.com", Utilities.GetFileContent("MyCertificate.crt")),
		ConfigServer: TcpServer.New(60003, "MyCertificate.crt", "MyKey.key"),

		SyncTopics:  []string{topics.PINGPONG},
		AsyncTopics: []string{topics.PING},

		ResolverConfigEndpoint: TcpEndpoint.New("127.0.0.1:60001", "example.com", Utilities.GetFileContent("MyCertificate.crt")),

		SyncResponseTimeoutMs: 10000,
		TcpTimeoutMs:          5000,
	})).Start()
	if err != nil {
		panic(err)
	}

	err = Node.New(Config.Node{
		Name:   "nodeBrokerWebsocketHTTP",
		Logger: Utilities.NewLogger(ERROR_LOG_FILE_PATH, ERROR_LOG_FILE_PATH, ERROR_LOG_FILE_PATH, ERROR_LOG_FILE_PATH, nil),
	}, Broker.New(Config.Broker{
		Server:       TcpServer.New(60004, "MyCertificate.crt", "MyKey.key"),
		Endpoint:     TcpEndpoint.New("127.0.0.1:60004", "example.com", Utilities.GetFileContent("MyCertificate.crt")),
		ConfigServer: TcpServer.New(60005, "MyCertificate.crt", "MyKey.key"),

		AsyncTopics: []string{topics.PONG},

		ResolverConfigEndpoint: TcpEndpoint.New("127.0.0.1:60001", "example.com", Utilities.GetFileContent("MyCertificate.crt")),

		SyncResponseTimeoutMs: 10000,
		TcpTimeoutMs:          5000,
	})).Start()
	if err != nil {
		panic(err)
	}

	applicationWebsocketHTTP := appWebsocketHTTP.New()
	Module.StartCommandLineInterface(Module.NewMultiModule(
		Node.New(Config.Node{
			Name:   "nodeApp",
			Logger: Utilities.NewLogger(ERROR_LOG_FILE_PATH, ERROR_LOG_FILE_PATH, ERROR_LOG_FILE_PATH, ERROR_LOG_FILE_PATH, nil),
		}, app.New()),
		Node.New(Config.Node{
			Name:   "nodeAppWebsocketHTTP",
			Logger: Utilities.NewLogger(ERROR_LOG_FILE_PATH, ERROR_LOG_FILE_PATH, ERROR_LOG_FILE_PATH, ERROR_LOG_FILE_PATH, nil),
		}, applicationWebsocketHTTP),
	))
}
