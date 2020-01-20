package main

import (
	"awesomeProject1/config"
	"awesomeProject1/httpProxy"
	"encoding/json"
	"flag"
	yaml "gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
	"os"
)

func initConf(confPath string) config.AllConf {
	log.Println("使用配置文件: ", confPath)
	byteConf, err := ioutil.ReadFile(confPath)
	if err != nil {
		log.Println("读取配置文件出错: ", err)
	}
	var confData config.AllConf
	err = yaml.Unmarshal(byteConf, &confData)
	if err != nil {
		log.Fatal("解析配置文件出错: ", err)
	}

	return confData
}

func IsFileExisted(f string) bool {
	fi, e := os.Stat(f)
	if e != nil {
		return false
	}
	return !fi.IsDir()
}

func startProxy(option config.AllConf) {
	// 启动 http 代理，在此之前先将配置设置为全局变量
	config.Config = option
	for i := 0; i < len(option.Http); i++ {
		httpConf := &option.Http[i]
		if httpConf.Ishttps {
			go httpProxy.StartHttpsProxy(*httpConf)
		} else {
			go httpProxy.StartHttpProxy(*httpConf)
		}
	}

	log.Println("所有任务已启动")
	// 主协程陷入永久休眠
	config.Wg.Add(1)
	config.Wg.Wait()
}

func main() {
	log.Println("Gateway 1.0" +
		"\n著作人：言西早" +
		"\n日期：2019.11.27")
	confPath := flag.String("c", "config.yaml", "配置文件路径")
	flag.Parse()
	if !IsFileExisted(*confPath) {
		log.Fatal("配置文件不存在: ", *confPath)
	}
	conf := initConf(*confPath)
	v, e := json.Marshal(conf)
	if e != nil {
		log.Fatal("将配置转为json失败: ", e)
	}
	log.Println("启动配置项：", string(v))
	startProxy(conf)
}
