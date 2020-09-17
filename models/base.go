package models

import (
	"beegoTool/global/mysql"
	redisHelp "beegoTool/global/redis"
	"xorm.io/xorm"
)

func getDb() *xorm.Engine {

	return mysql.GetSession("default").Master()
}

func getRedis() *redisHelp.RClient {
	return redisHelp.Redis("default")
}

func GetDbSession() *xorm.Session {
	return getDb().NewSession()
}
