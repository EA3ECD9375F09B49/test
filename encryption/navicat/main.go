package main

import (
	"bufio"
	"crypto/aes"
	"crypto/cipher"
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"os"
	"strings"
)

type NavicatPassword struct {
	version    int
	aesKey     string
	aesIv      string
	blowString string
	blowKey    []byte
	blowIv     []byte
}

func NewNavicatPassword(version int) *NavicatPassword {
	np := &NavicatPassword{
		version:    version,
		aesKey:     "libcckeylibcckey",
		aesIv:      "libcciv libcciv ",
		blowString: "3DC5CA39",
		blowKey:    nil,
		blowIv:     nil,
	}
	data := md5.Sum([]byte(np.blowString))
	np.blowKey = data[:]
	np.blowIv, _ = hex.DecodeString("d9c7c3c8870d64bd")
	return np
}

func (np *NavicatPassword) EncryptTwelve(str string) string {
	key := []byte(np.aesKey)
	iv := []byte(np.aesIv)

	block, err := aes.NewCipher(key)
	if err != nil {
		panic(err)
	}

	padLen := aes.BlockSize - len(str)%aes.BlockSize
	padText := []byte{byte(padLen)}
	padText = append(padText, make([]byte, padLen-1)...)
	strBytes := []byte(str)
	strBytes = append(strBytes, padText...)

	mode := cipher.NewCBCEncrypter(block, iv)
	ciphertext := make([]byte, len(strBytes))
	mode.CryptBlocks(ciphertext, strBytes)

	return fmt.Sprintf("%X", ciphertext)
}

func (np *NavicatPassword) DecryptTwelve(encrypted string) string {
	key := []byte(np.aesKey)
	iv := []byte(np.aesIv)

	block, err := aes.NewCipher(key)
	if err != nil {
		panic(err)
	}

	encryptedBytes, _ := hex.DecodeString(encrypted)
	mode := cipher.NewCBCDecrypter(block, iv)
	decrypted := make([]byte, len(encryptedBytes))
	mode.CryptBlocks(decrypted, encryptedBytes)

	return string(decrypted)
}

func main() {
	np := NewNavicatPassword(12)

	//encrypted := "B2B566B7FE53F4B6FB0360EDEBA2F1616BC5225E56619017C6DC82C8CD952085"
	//decrypted := np.DecryptTwelve(encrypted)
	//fmt.Println(decrypted)
	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Printf("please enter string to decrypt\n")
		data, err := reader.ReadString('\n')
		data = strings.Replace(data, "\n", "", -1)
		if err != nil {
			return
		}
		fmt.Printf("The word to decrypt is %s \n",
			data)
		decrypted := np.DecryptTwelve(data)
		fmt.Printf("decString is : %s\n \n", decrypted)
	}
}
