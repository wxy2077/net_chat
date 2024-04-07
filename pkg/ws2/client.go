package ws2

import (
	"fmt"
	"github.com/gogo/protobuf/proto"
	"github.com/gorilla/websocket"
	"net-chat/global"
	"net-chat/pkg/protocol"
	"net/http"
	"time"
)

var upGrader = websocket.Upgrader{
	HandshakeTimeout: 2 * time.Second, //握手超时时间
	ReadBufferSize:   1024,            //读缓冲大小
	WriteBufferSize:  1024,            //写缓冲大小
	CheckOrigin:      func(r *http.Request) bool { return true },
	Error:            func(w http.ResponseWriter, r *http.Request, status int, reason error) {},
}

type Client struct {
	Conn *websocket.Conn
	Name string
	Send chan []byte
}

func (c *Client) Read() {
	defer func() {
		MyServer.Unregister <- c
		_ = c.Conn.Close()
	}()

	for {
		c.Conn.PongHandler()
		_, message, err := c.Conn.ReadMessage()
		if err != nil {
			MyServer.Unregister <- c
			_ = c.Conn.Close()
			break
		}

		msg := &protocol.Message{}
		_ = proto.Unmarshal(message, msg)

		// pong
		if msg.Type == global.HeatBeat {
			pong := &protocol.Message{
				Content: global.PONG,
				Type:    global.HeatBeat,
			}
			pongByte, err2 := proto.Marshal(pong)
			if nil != err2 {
				//log.Logger.Error("client marshal message error", log.Any("client marshal message error", err2.Error()))
			}
			_ = c.Conn.WriteMessage(websocket.BinaryMessage, pongByte)
		} else {
			//if config.GetConfig().MsgChannelType.ChannelType == constant.KAFKA {
			//	kafka.Send(message)
			//} else {
			MyServer.Broadcast <- message
			//}
		}
	}
}

func (c *Client) Write() {
	defer func() {
		_ = c.Conn.Close()
	}()

	for {
		select {
		//正常情况是server发来了数据。如果前端断开了连接，read()会触发client.send管道的关闭，该case会立即执行。从而执行!ok里的return，从而执行defer
		case message, ok := <-c.Send:
			//client.send该管道被hub关闭
			if !ok {
				//写一条关闭信息就可以结束一切
				_ = c.Conn.WriteMessage(websocket.CloseMessage, []byte{})
				return
			}
			_ = c.Conn.WriteMessage(websocket.BinaryMessage, message)
		}
	}
}

func ServeWs(w http.ResponseWriter, r *http.Request, uid string) {

	conn, err := upGrader.Upgrade(w, r, nil) //http升级为websocket协议
	if err != nil {
		fmt.Printf("upgrade error: %v\n", err)

		return
	}
	fmt.Printf("connect to client %s\n", conn.RemoteAddr().String())

	client := &Client{
		Name: uid,
		Conn: conn,
		Send: make(chan []byte),
	}

	//注册client
	MyServer.Register <- client
	//启动子协程，运行ServeWs的协程退出后子协程也不会能出
	//websocket是全双工模式，可以同时read和write
	go client.Read()
	go client.Write()
}
