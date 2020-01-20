package config

const (
	QueueRedis = "redis"
)

type authOption struct {
	TokenKey    string
	TokenMapKey string
}

type NetAddress struct {
	Host string
	Port uint
}

type HttpConf struct {
	LocalHost NetAddress
	Ishttps   bool
	Domain    string
	Timeout   uint
	Auth      authOption
	Hosts     []NetAddress
}

//type WebsocketConf struct {
//	Timeout   uint
//	QueueName string
//	UpQueue   string
//	DownQueue string
//	LocalHost NetAddress
//	Hosts     []NetAddress
//}

type RedisConf struct {
	Host      string
	Port      uint
	Db        uint
	QueueName string
}

type AllConf struct {
	Http []HttpConf
	//Websocket []WebsocketConf
}
