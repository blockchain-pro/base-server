package server

import (
	"github.com/blockchain-pro/base-server/config"
	"github.com/blockchain-pro/base-server/util/os"
)

var appConfig *config.Config

type BaseApp interface {
	Init()
}

func Init(ConfigPath string, app BaseApp) {
	app.Init()
	os.WaitQuit()
}
