package main

import (
	"bufio"
	"crypto/aes"
	"crypto/cipher"
	"crypto/md5"
	"encoding/base64"
	"encoding/hex"
	"errors"
	"fmt"
	"os"
	"strings"
)

func main() {
	//encodingAesKey := "u9PBrZ2nudh8sD0GMQ9cWYznZOLe7CSd"
	//data := "CT+QSxoHVm68PamNsuFv4BFDMFdfjYkSvc+qyhnYbHzn1LnEBcxuON0aH0L5dMaD5fd07fa6a9MRPOTaY/4rSX0a64TCGiDC2u+eohlfiVm6oTxhOiqQq0dYZb/INt8luXveUxjjaPkwXhzrNg6mpw=="
	json := "{\"aa\": \"bb\"}"
	fmt.Println(json)

	fmt.Printf("please enter decodingAesKey\n")

	reader := bufio.NewReader(os.Stdin)
	encodingAesKey, err := reader.ReadString('\n')
	if err != nil {
		return
	}
	encodingAesKey = strings.Trim(encodingAesKey, "\f\t\r\n ")
	// convert CRLF to LF
	encodingAesKey = strings.Replace(encodingAesKey, "\n", "", -1)
	fmt.Printf("The decodingAesKey is %s \n",
		encodingAesKey)

	fmt.Printf("please enter ivKey\n")
	var ivKey string
	ivKey, err = reader.ReadString('\n')
	if err != nil {
		return
	}
	ivKey = strings.Trim(ivKey, "\f\t\r\n ")
	// convert CRLF to LF
	ivKey = strings.Replace(ivKey, "\n", "", -1)

	fmt.Printf("The ivKey is %s \n",
		ivKey)

	var data string
	for {
		fmt.Printf("please enter string to decrypt\n")
		data, err = reader.ReadString('\n')
		if err != nil {
			return
		}
		data = strings.Trim(data, "\f\t\r\n ")
		// convert CRLF to LF
		data = strings.Replace(data, "\n", "", -1)
		fmt.Printf("The word to decrypt is %s \n",
			data)
		decString := GetRealString(encodingAesKey, data, ivKey)
		fmt.Printf("decString is : %v \n", decString)
	}

}

func GetRealString(encodingAesKey string, data string, ivKey string) string {
	dataTmp, err := base64.StdEncoding.DecodeString(data)
	if err != nil {
		panic(err)
	}
	var md5Str = Md5EncodeToString(encodingAesKey)
	fmt.Printf("encrypted String is : %v \n", data)
	fmt.Printf("base64decoded String is : %v \n", dataTmp)
	fmt.Printf("aesKey is : %v \n", encodingAesKey)
	fmt.Printf("md5 String is : %v \n", md5Str)

	ivByte := make([]byte, 0)
	if ivKey != "no" {
		fmt.Printf("The ivKey is %s \n",
			ivKey)
		ivByte = []byte(ivKey)
	} else {
		fmt.Println("The ivKey is empty and we will go by our md5")
		ivByte = getIv(md5Str)
	}

	rs, err := AesCBCPk7Decrypt(dataTmp, getAesKey(md5Str), ivByte)
	if err != nil {
		panic(err)
	}

	return string(rs)
}

// Aes cbc 解密, pkcs7 填充
func AesCBCPk7Decrypt(crypted, key []byte, iv []byte) ([]byte, error) {
	if len(crypted) < 1 {
		return []byte(""), errors.New("crypted is empty")
	}
	if len(key) < 1 {
		return []byte(""), errors.New("key is empty")
	}
	if len(iv) < 1 {
		return []byte(""), errors.New("iv is empty")
	}

	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}

	// 加入判断条件防止 panic
	blockSize := block.BlockSize()
	if len(key) < blockSize {
		return nil, errors.New("key too short")
	}
	if len(crypted)%blockSize != 0 {
		return nil, errors.New("crypto/cipher: input not full blocks")
	}

	blockMode := cipher.NewCBCDecrypter(block, iv)
	origData := make([]byte, len(crypted))
	blockMode.CryptBlocks(origData, crypted)
	origData = PKCS7UnPadding(origData, blockSize)
	return origData, nil
}

func Md5EncodeToString(s string) string {
	hexCode := md5.Sum([]byte(s))
	return hex.EncodeToString(hexCode[:])
}

func getAesKey(key string) []byte {
	if len(key) != 32 {
		panic("error secret key")
	}
	return []byte(key[2:7] + key[11:15] + key[18:25])
}

func getIv(key string) []byte {
	if len(key) != 32 {
		panic("error secret key")
	}
	return []byte(key[4:9] + key[16:23] + key[25:29])
}

//// Aes cbc 解密, pkcs7 填充
//func AesCBCPk7Decrypt(crypted, key []byte, iv []byte) ([]byte, error) {
//	if len(crypted) < 1 {
//		return []byte(""), errors.New("crypted is empty")
//	}
//	if len(key) < 1 {
//		return []byte(""), errors.New("key is empty")
//	}
//	if len(iv) < 1 {
//		return []byte(""), errors.New("iv is empty")
//	}
//
//	block, err := aes.NewCipher(key)
//	if err != nil {
//		return nil, err
//	}
//
//	// 加入判断条件防止 panic
//	blockSize := block.BlockSize()
//	if len(key) < blockSize {
//		return nil, errors.New("key too short")
//	}
//	if len(crypted)%blockSize != 0 {
//		return nil, errors.New("crypto/cipher: input not full blocks")
//	}
//
//	blockMode := cipher.NewCBCDecrypter(block, iv)
//	origData := make([]byte, len(crypted))
//	blockMode.CryptBlocks(origData, crypted)
//	origData = PKCS7UnPadding(origData, blockSize)
//	return origData, nil
//}

// PKCS7 填充
func PKCS7UnPadding(plantText []byte, blockSize int) []byte {
	length := len(plantText)
	unpadding := int(plantText[length-1])
	return plantText[:(length - unpadding)]
}
