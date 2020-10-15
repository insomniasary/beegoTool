package webSocket

import (
	"fmt"
	"github.com/gorilla/websocket"
)

var clients = make(map[string]*websocket.Conn,0)
type WebSocket struct {
	WebsocketBase
}

func (WebSocket *WebSocket) WS(){

	name:= WebSocket.GetString(`name`)
	if clients[name] != nil {
		join(name,WebSocket.Con)
		//用户加入
	}

	defer func() {
		leave(name)
		//用户离开
		WebSocket.Con.Close()
	}()


	for {
		_,msgStr,err := WebSocket.Con.ReadMessage()

		if err != nil {
			break
		}else{
			fmt.Println(string(msgStr))
			WebSocket.Con.WriteMessage(1,[]byte("success"))
		}
		//fmt.Println(msgStr)

	}

}

func join(name string,conn *websocket.Conn){
	clients[name] = conn
}
func leave(name string){
	delete(clients,name)

}
