package main

import "fmt"

func main() {
	labels := make([]int, 100)
	for i := 0; i < 100; i++ {
		labels[i] = i
	}

	chIn := make(chan int, 20)
	chOut := make(chan int, 20)

	for i := 0; i < 10; i++ {
		go func(chIn chan int, chOut chan<- int) {

			for gmailLabels := range chIn {
				fmt.Println("processing", gmailLabels)
				// Performs some operation with the label `d`

				chOut <- gmailLabels
			}

		}(chIn, chOut)

	}

	go func(chIn chan int) {
		defer close(chIn)
		for _, l := range labels {
			fmt.Println("sending", l)

			chIn <- l
		}
	}(chIn)

	for i := 0; i < len(labels); i++ {
		lab := <-chOut
		fmt.Printf("Done %v\n", lab)
	}
}
