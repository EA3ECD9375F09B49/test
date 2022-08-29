package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	paragraph := "HelloX42ImaY3235stringZ30HelloX42ImaY3235stringZ30HelloX42ImaY3235stringZ30HelloX42ImaY3235stringZ30HelloX42ImaY3235stringZ30HelloX42ImaY3235stringZ30HelloX42ImaY3235stringZ30HelloX42ImaY3235stringZ30HelloX42ImaY3235stringZ30HelloX42ImaY3235stringZ30HelloX42ImaY3235stringZ30HelloX42ImaY3235stringZ30HelloX42ImaY32"
	start, end := getSubStringParagraphIndex(paragraph, 32)
	fmt.Printf("start:%d|end:%d", start, end)
}

func getSubStringParagraphIndex(paragraph string, stringIndexLength int) (int, int) {
	paraLength := len(paragraph)
	s := rand.NewSource(time.Now().UnixNano())
	r := rand.New(s) // initialize local pseudorandom generator
	start := r.Intn(paraLength)
	end := start + stringIndexLength
	if end > paraLength {
		start, end = getSubStringParagraphIndex(paragraph, stringIndexLength)
	}
	return start, end
}
