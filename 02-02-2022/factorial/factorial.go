package main

import (
	"fmt"
	"time"
)

const LIM = 41

func Factorial(n uint64) (result uint64) {
	if n > 0 {
		result = n * Factorial(n-1)
		return result
	}
	return 1
}

func main() {
	fmt.Println("==================FACTORIAL==================")
	start := time.Now()
	for i := uint64(0); i < LIM; i++ {
		fmt.Printf("Factorial for %d is : %d \n", i, Factorial(uint64(i)))
	}
	end := time.Now()
	fmt.Printf("Calculation finished in %s \n", end.Sub(start)) //Calculation finished in 2.0002ms
}
