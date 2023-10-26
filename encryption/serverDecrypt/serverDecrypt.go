package main

import (
	"crypto/aes"
	"crypto/cipher"
	"encoding/base64"
)

func main() {
	res := "2Kgnyg/2eflRNCdZz6Mncm0HDOP42rat5i1c+0c1PU0DDW3zRCK6o/+klAT3AEP1IjHbYEVM6MxpduA1GpIbvHBX6Go33ElqW82G2EjX8q4="
	dt, err := base64.StdEncoding.DecodeString(res)
	if err != nil {
		return
	}
	temp := AesDecrypt(dt, []byte("01F5B556E6511930"))
	result := string(temp)
	println(result)
}

func AesDecrypt(crypted, key []byte) []byte {
	block, err := aes.NewCipher(key)
	if err != nil {
		return []byte("")
	}
	blockMode := NewECBDecrypter(block)
	origData := make([]byte, len(crypted))
	blockMode.CryptBlocks(origData, crypted)
	origData = PKCS5UnPadding(origData)
	return origData
}

func NewECBDecrypter(b cipher.Block) cipher.BlockMode {
	return (*ecbDecrypter)(newECB(b))
}

func newECB(b cipher.Block) *ecb {
	return &ecb{
		b:         b,
		blockSize: b.BlockSize(),
	}
}

type ecb struct {
	b         cipher.Block
	blockSize int
}

type ecbDecrypter ecb

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

// PKCS5UnPadding ...
func PKCS5UnPadding(origData []byte) []byte {
	length := len(origData)
	// 去掉最后一个字节 unpadding 次
	unpadding := int(origData[length-1])
	if length <= unpadding {
		return nil
	}
	return origData[:(length - unpadding)]
}
