package queue

// 将与消息的交互封装到此包中，将配置项设置为全局变量，只初始化一次

type QueueUtil interface {
	Init()
	ProduceMessage()
	ConsumeMessage()
}
