package routers

import (
	"beegoTool/controllers"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/plugins/cors"
	"github.com/json-iterator/go/extra"
)

func init() {
	beego.InsertFilter("*", beego.BeforeRouter, cors.Allow(&cors.Options{
		AllowAllOrigins:  true,
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Authorization", "Access-Control-Allow-Origin", "Access-Control-Allow-Headers", "Content-Type", "X-Api-Token", "x-requested-with"},
		ExposeHeaders:    []string{"Content-Length", "Access-Control-Allow-Origin", "Access-Control-Allow-Headers", "Content-Type"},
		AllowCredentials: true,
	}))
	ns := beego.NewNamespace("/*",
		//  用于跨域请求
		beego.NSRouter("*", &controllers.Base{}, "OPTIONS:Options"))
	beego.AddNamespace(ns)
	// 容忍字符串和数字互转
	extra.RegisterFuzzyDecoders()
	beego.Router("/",&controllers.Index{},"*:GetIndex")
}
