package utils

import (
	"fmt"
	"io/ioutil"
	"net"
	"net/http"
	"time"
)

const (
	RequestTimeout int = 30 // 请求以及连接的超时时间
)

var (
	HttpClient = &http.Client{
		Timeout: time.Duration(RequestTimeout) * time.Second,
	} // 不过滤掉证书检查的 http client
	client = &http.Client{
		Transport: &http.Transport{
			DialContext: (&net.Dialer{
				Timeout: 5 * time.Second,
			}).DialContext,
			DisableKeepAlives: true,
		},
		Timeout: 5 * time.Second,
	}
)

func GetUrl(url string) (result string) {
	resp, err := client.Get(url)
	result = ""
	if err != nil {
		_ = fmt.Sprintf("getUrl:url:%v err:%v", url, err)
		return
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		_ = fmt.Sprintf("getUrl:url:%v err:%v", url, err)
		return
	}
	result = string(body)
	return
}
