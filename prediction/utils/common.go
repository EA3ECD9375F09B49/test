package utils

import (
	"crypto/md5"
	"crypto/sha1"
	"encoding/hex"
	"fmt"
	"io/ioutil"
	"math/rand"
	"net/http"

	//"net/http"
	"net/url"
	"strconv"
	"strings"
	"time"

	jsoniter "github.com/json-iterator/go"
)

var Cjson = jsoniter.ConfigCompatibleWithStandardLibrary

type Config struct {
	ConnectTimeout   time.Duration
	ReadWriteTimeout time.Duration
}

func PostForm(path string, postData url.Values) ([]byte, error) {
	var s []byte
	resp, err := HttpClient.PostForm(path, postData)

	if err != nil {
		return s, err
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return s, err
	}
	return body, nil
}

func POSTJson(path string, postData map[string]interface{}, header map[string]string) ([]byte, error) {

	jsonStr, err := Cjson.Marshal(postData)
	if err != nil {
		return []byte(""), err
	}

	payload := strings.NewReader(string(jsonStr))
	req, _ := http.NewRequest("POST", path, payload)
	req.Header.Add("content-type", "application/json")

	for key, value := range header {
		req.Header.Add(key, value)
	}

	resp, err := HttpClient.Do(req)
	if err != nil {
		return []byte(""), err
	}
	defer resp.Body.Close()

	respBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return []byte(""), err
	}

	return respBytes, nil
}

func RandInt64(min, max int64) int64 {
	if min >= max || min == 0 || max == 0 {
		return max
	}
	return rand.Int63n(max-min) + min
}
func Md5V(str string) string {
	h := md5.New()
	h.Write([]byte(str))
	return hex.EncodeToString(h.Sum(nil))
}

func GetSha1(str string) string {
	h := sha1.New()
	h.Write([]byte(str))
	str = fmt.Sprintf("%x", h.Sum(nil))
	return str
}

func GetInt(v interface{}) (int, error) {
	switch v := v.(type) {
	case float64:
		return int(v), nil
	case string:
		c, err := strconv.Atoi(v)
		if err != nil {
			return 0, err
		}
		return c, nil
	case int:
		return v, nil
	default:
		return 0, fmt.Errorf("conversion to int from %T not supported", v)
	}
}

func FilterListInfoStr(str string) (newStr string) {
	if len(str) < 1 {
		return
	}
	arr := strings.Split(str, "|")
	if len(arr) < 1 {
		return
	}
	newStr = arr[0]
	return
}

func InArray(val string, array []string) (exists bool) {
	exists = false
	for _, v := range array {
		if val == v {
			exists = true
			return
		}
	}
	return
}

//type IdsPoint []string
func InArrayPointer(val string, idsPoint *[]string) (exists bool) {
	exists = false
	for _, v := range *idsPoint {
		if val == v {
			exists = true
			return
		}
	}
	return
}

func GetSuffix(suffix, seperate string) (respSuffix string) {
	respSuffix = suffix
	if len(suffix) < 1 {
		return
	}
	respSuffix = seperate + "" + strings.ToUpper(suffix)
	return
}

func GetValidKeyname(keyName, suffix, seperate string) (respKeyName string) {
	respKeyName = keyName + "" + GetSuffix(suffix, seperate)
	return
}
