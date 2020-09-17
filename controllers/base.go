package controllers

import (
	"beegoTool/services"
	"fmt"
	"github.com/astaxie/beego"
)

type Base struct {
	beego.Controller

}

func (this *Base) Prepare() {
	fmt.Println("执行接口初始化前置信息")
	//初始化
}
func (this *Base) Options() {
	this.Data["json"] = map[string]interface{}{"status": 200, "message": "ok", "moreinfo": ""}
	this.ServeJSON()
}
func (this *Base) SuccReturn(retData interface{}) {
	retJson := map[string]interface{}{
		"APISTATUS": "2000",
		"APIDEC":    "ok",
		"APIDATA":   retData,
	}
	if retData == nil {
		retJson["APIDATA"] = []int{}
	}
	this.Data["json"] = retJson
	this.ServeJSON()
	//this.Abort("200")
}

func (this *Base) FailReturn(err services.CskError) {
	retJson := make(map[string]interface{})
	if err == nil {
		retJson["APISTATUS"] = "2000"
		retJson["APIDEC"] = "OK"
	} else {
		retJson[`APIDATA`] = []int{}
		retJson["APISTATUS"], retJson["APIDEC"] = err.Error()
	}

	this.Data["json"] = retJson
	this.ServeJSON()
}

func (this *Base) FailReturnStatus(err services.CskError, status int) {
	retJson := make(map[string]interface{})
	if err == nil {
		retJson["APISTATUS"] = "2000"
		retJson["msg"] = ""
	} else {
		retJson["APISTATUS"], retJson["msg"] = err.Error()
	}
	this.Ctx.ResponseWriter.WriteHeader(status)
	this.Data["json"] = retJson
	this.ServeJSON()
}

