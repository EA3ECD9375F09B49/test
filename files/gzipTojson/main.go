package main

import (
	"bytes"
	"compress/gzip"
	"encoding/json"
	"fmt"
	"golang.org/x/text/encoding/charmap"
	"io/ioutil"
	"os"
)

func convert(utf8 []byte) (string, error) {
	latin1, err := charmap.ISO8859_1.NewEncoder().Bytes(utf8)
	if err != nil {
		return "", err
	}
	return string(latin1), nil
}

// 解压
func decompression(in []byte) ([]byte, error) {
	newIn, err := convert(in) //编码为ISO8859_1
	if err != nil {
		return nil, err
	}
	reader, err := gzip.NewReader(bytes.NewReader([]byte(newIn)))
	if err != nil {
		return nil, err
	}
	defer reader.Close()
	return ioutil.ReadAll(reader)
}

func main() {
	dir := "./fromFiles"
	toDir := "./toFiles"
	files, err := ioutil.ReadDir(dir)
	if err != nil {
		panic(err)
	}

	for _, file := range files {
		if !file.IsDir() {
			// 打印文件名
			fmt.Println(file.Name())
			// 打开文件
			f, err := os.Open(dir + "/" + file.Name())
			if err != nil {
				panic(err)
			}
			defer f.Close()
			// Create a gzip reader
			data, _ := ioutil.ReadAll(f)
			// Read the JSON data
			buf, err := decompression(data)
			if err != nil {
				panic(err)
			}
			// Unmarshal the JSON data
			var obj interface{}
			err = json.Unmarshal(buf, &obj)
			if err != nil {
				panic(err)
			}

			// Write the JSON data to another file
			output, err := json.MarshalIndent(obj, "", "  ")
			if err != nil {
				panic(err)
			}
			err = ioutil.WriteFile(toDir+"/"+"output_"+file.Name(), output, 0644)
			if err != nil {
				panic(err)
			}

		}
	}

	fmt.Println("JSON data saved to output.json")
}
