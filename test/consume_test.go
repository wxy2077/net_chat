package test

import (
	"fmt"
	"testing"
	"time"
)

type Message struct {
	ID      int
	Content string
	Sender  string
	Time    int64
}

var (
	messageChannel = make(chan Message, 100)
	batchInterval  = 3 * time.Second // 定时器间隔
)

func Test_consume(t *testing.T) {
	// 启动消息生成器goroutine
	go produceMessages()

	// 启动消息消费者goroutine
	go consumeMessages()

	time.Sleep(20 * time.Second)
}

func produceMessages() {
	for i := 0; i < 500; i++ { // 模拟产生500条消息
		msg := Message{
			ID:      i,
			Content: fmt.Sprintf("Message %d", i),
			Sender:  "User1",
			Time:    time.Now().Unix(),
		}
		messageChannel <- msg
	}
	close(messageChannel)
}

func consumeMessages() {
	var batch []Message
	timer := time.NewTimer(batchInterval)
	defer timer.Stop()

	for {
		select {
		case msg, ok := <-messageChannel:
			if !ok {
				// Channel已关闭，处理剩余消息
				if len(batch) > 0 {
					saveToDatabase(batch)
				}
				return
			}
			batch = append(batch, msg)
			if len(batch) >= 100 {
				saveToDatabase(batch)
				batch = nil // 清空批次
				timer.Reset(batchInterval)
			}
		case <-timer.C:
			if len(batch) > 0 {
				saveToDatabase(batch)
				batch = nil // 清空批次
			}
			timer.Reset(batchInterval)
		}
	}
}

func saveToDatabase(messages []Message) {
	fmt.Printf("Saving %d messages to database\n", len(messages))
	// 模拟数据库保存
	time.Sleep(1 * time.Second)
}
