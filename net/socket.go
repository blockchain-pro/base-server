package net

import (
	"github.com/blockchain-pro/base-server/config"
	"fmt"
)

func InitSocket(cf *config.SocketServer) {
	if !cf.Enable {
		return
	}
	fmt.Println("Socket Enable Port :", cf.PORT)


}
