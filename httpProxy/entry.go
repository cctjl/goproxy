package httpProxy

/**
其中代理的数据处理部分调用相同函数实现
*/

import (
	"awesomeProject1/config"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"
	"time"
)

func StartHttpProxy(option config.HttpConf) {
	// 监听对应地址，设置过期时间为配置时间值
	strLocalAddress := string(option.LocalHost.Host) + ":" + strconv.Itoa(int(option.LocalHost.Port))
	mux := http.NewServeMux()
	mux.HandleFunc("/", handler)
	server := &http.Server{
		Addr:              strLocalAddress,
		Handler:           mux,
		TLSConfig:         nil,
		ReadTimeout:       time.Second * 3,
		ReadHeaderTimeout: time.Second * 1,
		WriteTimeout:      time.Second * 3,
		IdleTimeout:       time.Duration(option.Timeout),
		MaxHeaderBytes:    8072,
		TLSNextProto:      nil,
		ConnState:         nil,
		ErrorLog:          nil,
		BaseContext:       nil,
		ConnContext:       nil,
	}
	log.Println("starting: " + strLocalAddress)
	log.Fatal(server.ListenAndServe())
}

func StartHttpsProxy(option config.HttpConf) {

}

func handler(w http.ResponseWriter, r *http.Request) {
	_, e := fmt.Fprintln(w, "hello world")

	if e != nil {
		log.Println("写入消息失败: " + getUrl(r))
	}
}

func getUrl(r *http.Request) string {
	scheme := "http://"
	if r.TLS != nil {
		scheme = "https://"
	}
	url := strings.Join([]string{scheme, r.Host, r.RequestURI}, "")
	return url
}
