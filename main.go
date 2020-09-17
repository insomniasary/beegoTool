/**
 * 入口文件
 * Author: yansheng
 * RegTime: 2019/5/18
 */
package main

import (
	"beegoTool/global"
	_ "beegoTool/routers"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
)

func main() {
	log := logs.NewLogger()
	log.SetLogger(logs.AdapterFile, `{"filename":"./tmp/process.log","level":7,"maxlines":0,"maxsize":0,"daily":true,"maxdays":7,"color":false}`)
	log.Async()
	global.InitEnv(log)
	beego.Run()
}
