package ws

import (
	"fmt"
	"sync"
)

var HubServer = newHub()

type Hub struct {
	clients    map[int64]*Client //维护所有的client
	mutex      *sync.Mutex
	broadcast  chan []byte  //广播消息
	register   chan *Client //注册
	unregister chan *Client //注销

}

func newHub() *Hub {
	return &Hub{
		mutex:      &sync.Mutex{},
		clients:    make(map[int64]*Client),
		broadcast:  make(chan []byte), //同步管道，确保hub消息不堆积，同时多个client给hub发数据会阻塞
		register:   make(chan *Client),
		unregister: make(chan *Client),
	}
}

func (hub *Hub) GetClient(id int64) (*Client, bool) {

	hub.mutex.Lock()
	defer hub.mutex.Unlock()

	client, ok := hub.clients[id]
	return client, ok
}

func (hub *Hub) SetClient(id int64, client *Client) {
	hub.mutex.Lock()
	defer hub.mutex.Unlock()

	hub.clients[id] = client
}

func (hub *Hub) DelClient(id int64) {
	hub.mutex.Lock()
	defer hub.mutex.Unlock()

	delete(hub.clients, id)
}

func (hub *Hub) Run() {
	for {
		select {
		case client := <-hub.register:
			fmt.Printf("online register:%d\n", client.userID)
			//client上线，注册
			hub.SetClient(client.userID, client)

		case client := <-hub.unregister:
			//查询当前client是否存在
			if _, exist := hub.GetClient(client.userID); exist {
				close(client.Send)
				hub.DelClient(client.userID)
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
