package httpProxy

// 通过 chan 实现限时

import (
	"awesomeProject1/config"
	"io/ioutil"
	"log"
	"math/rand"
	"net/http"
	"strconv"
	"time"
)

func DoAuth(r *http.Request) BaseAuthResult {
	result := BaseAuthResult{true, "认证通过", nil}
	return result
}

/**
将请求转发到对应地址，先将 header 复制，然后再加上 body 按原请求方法发过去，考虑到配置的超时时间等问题
需要从 httpConf 中的 host 中选一个转发，todo，先随机选一个
*/
func DoProxy(w *http.ResponseWriter, r *http.Request, httpConf *config.HttpConf) {
	dstAddress := httpConf.Hosts[rand.Intn(len(httpConf.Hosts))]
	dstDomain := dstAddress.Host
	if dstAddress.Port != 80 && dstAddress.Port != 0 {
		dstDomain = dstAddress.Host + ":" + strconv.FormatUint(uint64(dstAddress.Port), 10)
	}
	url := GetUrl(r, dstDomain)

	req, err := http.NewRequest(r.Method, url, r.Body)
	log.Println("请求方式： ", r.Method, "，url：", r.URL.String(), "，转发到url：", url)
	if err != nil {
		log.Println(r.RequestURI, " 构造请求出错：", err)
		http.Error(*w, "转发请求失败", http.StatusInternalServerError)
		return
	}
	CopyHeaders(&req.Header, &r.Header, true)
	client := &http.Client{Timeout: time.Duration(httpConf.Timeout) * time.Second}
	resp, err := client.Do(req)
	if err != nil {
		log.Println(r.RequestURI, " 后端请求出错：", err)
		http.Error(*w, "转发请求超时", http.StatusRequestTimeout)
		return
	}

	setResponseHeaders(*w, resp)
	defer resp.Body.Close()
	respBody, _ := ioutil.ReadAll(resp.Body)
	(*w).Write(respBody)
}

func setResponseHeaders(w http.ResponseWriter, response *http.Response) {
	for k, v := range response.Header {
		w.Header()[k] = v
	}
}

func CopyHeaders(dst *http.Header, src *http.Header, keepDestHeaders bool) {
	if !keepDestHeaders {
		for k := range *dst {
			dst.Del(k)
		}
	}
	for k, vs := range *src {
		for _, v := range vs {
			dst.Add(k, v)
		}
	}
}
