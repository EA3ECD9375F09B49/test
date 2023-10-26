package main

import (
	"fmt"
	"time"
)

var funcMap = map[string]func(value string) (<-chan string, error){
	"YBTY": Handler,
}

func Handler(value string) (<-chan string, error) {
	ops := make(chan string, 3)
	go func() {
		defer close(ops)

		for i := 0; i < 3; i++ {
			err := packOpt(ops, value, i)
			if err != nil {
				return
			}
			time.Sleep(10 * time.Second)
		}
	}()
	return ops, nil
}

func packOpt(ops chan string, value string, count int) error {
	ops <- fmt.Sprintf("value %v count %d success", value, count)
	return nil
}

func main() {

	key := "YBTY"

	handler, ok := funcMap[key]
	if !ok {
		fmt.Println("error")
	}

	ops, _ := handler("Im trying " + key)

	for op := range ops {
		fmt.Println("结果为", op)
	}
}
