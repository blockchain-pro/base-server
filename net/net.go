package net

import "github.com/blockchain-pro/base-server/config"

func Init(config *config.Config) {
	InitWeb(&config.WebServer)
	InitSocket(&config.SocketServer)
	InitWebSocket(&config.WebSocketServer)

}