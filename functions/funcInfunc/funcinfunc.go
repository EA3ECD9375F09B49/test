package main

import "fmt"

func f2() int {
	return 10
}

func f5(x func() int) func(int,int)int{
	return ff
}
func ff(a,b int)int  {
	return a+b
}
func main()  {
  f7 := f5(f2)
  aa := f7(1,2)
  fmt.Printf("%v",aa)
}
