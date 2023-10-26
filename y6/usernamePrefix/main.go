package main

import (
	"fmt"
	"strings"
)

func main() {

	inputStrings := []string{"obajnt66017"}

	for _, input := range inputStrings {
		input = strings.TrimSpace(input)
		user := TrimJNUserName(input)
		fmt.Printf("字符串 \"%s\" 匹配了子字符串 \"%s\"\n", input, user)
		fmt.Printf("后缀 \"%s\" 用户 \"%s\"\n", GetJnIdPosfix(input), input)
	}

	//var aa string
	//bb := "p03p03wcg1124"
	//prefix := "p03"
	//if strings.HasPrefix(bb, prefix) {
	//	aa = strings.TrimPrefix(bb, prefix)
	//}
	//
	//aa, _ = SubString(bb, 5, len(bb))
	//
	////aa, _ := url.Parse("https://api.y1uldxc.com:17025/file/fastdfs/download/image?filePath=group1/M00/18/A3/CgURtmQauP-AamWnAAAU7b87fjk853.png")
	////aa := strings.Split("aaaa:9092", ",")
	//fmt.Println(aa)
}

func TrimJNUserName(username string) string {
	substrings := []string{"oub", "ajn", "ob"}
	for _, substr := range substrings {
		if strings.HasPrefix(username, substr+"_") {
			username = strings.TrimPrefix(username, substr+"_")
			break
		}
		if strings.HasPrefix(username, substr) {
			username = strings.TrimPrefix(username, substr)
			break
		}
	}
	return username
}

func GetJnIdPosfix(input string) string {
	substrings := []string{"oub", "ajn", "ob"}

	prefixMap := map[string]string{
		"oub": "oub",
		"ajn": "ajn",
		"ob":  "oub",
	}
	result := "oub"
	for _, substr := range substrings {
		if strings.HasPrefix(input, substr) {
			if val, ok := prefixMap[substr]; ok {
				result = val
				fmt.Printf("字符串 \"%s\" 匹配了子字符串 \"%s\"\n", input, result)
				break
			}
		}
	}
	return result
}

func SubString(source string, start, end int) (string, error) {
	var r = []rune(source)
	length := len(r)
	if start < 0 || end > length || start > end {
		return "", fmt.Errorf("SubString Out of range! source: %s, length: %d, start: %d, end: %d", source, length, start, end)
	}
	if start == 0 && end == length {
		return source, nil
	}
	return string(r[start:end]), nil
}
