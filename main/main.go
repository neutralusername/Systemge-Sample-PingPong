package main

import (
	"Systemge/Broker"
	"Systemge/Config"
	"Systemge/Helpers"
	"Systemge/Node"
	"Systemge/Resolver"
	"SystemgeSamplePingPong/app"
	"SystemgeSamplePingPong/appWebsocketHTTP"
	"SystemgeSamplePingPong/topics"
)

const ERROR_LOG_FILE_PATH = "error.log"

func main() {
	Node.StartCommandLineInterface(true,
		Node.New(&Config.Node{
			Name: "nodeResolver",
			Logger: &Config.Logger{
				InfoPath:    ERROR_LOG_FILE_PATH,
				DebugPath:   ERROR_LOG_FILE_PATH,
				ErrorPath:   ERROR_LOG_FILE_PATH,
				WarningPath: ERROR_LOG_FILE_PATH,
				QueueBuffer: 10000,
			},
		}, Resolver.New(&Config.Resolver{
			Server: &Config.TcpServer{
				Port:        60000,
				TlsCertPath: "MyCertificate.crt",
				TlsKeyPath:  "MyKey.key",
			},
			ConfigServer: &Config.TcpServer{
				Port:        60001,
				TlsCertPath: "MyCertificate.crt",
				TlsKeyPath:  "MyKey.key",
			},
			TcpTimeoutMs: 5000,
		})),
		Node.New(&Config.Node{
			Name: "nodeBrokerApp",
			Logger: &Config.Logger{
				InfoPath:    ERROR_LOG_FILE_PATH,
				DebugPath:   ERROR_LOG_FILE_PATH,
				ErrorPath:   ERROR_LOG_FILE_PATH,
				WarningPath: ERROR_LOG_FILE_PATH,
				QueueBuffer: 10000,
			},
		}, Broker.New(&Config.Broker{
			Server: &Config.TcpServer{
				Port:        60002,
				TlsCertPath: "MyCertificate.crt",
				TlsKeyPath:  "MyKey.key",
			},
			Endpoint: &Config.TcpEndpoint{
				Address: "127.0.0.1:60002",
				Domain:  "example.com",
				TlsCert: Helpers.GetFileContent("MyCertificate.crt"),
			},
			ConfigServer: &Config.TcpServer{
				Port:        60003,
				TlsCertPath: "MyCertificate.crt",
				TlsKeyPath:  "MyKey.key",
			},

			SyncTopics:  []string{topics.PINGPONG},
			AsyncTopics: []string{topics.PING},

			ResolverConfigEndpoint: &Config.TcpEndpoint{
				Address: "127.0.0.1:60001",
				Domain:  "example.com",
				TlsCert: Helpers.GetFileContent("MyCertificate.crt"),
			},

			SyncResponseTimeoutMs: 10000,
			TcpTimeoutMs:          5000,
		})),
		Node.New(&Config.Node{
			Name: "nodeBrokerWebsocketHTTP",
			Logger: &Config.Logger{
				InfoPath:    ERROR_LOG_FILE_PATH,
				DebugPath:   ERROR_LOG_FILE_PATH,
				ErrorPath:   ERROR_LOG_FILE_PATH,
				WarningPath: ERROR_LOG_FILE_PATH,
				QueueBuffer: 10000,
			},
		}, Broker.New(&Config.Broker{
			Server: &Config.TcpServer{
				Port:        60004,
				TlsCertPath: "MyCertificate.crt",
				TlsKeyPath:  "MyKey.key",
			},
			Endpoint: &Config.TcpEndpoint{
				Address: "127.0.0.1:60004",
				Domain:  "example.com",
				TlsCert: Helpers.GetFileContent("MyCertificate.crt"),
			},
			ConfigServer: &Config.TcpServer{
				Port:        60005,
				TlsCertPath: "MyCertificate.crt",
				TlsKeyPath:  "MyKey.key",
			},

			AsyncTopics: []string{topics.PONG},

			ResolverConfigEndpoint: &Config.TcpEndpoint{
				Address: "127.0.0.1:60001",
				Domain:  "example.com",
				TlsCert: Helpers.GetFileContent("MyCertificate.crt"),
			},

			SyncResponseTimeoutMs: 10000,
			TcpTimeoutMs:          5000,
		})),
		Node.New(&Config.Node{
			Name: "nodeApp",
			Logger: &Config.Logger{
				InfoPath:    ERROR_LOG_FILE_PATH,
				DebugPath:   ERROR_LOG_FILE_PATH,
				ErrorPath:   ERROR_LOG_FILE_PATH,
				WarningPath: ERROR_LOG_FILE_PATH,
				QueueBuffer: 10000,
			},
		}, app.New()),
		Node.New(&Config.Node{
			Name: "nodeAppWebsocketHTTP",
			Logger: &Config.Logger{
				InfoPath:    ERROR_LOG_FILE_PATH,
				DebugPath:   ERROR_LOG_FILE_PATH,
				ErrorPath:   ERROR_LOG_FILE_PATH,
				WarningPath: ERROR_LOG_FILE_PATH,
				QueueBuffer: 10000,
			},
		}, appWebsocketHTTP.New()),
	)
}
