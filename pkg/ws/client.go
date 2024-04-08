package ws

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/websocket"
	"net-chat/global"
	"net-chat/pkg/protocol"
	"net/http"
	"time"
)

var (
	pongWait         = 60 * time.Second  //等待时间
	pingPeriod       = 9 * pongWait / 10 //周期54s
	maxMsgSize int64 = 512               //消息最大长度
	writeWait        = 10 * time.Second  //

)
var (
	newLine = []byte{'\n'}
	space   = []byte{' '}
)
var upGrader = websocket.Upgrader{
	HandshakeTimeout: 2 * time.Second, //握手超时时间
	ReadBufferSize:   1024,            //读缓冲大小
	WriteBufferSize:  1024,            //写缓冲大小
	CheckOrigin:      func(r *http.Request) bool { return true },
	Error:            func(w http.ResponseWriter, r *http.Request, status int, reason error) {},
}

func ServeWs(w http.ResponseWriter, r *http.Request, uid int64) {
	conn, err := upGrader.Upgrade(w, r, nil) //http升级为websocket协议
	if err != nil {
		fmt.Printf("upgrade error: %v\n", err)
		return
	}

	fmt.Printf("connect to client %s, uid:%d\n", conn.RemoteAddr().String(), uid)

	//每来一个前端请求，就会创建一个client
	client := &Client{
		conn:   conn,
		Send:   make(chan *protocol.Message),
		userID: uid,
	}

	//向hub注册client
	HubServer.register <- client

	//启动子协程，运行ServeWs的协程退出后子协程也不会能出
	//websocket是全双工模式，可以同时read和write
	go client.read()
	go client.write()
}

type Client struct {
	Send   chan *protocol.Message
	conn   *websocket.Conn
	userID int64
}

func (c *Client) read() {
	defer func() {
		//hub中注销client
		HubServer.unregister <- c
		//global.Log.Infof("uid:%d-client %s disconnect", c.userID, c.conn.RemoteAddr().String())
		//关闭websocket管道
		_ = c.conn.Close()
	}()

	//一次从管管中读取的最大长度
	c.conn.SetReadLimit(maxMsgSize)
	//连接中，每隔54秒向客户端发一次ping，客户端返回pong，所以把SetReadDeadline设为60秒，超过60秒后不允许读
	_ = c.conn.SetReadDeadline(time.Now().Add(pongWait))
	//心跳
	c.conn.SetPongHandler(func(appData string) error {
		//每次收到pong都把deadline往后推迟60秒
		_ = c.conn.SetReadDeadline(time.Now().Add(pongWait))
		return nil
	})

	for {
		//如果前端主动断开连接，运行会报错，for循环会退出。注册client时，hub中会关闭client.send管道
		_, msg, err := c.conn.ReadMessage()
		if err != nil {
			//如果以意料之外的关闭状态关闭，就打印日志
			if websocket.IsUnexpectedCloseError(err, websocket.CloseAbnormalClosure, websocket.CloseGoingAway) {
				fmt.Printf("read from websocket err: %v\n", err)
			}
			//ReadMessage失败，关闭websocket管道、注销client，退出
			break
		} else {
			//换行符替换成空格，去除首尾空格
			fmt.Printf("\n\n消息:%s", msg)

			c.receiveOption(msg)
		}
	}
}

// 从hub的broadcast那儿读限数据，写到websocket连接里面去
func (c *Client) write() {
	//给前端发心跳，看前端是否还存活
	ticker := time.NewTicker(pingPeriod)
	defer func() {
		//ticker不用就stop，防止协程泄漏
		ticker.Stop()
		fmt.Printf("close connection to %s\n", c.conn.RemoteAddr().String())
		//给前端写数据失败，关闭连接
		_ = c.conn.Close()
	}()

	for {
		select {
		//正常情况是hub发来了数据。如果前端断开了连接，read()会触发client.send管道的关闭，该case会立即执行。从而执行!ok里的return，从而执行defer
		case msg, ok := <-c.Send:
			//client.send该管道被hub关闭
			if !ok {
				//写一条关闭信息就可以结束一切
				_ = c.conn.WriteMessage(websocket.CloseMessage, []byte{})
				return
			}
			//10秒内必须把信息写给前端（写到websocket连接里去），否则就关闭连接
			_ = c.conn.SetWriteDeadline(time.Now().Add(writeWait))
			//通过NextWriter创建一个新的writer，主要是为了确保上一个writer已经被关闭，即它想写的内容已经flush到conn里去
			if writer, err := c.conn.NextWriter(websocket.TextMessage); err != nil {
				return
			} else {

				binMsg, _ := json.Marshal(msg)
				_, _ = writer.Write(binMsg)
				_, _ = writer.Write(newLine) //每发一条消息，都加一个换行符
				//为了提升性能，如果c.send里还有消息，则趁这一次都写给前端
				n := len(c.Send)
				for i := 0; i < n; i++ {
					binMsg, _ = json.Marshal(<-c.Send)
					_, _ = writer.Write(binMsg)
					_, _ = writer.Write(newLine)
				}
				if err = writer.Close(); err != nil {
					return //结束一切
				}
			}
		case <-ticker.C:
			//fmt.Printf("\nsend ping message\n")
			_ = c.conn.SetWriteDeadline(time.Now().Add(writeWait))
			//心跳保持，给浏览器发一个PingMessage，等待浏览器返回PongMessage
			if err := c.conn.WriteMessage(websocket.PingMessage, nil); err != nil {
				return //写websocket连接失败，说明连接出问题了，该c可以over了
			}
		}
	}
}

func (c *Client) receiveOption(res []byte) {
	msg := &protocol.Message{}
	err := json.Unmarshal(res, msg)
	if err != nil {

		global.Log.Errorf("用户:%d-数据格式错误%s\n", c.userID, err.Error())
		return
	}
	// 默认发送消息
	c.sendMsg(msg)

}

func (c *Client) sendMsg(msg *protocol.Message) {

	client, ok := HubServer.clients[msg.To]
	if ok {
		msg.From = c.userID
		client.Send <- msg
	} else {
		fmt.Printf("\n未找到此用户:%d", msg.To)
	}
}
