package initialize

import "net-chat/pkg/ws"

func InitWsHub() {
	go ws.HubServer.Run()
}
