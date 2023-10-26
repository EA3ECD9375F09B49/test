package main

import (
	"bufio"
	"crypto/aes"
	"crypto/cipher"
	"encoding/hex"
	"errors"
	"fmt"
	"os"
	"strings"
)

func main() {

	var AesKey = map[string]string{
		"fat": "A85F#H80q2u8Qt0k",
		"uat": "kD%voCkU&RP(*H)1zJE2oTdFR)#Wes7X",
	}

	var Iv = map[string]string{
		"fat": "gv0-_R6E@ncCQAUj",
		"uat": "YBh2x%#rH!prkYj5",
	}
	//encodingAesKey := "u9PBrZ2nudh8sD0GMQ9cWYznZOLe7CSd"
	//data := "CT+QSxoHVm68PamNsuFv4BFDMFdfjYkSvc+qyhnYbHzn1LnEBcxuON0aH0L5dMaD5fd07fa6a9MRPOTaY/4rSX0a64TCGiDC2u+eohlfiVm6oTxhOiqQq0dYZb/INt8luXveUxjjaPkwXhzrNg6mpw=="
	reader := bufio.NewReader(os.Stdin)
	fmt.Printf("请输入环境 测试输入 fat, 预发请输入 uat \n")
	var env string
	env, err := reader.ReadString('\n')

	env = strings.Trim(env, "\f\t\r\n ")
	// convert CRLF to LF
	env = strings.Replace(env, "\n", "", -1)
	if err != nil {
		return
	}
	encodingAesKey := AesKey[env]
	ivKey := Iv[env]

	if len(encodingAesKey) <= 0 || len(ivKey) <= 0 {
		fmt.Printf("环境输入错误 \n")
		return
	}

	var data string
	for {
		fmt.Printf("继续请输入 y, 退出请输入 n 后回车 \n")

		var cnt string
		cnt, err = reader.ReadString('\n')
		cnt = strings.Trim(cnt, "\f\t\r\n ")
		// convert CRLF to LF
		cnt = strings.Replace(cnt, "\n", "", -1)
		if err != nil {
			return
		}
		if cnt != "y" {
			return
		}

		fmt.Printf("请输入需要解密的字符\n")
		data, err = reader.ReadString('\n')
		data = strings.Trim(data, "\f\t\r\n ")
		data = strings.Replace(data, "\n", "", -1)
		if err != nil {
			return
		}
		decString := GetRealString(encodingAesKey, data, ivKey)
		fmt.Printf("解密后的值为 : %v \n", decString)
	}

}

func GetRealString(encodingAesKey string, data string, ivKey string) string {
	dataTmp, err := hex.DecodeString(data)
	if err != nil {
		panic(err)
	}
	ivByte := make([]byte, 0)
	var aeskey []byte
	ivByte = []byte(ivKey)
	aeskey = []byte(encodingAesKey)

	rs, err := AesCBCPk7Decrypt(dataTmp, aeskey, ivByte)
	if err != nil {
		panic(err)
	}

	return rs
}

// Aes cbc 解密, pkcs7 填充
func AesCBCPk7Decrypt(encryption, key []byte, iv []byte) (string, error) {
	if len(encryption) < 1 {
		return "", errors.New("encryption is empty")
	}
	if len(key) < 1 {
		return "", errors.New("key is empty")
	}
	if len(iv) < 1 {
		return "", errors.New("iv is empty")
	}

	block, err := aes.NewCipher(key)
	if err != nil {
		return "", err
	}

	// 加入判断条件防止 panic
	blockSize := block.BlockSize()
	if len(key) < blockSize {
		return "", errors.New("key too short")
	}
	if len(encryption)%blockSize != 0 {
		return "", errors.New("crypto/cipher: input not full blocks")
	}

	blockMode := cipher.NewCBCDecrypter(block, iv)
	origData := make([]byte, len(encryption))
	blockMode.CryptBlocks(origData, encryption)
	origData = PKCS7UnPadding(origData, blockSize)
	return string(origData), nil
}

// PKCS7 填充
func PKCS7UnPadding(plantText []byte, blockSize int) []byte {
	length := len(plantText)
	unpadding := int(plantText[length-1])
	return plantText[:(length - unpadding)]
}
