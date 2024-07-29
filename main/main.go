package main

import (
	"SystemgeSamplePingPong/app"
	"SystemgeSamplePingPong/appWebsocketHTTP"
	"SystemgeSamplePingPong/topics"

	"github.com/neutralusername/Systemge/Config"
	"github.com/neutralusername/Systemge/Dashboard"
	"github.com/neutralusername/Systemge/Helpers"
	"github.com/neutralusername/Systemge/Node"
	"github.com/neutralusername/Systemge/Tools"
)

const LOGGER_PATH = "logs.log"

func main() {
	Tools.NewLoggerQueue(LOGGER_PATH, 10000)
	Node.New(&Config.Node{
		Name:           "dashboard",
		RandomizerSeed: Tools.GetSystemTime(),
	}, Dashboard.New(&Config.Dashboard{
		Server: &Config.TcpServer{
			Port: 8081,
		},
		NodeStatusIntervalMs:           1000,
		NodeSystemgeCounterIntervalMs:  1000,
		NodeWebsocketCounterIntervalMs: 1000,
		NodeBrokerCounterIntervalMs:    1000,
		NodeResolverCounterIntervalMs:  1000,
		HeapUpdateIntervalMs:           1000,
		NodeSpawnerCounterIntervalMs:   1000,
		NodeHTTPCounterIntervalMs:      1000,
		GoroutineUpdateIntervalMs:      1000,
		AutoStart:                      true,
		AddDashboardToDashboard:        true,
	},
		Node.New(&Config.Node{
			Name:              "nodeResolver",
			RandomizerSeed:    Tools.GetSystemTime(),
			InfoLoggerPath:    LOGGER_PATH,
			WarningLoggerPath: LOGGER_PATH,
			ErrorLoggerPath:   LOGGER_PATH,
		}, Node.NewResolverApplication(&Config.Resolver{
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
			TopicResolutions: map[string]*Config.TcpEndpoint{
				topics.PINGPONG: {
					Address: "127.0.0.1:60002",
					Domain:  "example.com",
					TlsCert: Helpers.GetFileContent("MyCertificate.crt"),
				},
				topics.PING: {
					Address: "127.0.0.1:60002",
					Domain:  "example.com",
					TlsCert: Helpers.GetFileContent("MyCertificate.crt"),
				},
				topics.PONG: {
					Address: "127.0.0.1:60004",
					Domain:  "example.com",
					TlsCert: Helpers.GetFileContent("MyCertificate.crt"),
				},
			},
			TcpTimeoutMs: 5000,
		})),
		Node.New(&Config.Node{
			Name:              "nodeBrokerApp",
			RandomizerSeed:    Tools.GetSystemTime(),
			InfoLoggerPath:    LOGGER_PATH,
			WarningLoggerPath: LOGGER_PATH,
			ErrorLoggerPath:   LOGGER_PATH,
		}, Node.NewBrokerApplication(&Config.Broker{
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

			SyncResponseTimeoutMs: 10000,
			TcpTimeoutMs:          5000,
		})),
		Node.New(&Config.Node{
			Name:              "nodeBrokerWebsocketHTTP",
			RandomizerSeed:    Tools.GetSystemTime(),
			InfoLoggerPath:    LOGGER_PATH,
			WarningLoggerPath: LOGGER_PATH,
			ErrorLoggerPath:   LOGGER_PATH,
		}, Node.NewBrokerApplication(&Config.Broker{
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

			SyncResponseTimeoutMs: 10000,
			TcpTimeoutMs:          5000,
		})),
		Node.New(&Config.Node{
			Name:              "nodeApp",
			RandomizerSeed:    Tools.GetSystemTime(),
			InfoLoggerPath:    LOGGER_PATH,
			WarningLoggerPath: LOGGER_PATH,
			ErrorLoggerPath:   LOGGER_PATH,
		}, app.New()),
		Node.New(&Config.Node{
			Name:              "nodeAppWebsocketHTTP",
			RandomizerSeed:    Tools.GetSystemTime(),
			InfoLoggerPath:    LOGGER_PATH,
			WarningLoggerPath: LOGGER_PATH,
			ErrorLoggerPath:   LOGGER_PATH,
		}, appWebsocketHTTP.New()),
	)).StartBlocking()
}
