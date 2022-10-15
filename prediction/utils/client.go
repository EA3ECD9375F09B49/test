package utils

import (
	"bytes"
	"errors"
	"io/ioutil"
	"net"
	"net/http"
	"time"
)

func NewHttpClientPoolNoMaxPerHost(idelConnCount int) *http.Client {
	defaultTransPort := &http.Transport{
		Proxy: http.ProxyFromEnvironment,
		DialContext: (&net.Dialer{
			Timeout:   30 * time.Second,
			KeepAlive: 15 * time.Second,
		}).DialContext,
		TLSHandshakeTimeout:   10 * time.Second,
		MaxIdleConns:          idelConnCount,
		MaxIdleConnsPerHost:   idelConnCount, //控制对三方的keepalive负载，超过后net/http会自动新建，可能会产生很多被三方丢弃的established
		IdleConnTimeout:       90 * time.Second,
		ExpectContinueTimeout: 1 * time.Second,
	}
	return &http.Client{
		Transport: defaultTransPort,
		Timeout:   30 * time.Second,
	}
}
func GetObSportHttpClient() *http.Client {
	return &http.Client{
		Timeout: 30 * time.Second,
	}
}

func PostV2(path string, postData []byte, header map[string]string, httpClient *http.Client) ([]byte, error) {
	return HttpUtilV2(path, "POST", postData, header, httpClient)
}
func GetV2(path string, postData []byte, header map[string]string, httpClient *http.Client) ([]byte, error) {
	return HttpUtilV2(path, "GET", postData, header, httpClient)
}
func HttpUtilV2(path, method string, postData []byte, header map[string]string, httpClient *http.Client) ([]byte, error) {
	var s []byte
	if httpClient == nil {
		return s, errors.New("尚未指定http.Client")
	}
	payload := bytes.NewReader(postData)
	req, err := http.NewRequest(method, path, payload)
	if err != nil {
		return s, err
	}
	for key, value := range header {
		req.Header.Add(key, value)
	}
	res, err := httpClient.Do(req)
	if err != nil {
		return s, err
	}
	if res != nil {
		defer res.Body.Close()
	}
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return s, err
	}
	return body, nil
}
