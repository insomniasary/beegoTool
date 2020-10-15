package webSocket

import (
	"beegoTool/controllers"
	"fmt"
	"github.com/gorilla/websocket"
	"net/http"
)

type WebsocketBase struct {
	controllers.Base
	Con *websocket.Conn
}
func (this *WebsocketBase) Prepare() {
	con,err := (&websocket.Upgrader{CheckOrigin: func(r *http.Request) bool {
		return true
	}}).Upgrade(this.Ctx.ResponseWriter, this.Ctx.Request, nil)
	if _, ok := err.(websocket.HandshakeError); ok {
		fmt.Println(ok)
	}
	this.Con = con
}