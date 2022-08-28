package main

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"crypto/md5"
	"encoding/base64"
	"encoding/hex"
	"errors"
	"fmt"
)

func main() {
	//encodingAesKey := "u9PBrZ2nudh8sD0GMQ9cWYznZOLe7CSd"
	//data := 'CT+QSxoHVm68PamNsuFv4BFDMFdfjYkSvc+qyhnYbHzn1LnEBcxuON0aH0L5dMaD5fd07fa6a9MRPOTaY/4rSX0a64TCGiDC2u+eohlfiVm6oTxhOiqQq0dYZb/INt8luXveUxjjaPkwXhzrNg6mpw=='
	//md5Str := '82bd3492908cad01fd06c5d79bb86a94'
	//str := "testtest:WgAewQMs3idC4aGnjGtBXw==@tcp(w_db.service.consul:3306)/tybdata?loc=Local&parseTime=true"
	//encString := GetEncString(encodingAesKey, str)
	//fmt.Printf("encString is : %v", encString)

	fmt.Printf("please enter encodingAesKey\n")
	var encodingAesKey string
	scan, err := fmt.Scan(&encodingAesKey)
	if err != nil {
		return
	}
	fmt.Printf("The encodingAesKey is %s and scan status is %d \n",
		encodingAesKey, scan)

	var str string
	for {
		fmt.Printf("please enter string to encrypt\n")
		scan1, err := fmt.Scan(&str)
		if err != nil {
			return
		}
		fmt.Printf("The word to encString is %s and scan status is %d \n",
			str, scan1)
		encString := GetEncString(encodingAesKey, str)
		fmt.Printf("encString is : %v \n", encString)
	}

}

func GetEncString(encodingAesKey string, originalString string) string {
	var md5Str = Md5EncodeToString(encodingAesKey)
	fmt.Printf("original String is : %v \n", originalString)
	fmt.Printf("aesKey is : %v \n", encodingAesKey)
	fmt.Printf("md5 String is : %v \n", md5Str)
	rs, err := AesCBCPk7Encrypt([]byte(originalString), getAesKey(md5Str), getIv(md5Str))
	if err != nil {
		panic(err)
	}
	str := string(rs)
	base64EncodedString := base64.StdEncoding.EncodeToString([]byte(str))
	if err != nil {
		panic(err)
	}
	return base64EncodedString
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

// Aes cbc 加密, pkcs7 填充
func AesCBCPk7Encrypt(origData, key []byte, iv []byte) ([]byte, error) {
	if len(origData) < 1 {
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

	blockSize := block.BlockSize()
	origData = PKCS7Padding(origData, blockSize)
	blockMode := cipher.NewCBCEncrypter(block, iv[:blockSize])
	crypted := make([]byte, len(origData))
	blockMode.CryptBlocks(crypted, origData)
	return crypted, nil
}

// PKCS7 填充
func PKCS7Padding(ciphertext []byte, blockSize int) []byte {
	padding := blockSize - len(ciphertext)%blockSize
	padtext := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(ciphertext, padtext...)
}
