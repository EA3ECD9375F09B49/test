package main

import "fmt"

func main() {
	message := "0123456789afcdefghijklmn"
	//for i := 0; i < len(message); i += 2 {
	//	fmt.Println(string(message[i]))
	//}
	for i := 0; i < len(message); i++ {
		if i%2 == 1 {
			fmt.Println("even:" + string(message[i]))
		} else {
			fmt.Println("odd:" + string(message[i]))
		}

	}
}
