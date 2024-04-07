package ws

import (
	"fmt"
	"net-chat/pkg/protocol"
	"sync"
)

var HubServer = newHub()

type Hub struct {
	clients    map[int64]*Client //维护所有的client
	mutex      *sync.Mutex
	broadcast  chan *protocol.Message //广播消息
	register   chan *Client           //注册
	unregister chan *Client           //注销

}

func newHub() *Hub {
	return &Hub{
		mutex:      &sync.Mutex{},
		clients:    make(map[int64]*Client),
		broadcast:  make(chan *protocol.Message), //同步管道，确保hub消息不堆积，同时多个client给hub发数据会阻塞
		register:   make(chan *Client),
		unregister: make(chan *Client),
	}
}

func (hub *Hub) Run() {
	for {
		select {
		case client := <-hub.register:
			fmt.Printf("上线注册:%d\n", client.userID)
			//client上线，注册
			hub.clients[client.userID] = client
		case client := <-hub.unregister:
			//查询当前client是否存在
			if _, exists := hub.clients[client.userID]; exists {
				//注销client 通道
				close(client.Send)
				//删除注销的client
				delete(hub.clients, client.userID)
			}
		case msg := <-hub.broadcast:
			//将message广播给每一位client
			for _, client := range hub.clients {
				select {
				case client.Send <- msg:
				//异常client处理
				default:
					close(client.Send)
					//删除异常的client
					delete(hub.clients, client.userID)
				}
			}
		}
	}
}
