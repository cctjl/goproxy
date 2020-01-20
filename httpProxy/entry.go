package httpProxy

/**
其中代理的数据处理部分调用相同函数实现
*/

import (
	"awesomeProject1/config"
	"log"
	"net/http"
	"strconv"
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
		IdleTimeout:       time.Duration(option.Timeout) * time.Second,
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
	host := r.Host
	currConfig := config.HttpConf{}
	for _, v := range config.Config.Http {
		if v.Domain == host {
			currConfig = v
			break
		}
	}
	if currConfig.Domain == "" {
		log.Println("未找到对应域名配置：", host)
		http.Error(w, "未配置处理服务器地址", http.StatusNotFound)
		return
	}
	url := GetUrl(r, "")
	log.Println("收到请求：", url)

	if currConfig.Auth.TokenKey != "" {
		authResult := DoAuth(r)
		log.Println("认证未通过：", url)
		http.Error(w, "认证失败："+authResult.Msg, http.StatusForbidden)
		return
	}
	DoProxy(&w, r, &currConfig)
}
