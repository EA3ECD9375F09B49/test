package main

import (
	"crypto"
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/base64"
	"encoding/pem"
	"fmt"
	"net/url"
)

func main() {

	block, _ := pem.Decode([]byte(`
-----BEGIN PRIVATE KEY-----
  MIICeQIBADANBgkqhkiG9w0BAQEFAASCAmMwggJfAgEAAoGBALaiUegzB3VChAHOkBomXH//JWKdoUo3X9rujbdnC0NN0gmL50rMVF2tm7+7f7NTZMIaIOnqH6DlI9xKLI1xhBS4OPvmB0aX5UdrlhrDlkNtLGfzQ01ztyJfIabS4aGyyr3VDIbQc9DgNXQQJRw0OlFfdlKuYOhx088XXNS3J6U5AgMBAAECgYEAlCo7qNU9R7QexBiAb7aPdIfaHJclMsCZ01OVRSUjzgZkT7pgeh4qk5U0tee3QhR1ucLY0OxPnIhI/35thpKBQ65fccMAsekZvgbZiBtEiLAr3TGZmKZtWgatxxiqIUmXLoW/uIgPCwZYVgoDoOiHOTtrwI/zIwp0NtspUOOSdMECQQDmmljdigFZA+NkfxVs1pMmmkQ+WJkIzgzbvaTfRc11pgOsmOIrWnv7oQ1J8+2yy68uzsu17mjTILTNRfTiwoHjAkEAyr+IIoRB4sNPr8S/mHqkiwlUyurPOh2zCckrmTx77zdolrPm3GBYuAj6hfjkBw4sJ2LAOyUjG58ksVIPJYA3MwJBALC8UqdYAbhrVnfLPmxv9896JBt0Y7Vv1kMDkbzdDp7AYciCU7TOYH91621mWiLSIK1LKK2CzywgwPEiUJEKaekCQQC6nWSOcLG8KQ2VuUCo9mmxV9tmEo6+7Us3/KRWnSdt3dA0tk4OFdhTPGBrI9Wu8MqPTgOl7N7Ns8OwOtmaKNOlAkEA3wqudEhKnRghZqHEREtTi4zvigV4O5g8K53EurDfVjjgng02mk6FMoQyVlgX+b8VJYfBpNwUdrsHj5BFjPbvYw==
-----END PRIVATE KEY-----`))

	context := "appkey=tml0cj&platform=ios" //需要签名的文本

	if block == nil {
		panic("私钥错误")
		return
	}

	private, err := x509.ParsePKCS8PrivateKey(block.Bytes) //之前看java demo中使用的是pkcs8
	if err != nil {
		panic("PrivateKey error")
		return
	}
	h := crypto.Hash.New(crypto.SHA1) //进行SHA1的散列
	h.Write([]byte(context))
	hashed := h.Sum(nil)

	// 进行rsa加密签名
	signedData, err := rsa.SignPKCS1v15(rand.Reader, private.(*rsa.PrivateKey), crypto.SHA1, hashed)

	data := base64.StdEncoding.EncodeToString(signedData)
	encodedStr := url.QueryEscape(data)
	fmt.Println(encodedStr)

}
