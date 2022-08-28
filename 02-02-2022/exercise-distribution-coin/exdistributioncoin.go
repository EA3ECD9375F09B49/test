package main

import (
	"fmt"
)

var (
	coins = 5000
	users = []string{
		"Matthew","Sarah","Augustus","Heidi","Emilie","Peter","Giana","Adriano","Aaron","Elizabeth",
	}
	distribution = make(map[string]int,len(users))
)
func main()  {
  left :=dispatchCoin()
  fmt.Println("剩下：",left)
	fmt.Printf("%v",distribution)
}

//func dispatchCoin() int  {
//	var coin int
//	for _, name :=range users{
//		switch {
//		case strings.ContainsAny(name, "e | E"):
//			coin =  1
//		case strings.ContainsAny(name, "i | I"):
//			coin =  2
//		case strings.ContainsAny(name, "o | O"):
//			coin =  3
//		case strings.ContainsAny(name, "u | U"):
//			coin =  4
//		default:
//			coin = 0
//		}
//		if _,ok := distribution[name];!ok {
//			distribution[name]=coin
//		}else {
//			distribution[name] +=coin
//		}
//		coins -=coin
//	}
//	return coins
//}

func dispatchCoin() int  {
	var coin int
	for _, name :=range users{
		for _, c :=range name {
			switch c {
			case 'e','E':
				coin =  1
			case 'i','I':
				coin =  2
			case 'o','O':
				coin =  3
			case 'u','U':
				coin =  4
			default:
				coin = 0
			}
			if _,ok := distribution[name];!ok {
				distribution[name]=coin
			}else {
				distribution[name] +=coin
			}
			coins -=coin
		}

	}
	return coins
}