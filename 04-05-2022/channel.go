package main

import (
	"fmt"
	"sort"
	"sync"
)

type NumIndex struct {
	Offset int
	NumArr []*NUm
}

type NUm struct {
	A int
	B int
}

func main() {
	count := 100000 // 总条数

	lcount := 3000 //一次获取条数

	size := count/lcount + 1

	wg := sync.WaitGroup{}
	c := make(chan *NumIndex, size)

	for i := 0; i < size; i++ {
		wg.Add(1)
		j := i
		offset := j * lcount // 每次偏移量
		go getNum(offset, lcount, &wg, c)
	}
	wg.Wait()

	var numIndexArr []*NumIndex
	for i := 0; i < size; i++ {
		select {
		case tmp := <-c:
			numIndexArr = append(numIndexArr, tmp)
			//fmt.Println("tmp====== ", tmp)
		}
	}

	sort.Slice(numIndexArr, func(i, j int) bool {
		return numIndexArr[i].Offset < numIndexArr[j].Offset
	})

	for i := range numIndexArr {
		a := numIndexArr[i]
		fmt.Println("index: ", a.Offset)
	}

	fmt.Println("done!!!!!!!!!!!!")
}

func getNum(offset, lcount int, wg *sync.WaitGroup, c chan *NumIndex) {
	defer wg.Done()

	var tmpArr []*NUm

	tmp := &NUm{
		A: 0,
		B: 0,
	}

	tmpArr = append(tmpArr, tmp)

	var tmpNumIndex NumIndex
	tmpNumIndex.NumArr = tmpArr
	tmpNumIndex.Offset = offset

	c <- &tmpNumIndex
	return
}
