package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"sort"
	"strconv"
	"strings"
	"time"
)

func main() {
	//var brr = []int{0, 1, 12, 14, 40, 56, 60, 88, 78, 99}
	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Printf("please enter target \n")
		var target int
		sumString, err := reader.ReadString('\n')
		// convert CRLF to LF
		sumString = strings.Replace(sumString, "\n", "", -1)
		target, err = strconv.Atoi(sumString)
		if err != nil {
			return
		}
		fmt.Printf("The target is %d  \n",
			target)

		var start int
		fmt.Printf("please enter start int \n")

		startString, err := reader.ReadString('\n')
		// convert CRLF to LF
		startString = strings.Replace(startString, "\n", "", -1)
		start, err = strconv.Atoi(startString)
		if err != nil {
			return
		}
		fmt.Printf("The start int is %d  \n",
			start)

		var end int
		fmt.Printf("please enter end int \n")

		endstring, err := reader.ReadString('\n')
		// convert CRLF to LF
		endstring = strings.Replace(endstring, "\n", "", -1)
		end, err = strconv.Atoi(endstring)
		if err != nil {
			return
		}
		fmt.Printf("The end int is %d  \n",
			end)

		var companion int
		fmt.Printf("please enter companion int \n")
		companionString, err := reader.ReadString('\n')
		// convert CRLF to LF
		companionString = strings.Replace(companionString, "\n", "", -1)
		companion, err = strconv.Atoi(companionString)
		if err != nil {
			return
		}
		fmt.Printf("The companion int is %d  \n",
			end)

		fmt.Println(deviationOfSum(companion, start, end, target))
	}
}

func deviationOfSum(count int, start int, end int, target int) []int {
	tables := make(map[int]int)
	//nums := make([]int, end+1)
	var nums []int
	result := make([]int, count)
	k := 0
	for i := start; i <= end; i++ {
		//nums[i] = i
		nums = append(nums, i)
		tables[k] = nums[k]
		k++
	}
	tmpSum := 0
	for j := 0; j < count; j++ {
		var t int
		if j+1 == count {
			t = target - tmpSum
		} else {
			t = getRanDomNumberDeviation(nums, tables)
			if j > 0 {
				tmp := target - (tmpSum + t)
				if tmp < 0 {
					t = 0
				} else {
					t = tmp
				}
			}
		}
		tmpSum += t
		if tmpSum <= target {
			result[j] = t
		}
	}
	sort.Ints(result)
	for _, v := range result {
		if v > end {
			result = deviationOfSum(count, start, end, target)
			break
		}
	}
	return result
}

func getRanDomNumberDeviation(nums []int, tables map[int]int) int {
	t := getRandomIntData(nums)
	if _, ok := tables[t]; ok {
	} else {
		t = getRanDomNumberDeviation(nums, tables)
	}
	return t
}

func getRandomIntData(nums []int) int {
	s := rand.NewSource(time.Now().Unix())
	r := rand.New(s) // initialize local pseudorandom generator
	n := r.Intn(len(nums))
	t := nums[n]
	return t
}
