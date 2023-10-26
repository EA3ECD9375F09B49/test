package main

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/sha1"
	"encoding/base64"
	"errors"
	"fmt"
)

func main() {
	str := "J6tUXG5nPP9JcbBAdf+eqQ=="
	realkey, _ := getRealKey()
	iv := "SNeS5Bj5safieQLL"
	key, _ := AesSha1prng([]byte(realkey), 128)
	fmt.Println(string(key))
	pt, _ := AesCBCPk5DecryptBase64(str, string(key), []byte(iv))
	fmt.Println(pt)
}

// 获取真实密钥
func getRealKey() (string, error) {
	var encryptKey = "a01ogy0sTVqyNlOJHuKsw1s"
	if encryptKey == "" || len(encryptKey) < 7 {
		return "", errors.New("密钥长度太小")
	}
	// 去掉前三后四掩码
	realKey := encryptKey[3 : len(encryptKey)-4]
	if len(realKey) > 32 {
		return "", errors.New("密钥长度过长")
	}
	return realKey, nil
}

// SHA1PRNG 处理
func AesSha1prng(keyBytes []byte, encryptLength int) ([]byte, error) {
	aa := Sha1(keyBytes)
	hashs := Sha1(aa)
	maxLen := len(hashs)
	realLen := encryptLength / 8
	if realLen > maxLen {
		return nil, errors.New("invalid length!")
	}

	return hashs[0:realLen], nil
}

func Sha1(data []byte) []byte {
	h := sha1.New()
	h.Write(data)
	return h.Sum(nil)
}

// Aes cbc 解密, pkcs5 填充, base64编码
func AesCBCPk5DecryptBase64(encrypt string, key string, iv []byte) (string, error) {
	encryptBytes, err := base64.StdEncoding.DecodeString(encrypt)
	if err != nil {
		return "", err
	}
	str, err := AesCBCPk7Decrypt(encryptBytes, []byte(key), iv)
	if err != nil {
		return "", err
	}

	return string(str), nil
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

// PKCS7 填充
func PKCS7UnPadding(plantText []byte, blockSize int) []byte {
	length := len(plantText)
	unpadding := int(plantText[length-1])
	return plantText[:(length - unpadding)]
}
