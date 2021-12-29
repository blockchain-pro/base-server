package database

import (
	"github.com/blockchain-pro/base-server/config"
	"github.com/blockchain-pro/base-server/database/mongodb"
	"github.com/blockchain-pro/base-server/database/mysql"
	"github.com/blockchain-pro/base-server/database/redis"
)

type DataBase struct {
	MysqlDB *mysql.MysqlDB
	MongoDB *mongodb.MongoDB
	RedisDB *redis.RedisDB
}

func New(config *config.Config) *DataBase {
	newDateBase := new(DataBase)
	newDateBase.MysqlDB = mysql.New(&config.MySql)
	newDateBase.RedisDB = redis.New(&config.Redis)
	newDateBase.MongoDB = mongodb.New(&config.MongoDB)
	return newDateBase
}
