package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

var wsupgraders = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

func registerClient(c *gin.Context) {

	wsupgraders.CheckOrigin = func(r *http.Request) bool { return true }
	conn, err := wsupgraders.Upgrade(c.Writer, c.Request, nil)

	if err != nil {
		msg := fmt.Sprintf("Failed to set websocket upgrade: %+v", err)
		fmt.Println(msg)
		return
	}

	time.Sleep(time.Second * 5)
	res := getQuote()
	//sendMsg(res)
	sendWp(res)
	// mType, mByte, err := conn.ReadMessage()
	// fmt.Println("mByte: ", string(mByte))
	// fmt.Println("mType: ", mType)
	// fmt.Println("err: ", err)

	//conn.WriteMessage(websocket.TextMessage, []byte(fmt.Sprintf("%s", res)))

	conn.Close()
}
