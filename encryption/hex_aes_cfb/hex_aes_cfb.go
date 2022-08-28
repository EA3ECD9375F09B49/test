package main

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"encoding/hex"
	"errors"
	"fmt"
)

func main() {

}

func GetEncString(encodingAesKey string, originalString string, ivKey string) string {
	var aeskey []byte
	ivByte := make([]byte, 0)
	ivByte = []byte(ivKey)
	aeskey = []byte(encodingAesKey)

	rs, err := AESCFBEncrypt([]byte(originalString), aeskey, ivByte)
	if err != nil {
		panic(err)
	}
	str := string(rs)
	hexEncodedString := hex.EncodeToString([]byte(str))
	if err != nil {
		panic(err)
	}
	return hexEncodedString
}

// Aes cbc 加密, pkcs7 填充
func AESCFBEncrypt(origData, key []byte, iv []byte) ([]byte, error) {
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
	ciphertext := make([]byte, aes.BlockSize+len(origData))
	stream := cipher.NewCFBEncrypter(block, iv)
	stream.XORKeyStream(ciphertext[aes.BlockSize:], origData)
	fmt.Printf("%x\n", ciphertext)
	return ciphertext, nil
}

func GetRealString(encodingAesKey string, data string, ivKey string) string {
	dataTmp, err := hex.DecodeString(data)
	var aeskey []byte
	if err != nil {
		panic(err)
	}
	ivByte := make([]byte, 0)
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
	origData = ZeroUnPadding(origData)
	return string(origData), nil
}

func ZeroUnPadding(origData []byte) []byte {
	return bytes.TrimFunc(origData, func(r rune) bool {
		return r == rune(0)
	})
}
