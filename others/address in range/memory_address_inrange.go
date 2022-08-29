package main

import "fmt"

//for range 经典bug 坑的解释
func main() {
	a := []int{1, 2, 3, 4}

	//遍历切片
	for k, v := range a {
		fmt.Printf("key = %#v and value = %#v of original \n", k, v)
		fmt.Printf("original key and value address are %v == %v \n", &k, &v) //地址相同
	}

	//遍历切片
	for k, v := range a {
		kk, vv := k, v //new一个变量
		fmt.Printf("key k= %v and value v= %v of original \n", k, v)
		fmt.Printf("key kk= %v and value vv= %v of new \n", kk, vv)
		fmt.Printf("new key and value address are %v == %v \n", &kk, &vv) //地址不同，达到要求
	}
}
