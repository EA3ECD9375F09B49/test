package main

import (
	"crypto/aes"
	"crypto/cipher"
	"encoding/base64"
	"errors"
	"fmt"
)

func main() {
	//encodingAesKey := "u9PBrZ2nudh8sD0GMQ9cWYznZOLe7CSd"
	//data := "CT+QSxoHVm68PamNsuFv4BFDMFdfjYkSvc+qyhnYbHzn1LnEBcxuON0aH0L5dMaD5fd07fa6a9MRPOTaY/4rSX0a64TCGiDC2u+eohlfiVm6oTxhOiqQq0dYZb/INt8luXveUxjjaPkwXhzrNg6mpw=="

	fmt.Printf("please enter decodingAesKey\n")
	var encodingAesKey string
	scan, err := fmt.Scan(&encodingAesKey)
	if err != nil {
		return
	}
	fmt.Printf("The decodingAesKey is %s and scan status is %d \n",
		encodingAesKey, scan)

	fmt.Printf("please enter ivKey\n")

	var data string
	for {
		fmt.Printf("please enter string to decrypt\n")
		scan1, err := fmt.Scan(&data)
		if err != nil {
			return
		}
		fmt.Printf("The word to decrypt is %s and scan status is %d \n",
			data, scan1)
		decString, err := AesDecrypt(data, encodingAesKey)
		fmt.Printf("decString is : %v \n", decString)
	}

}

func AesDecrypt(str string, key string) (string, error) {
	if str == "" {
		fmt.Printf("AesDecrypt base64.StdEncoding.DecodeString(str) is empty")
		return "", errors.New("AesDecrypt base64.StdEncoding.DecodeString(str) is empty")
	}

	crypted, err := base64.StdEncoding.DecodeString(str)
	if err != nil {
		fmt.Printf("AesDecrypt base64.StdEncoding.DecodeString(str)", err.Error())
		return "", err
	}
	block, err := aes.NewCipher([]byte(key))
	if err != nil {
		fmt.Printf("AesDecrypt aes.NewCipher(key)", err.Error())
		return "", err
	}
	blockMode := NewECBDecrypter(block)
	origData := make([]byte, len(crypted))
	blockMode.CryptBlocks(origData, crypted)
	origData = PKCS5UnPadding(origData)
	return string(origData), nil
}

type ecbDecrypter ecb

type ecb struct {
	b         cipher.Block
	blockSize int
}

func newECB(b cipher.Block) *ecb {
	return &ecb{
		b:         b,
		blockSize: b.BlockSize(),
	}
}

func NewECBDecrypter(b cipher.Block) cipher.BlockMode {
	return (*ecbDecrypter)(newECB(b))
}
func (x *ecbDecrypter) BlockSize() int { return x.blockSize }
func (x *ecbDecrypter) CryptBlocks(dst, src []byte) {
	if len(src)%x.blockSize != 0 {
		panic("crypto/cipher: input not full blocks")
	}
	if len(dst) < len(src) {
		panic("crypto/cipher: output smaller than input")
	}
	for len(src) > 0 {
		x.b.Decrypt(dst, src[:x.blockSize])
		src = src[x.blockSize:]
		dst = dst[x.blockSize:]
	}
}

func PKCS5UnPadding(origData []byte) []byte {
	length := len(origData)
	// 去掉最后一个字节 unpadding 次
	unpadding := int(origData[length-1])
	return origData[:(length - unpadding)]
}
