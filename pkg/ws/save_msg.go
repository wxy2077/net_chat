package ws

import (
	"net-chat/global"
	"net-chat/model"
	"time"
)

var (
	// MessageChannel 消息通道 批量保存消息
	MessageChannel      = make(chan *model.Message, 100)
	GroupMessageChannel = make(chan *model.GroupMessage, 100)

	batchInterval = 10 * time.Minute // 定时器间隔10分钟, 保存一次channel中的数据
)

func SaveMsg() {
	var batchSingleMsg []*model.Message
	var batchGroupMsg []*model.GroupMessage

	timer := time.NewTimer(batchInterval)
	defer timer.Stop()

	for {
		select {
		case msg, ok := <-MessageChannel:
			if !ok {
				// Channel已关闭，处理剩余消息
				if len(batchSingleMsg) > 0 {
					if err := new(model.Message).BatchCreate(global.DB, batchSingleMsg); err != nil {
						global.Log.Error(err)
					}
				}
				return
			}
			batchSingleMsg = append(batchSingleMsg, msg)
			if len(batchSingleMsg) >= 100 {
				if err := new(model.Message).BatchCreate(global.DB, batchSingleMsg); err != nil {
					global.Log.Error(err)
				}
				batchSingleMsg = nil // 清空批次
				timer.Reset(batchInterval)
			}
		case msg, ok := <-GroupMessageChannel:
			if !ok {
				// Channel已关闭，处理剩余消息
				if len(batchGroupMsg) > 0 {
					if err := new(model.GroupMessage).BatchCreate(global.DB, batchGroupMsg); err != nil {
						global.Log.Error(err)
					}
				}
				return
			}
			batchGroupMsg = append(batchGroupMsg, msg)
			if len(batchGroupMsg) >= 100 {
				if err := new(model.GroupMessage).BatchCreate(global.DB, batchGroupMsg); err != nil {
					global.Log.Error(err)
				}
				batchGroupMsg = nil // 清空批次
				timer.Reset(batchInterval)
			}

		case <-timer.C:
			if len(batchSingleMsg) > 0 {
				if err := new(model.Message).BatchCreate(global.DB, batchSingleMsg); err != nil {
					global.Log.Error(err)
				}
				batchSingleMsg = nil // 清空批次
			}
			if len(batchGroupMsg) > 0 {
				if err := new(model.GroupMessage).BatchCreate(global.DB, batchGroupMsg); err != nil {
					global.Log.Error(err)
				}
				batchGroupMsg = nil // 清空批次
			}

			timer.Reset(batchInterval)
		}
	}
}
