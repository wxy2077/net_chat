package global

const (
	HeatBeat = "heartbeat"
	PONG     = "pong"

	// 消息类型，单聊或者群聊
	MessageTypeUser  = 1
	MessageTypeGroup = 2

	// 消息内容类型
	Text        = 1
	File        = 2
	Image       = 3
	Audio       = 4
	Video       = 5
	AudioOnline = 6
	VideoOnline = 7

	// 消息队列类型
	GoChannel = "go_channel"
	Kafka     = "kafka"
)
