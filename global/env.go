package global

import (
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
	"beegoTool/global/logger"
	"beegoTool/global/mysql"
	"beegoTool/global/redis"
	"runtime"
)

var Cfg = beego.AppConfig


func InitEnv(log *logs.BeeLogger) {
	runtime.GOMAXPROCS(runtime.NumCPU())
	var err error
	logger.InitLogger(log, IsDev())

	//elasticsearch

	//esConfig := elasticsearch6.Config{
	//	Username :esUser,
	//	Password :esPassword,
	//	Addresses : []string{esLink},
	//}
	//err = es.InitEs(`default`,esConfig)
	//if err != nil {
	//	fmt.Print(err)
	//}
	// database
	dbLink := Cfg.String("master_db")
	maxIdleConn, _ := Cfg.Int("master_db_max_idle_conn")
	maxOpenConn, _ := Cfg.Int("master_db_max_open_conn")
	//dbLink := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8", dbUser, dbPass, dbHost, dbPort, dbName) + "&loc=Asia%2FChongqing"

	//slaveDbLink := Cfg.String("slave_db")
	//slaveMaxidleconn, _ := Cfg.Int("slave_db_max_idle_conn")
	//slaveMaxopenconn, _ := Cfg.Int("slave_db_max_open_conn")
	master := mysql.Options{
		Dns:         dbLink,
		MaxIdleConn: maxIdleConn,
		MaxOpenConn: maxOpenConn,
	}
	//读写分离，详情可查看初始化源码
	slaves := make([]mysql.Options, 0)
	//if slaveDbLink != "" {
	//	slaves = append(slaves, mysql.Options{
	//		Dns:         slaveDbLink,
	//		MaxIdleConn: slaveMaxidleconn,
	//		MaxOpenConn: slaveMaxopenconn,
	//	})
	//}
	//mysql连接池
	err = mysql.InitMysql("default", master, slaves, IsDev())
	if err != nil {
		fmt.Println(err)
		//panic(err)
	}

	defaultRedisHost := Cfg.String("default_redis_host")
	defaultRedisPasswd := Cfg.String("default_redis_passwd")
	defaultRedisDb, _ := Cfg.Int("default_redis_db")
	defaultRedisMaxidle, _ := Cfg.Int("default_redis_maxidle")
	defaultRedisMaxactive, _ := Cfg.Int("default_redis_maxactive")

	redisOpt := redis.Options{
		MaxIdle:     defaultRedisMaxidle,
		MaxActive:   defaultRedisMaxactive,
		Host:        defaultRedisHost,
		PassWd:      defaultRedisPasswd,
		DB:          defaultRedisDb,
		IdleTimeout: 180,
	}
	redis.InitRedis("default", redisOpt)


}
func IsDev() bool {
	if beego.BConfig.RunMode != "prod" {
		return true
	}
	return false
}
