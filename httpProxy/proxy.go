package httpProxy

// 通过 chan 实现限时

import (
	"awesomeProject1/config"
	"net/http"
)

/**
代理逻辑处理，将请求根据配置找到后端服务器地址
*/
func FindRoute(request http.Request, conf config.HttpConf) config.NetAddress {

	return config.NetAddress{}
}

/**
将请求转发到对应地址
*/
func DoProxy(r http.Request, address config.NetAddress) {

}
