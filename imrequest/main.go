package main

import (
	"bytes"
	"crypto"
	"crypto/md5"
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/hex"
	"encoding/pem"
	"fmt"
	jsoniter "github.com/json-iterator/go"
	"imrequest/utils"
	"net/url"
	"regexp"
	"sort"
	"strings"
)

// 存数据时
var ObMap = map[string]string{
	"name":         "OBESPORT",
	"api":          "https://djbob.merchantapi10.com",
	"merchantCode": "26217611501329530",
	"secret_key":   "bda079f2c556a2a282c13b96a2904e1b",
	"web_hrl":      "https://search.gpthelp.com",
	"h5_url":       "https://item.gpthelp.com",
	"pull":         "https://djpull.i9js01.com",
}

var Cjson = jsoniter.ConfigCompatibleWithStandardLibrary

const privPemFmt = `-----BEGIN PRIVATE KEY-----
%s
-----END PRIVATE KEY-----`

// ob main
// https://djpull.i9js01.com/pull/order/queryScroll?agency=false&end_time=&last_order_id=0&merchant=38425828292093166&page_size=5000&sign=g8c58b70e28PKf265b6c6KXbbea8a69d524c9cVk&start_time=
func main() {

	DescFormula := "^[^\\s][\u4e00-\u9fa5\u3001\\d\\sA-Za-z（）/\\s-]{0,48}[^\\s](,[^\\s][\u4e00-\u9fa5\u3001\\dA-Za-z（）/\\s-]{0,48}[^\\s])*$"
	aa := "aaa（）bbb,cc（dd/ef-）"
	r := regexp.MustCompile(DescFormula)
	var bb string
	if r.MatchString(aa) { //"a,b,c,d,e"
		bb = "success"
	}
	fmt.Println(bb)

}

type Resource struct {
	Name string
}

func dummy(initializeResource bool) {
	var resource *Resource = nil
	init := false

	if initializeResource {
		resource = new(Resource)
		init = true
	}

	r := resource
	_ = r.Name
	fmt.Print(init)
}

func RsaSign(data []byte) ([]byte, error) {
	h := md5.New()
	h.Write(data)
	hashed := h.Sum(nil)
	s := []byte("")
	//获取私钥
	Md5Keys := fmt.Sprintf(privPemFmt, "MIIBUwIBADANBgkqhkiG9w0BAQEFAASCAT0wggE5AgEAAkEAnWsqywtNwmUbkKF+svHKgjKZ7Yktsqy2UR8JbsvEG8El/JwZMF+mM9yUIf7D4y983UOCRnSaIZyNnpQs/bCU6QIDAQABAkBlbHPg6IKMOjqdX//S6YhxhIq6icTgtvisoZOhSDYtLqxPvSEXcuB9T3C2XR6Fdfy/Pe8zsaMgtq0GCECSWitFAiEA2xdQ7GETWUtuBC5+L2EQjbrzyGLdpqP16/hmmGPYdPsCIQC38B57PXuaOaSZe6eYO73CSt+LUTF0oB5u4y4LLDgQawIgBIJ9WvBAmrTvxcxDMqx3z8MKY5SNZXG4jSvmyLo9aWMCIEL/PCsJfv7y5ghdqPLjN8AQQ5JcNZZUSCF9sCSzq4wXAiA6FTzgodPS5NAU3XP3Nj61xr7yOsb9UQ7uzvVsA6zMQw==")
	block, _ := pem.Decode([]byte(Md5Keys))
	//解析PKCS1格式的私钥
	priv, err := x509.ParsePKCS8PrivateKey(block.Bytes)
	if err != nil {
		return s, err
	}
	return rsa.SignPKCS1v15(rand.Reader, priv.(*rsa.PrivateKey), crypto.MD5, hashed)
}

func getQuerySignString(paramsMap map[string]string, signKey string) string {
	queryParams := make([]string, len(paramsMap))

	for k, v := range paramsMap {
		queryParams = append(queryParams, k+v)
	}
	sort.Strings(queryParams)

	//查询参数加秘钥组成明文
	clearText := strings.Join(queryParams, "") + signKey

	return utils.MD5Hash(clearText)
}

// IM MAIN
//func main() {
//	usLoc := time.FixedZone("US", -4*3600)
//	now := time.Now().In(usLoc).Format("2006-01-02 15:04:05.000")
//	now = "2022-11-21 07:06:13.000"
//	aa := utils.AesECBEncrypt(now, md5raw())
//	fmt.Printf(aa)
//}

func md5raw() string {
	has := md5.Sum([]byte("cb2834934d87a132"))
	key := fmt.Sprintf("%x", has)
	lens := len(key) / 2
	md5raw := ""
	for i := 0; i < lens; i++ {
		start := i * 2
		end := (i * 2) + 2
		substring := utils.Substring(key, start, end)
		hexByte, _ := hex.DecodeString(substring)
		md5raw = md5raw + string(hexByte)
	}

	return md5raw
}

/**
 * @Description:  签名算法
 * @receiver platform Platform
 * @param params 业务参数
 * @return string 最终拼接的get参数
 */
func generateSign(params map[string]string) string {
	reqMap := make(url.Values, len(params)+1)
	params["key"] = ObMap["secret_key"]
	keys := []string{}
	for key := range params {
		keys = append(keys, key)
	}
	sort.Sort(sort.StringSlice(keys))
	for key := range keys {
		reqMap.Add(keys[key], params[keys[key]])
	}
	reqStr := reqMap.Encode()
	hash := utils.MD5Hash(reqStr)
	temp1 := utils.GetRandomString(2)
	temp2 := utils.GetRandomString(2)
	temp3 := utils.GetRandomString(2)
	temp4 := utils.GetRandomString(2)
	var buffer bytes.Buffer
	buffer.WriteString(temp1)
	buffer.WriteString(hash[0:9])
	buffer.WriteString(temp2)
	buffer.WriteString(hash[9:17])
	buffer.WriteString(temp3)
	buffer.WriteString(hash[17:])
	buffer.WriteString(temp4)
	resultParams := make(url.Values, len(params)+1)
	resultParams.Add("sign", buffer.String())
	for key := range params {
		if key != "key" { //只拼接业务参数
			resultParams.Add(key, params[key])
		}
	}
	return resultParams.Encode()
}
