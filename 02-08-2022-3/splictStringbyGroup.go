package main

import (
	"fmt"
	"regexp"
	"sort"
	"strconv"
	"strings"
)

func main() {
	message := "10,4,p,32"
	//message := "1,3,a,3,14,19,5,b,13,2,1e,23,4,1,8,20"
	//lastInd := strings.LastIndex(message, ",")
	//fmt.Printf("arrayString %v \n", message[lastInd+1:])
	regex := *regexp.MustCompile(`(?s)([0-9]+),([0-9]+),([p|np]),([0-9]+)`)
	res := regex.FindAllStringSubmatch(message, -1)
	var baseString string
	var groupString string
	var positionString string
	var scLengthString string

	//var base int
	var group int
	for i := range res {
		baseString = res[i][1]
		groupString = res[i][2]
		positionString = res[i][3]
		scLengthString = res[i][4]
		fmt.Printf("%s,%s,%s,%s\n", baseString, groupString, positionString, scLengthString)
	}
	//base, _ = strconv.Atoi(baseString)
	group, _ = strconv.Atoi(groupString)

	RuleFormula := map[string]string{
		"10_p":  "%d,%d,%d",
		"10_np": "%d,%d",
		"16_p":  "%s,%s,%s",
		"16_np": "%s,%s",
		"32_p":  "%s,%s,%s",
		"32_np": "%s,%s",
	}
	formulaIndex := baseString + "_" + positionString
	currentRuleFormula := RuleFormula[formulaIndex]
	var resultString string
	for i := 1; i <= group; i++ {
		resultString += currentRuleFormula + ","
	}
	resultString += scLengthString
	fmt.Println(resultString)
}

func sortMatrix() {
	matrix := [3][3]int{
		{2, 3, 1},
		{6, 3, 5},
		{1, 4, 9},
	}

	sort.Slice(matrix[:], func(i, j int) bool {
		for x := range matrix[i] {
			if matrix[i][x] == matrix[j][x] {
				continue
			}
			return matrix[i][x] < matrix[j][x]
		}
		return false
	})

	fmt.Println(matrix) //[[1 4 9] [2 3 1] [6 3 5]]
}

func getStringbyGroup() {
	paragraph := "HelloX42ImaY3235stringZ30HelloX42ImaY3235stringZ30HelloX42ImaY3235stringZ30HelloX42ImaY3235stringZ30HelloX42ImaY3235stringZ30HelloX42ImaY3235stringZ30HelloX42ImaY3235stringZ30HelloX42ImaY3235stringZ30HelloX42ImaY3235stringZ30HelloX42ImaY3235stringZ30HelloX42ImaY3235stringZ30HelloX42ImaY3235stringZ30HelloX42ImaY32"

	message := "1,3,10,3,20,25,5,11,19,2,30,35,4,1,8,32"
	//message := "1,3,a,3,14,19,5,b,13,2,1e,23,4,1,8,20"
	lastInd := strings.LastIndex(message, ",")
	fmt.Printf("arrayString %v \n", message[lastInd+1:])
	regex := *regexp.MustCompile(`(?s)([a-vA-V0-9]+),([a-vA-V0-9]+),([a-vA-V0-9]+),?`)
	res := regex.FindAllStringSubmatch(message, -1)

	sort.Slice(res[:], func(i, j int) bool {
		sortIndexer := 1
		if res[i][sortIndexer] == res[j][sortIndexer] {
			return false
		}
		return res[i][sortIndexer] < res[j][sortIndexer]
	})

	var resultString string
	for i := range res {
		//like Java: match.group(1), match.gropu(2), etc
		fmt.Printf("%s,%s,%s\n", res[i][1], res[i][2], res[i][3])
		start, err := strconv.Atoi(res[i][2])
		end, err := strconv.Atoi(res[i][3])
		if err != nil {

		}
		resultString += string([]byte(paragraph)[start:end])
	}
	fmt.Printf("resultString :%s", resultString)
	//arrayString := strings.SplitN(message, ",", 4)

	//re := regexp.MustCompile("^\\d+(,\\d+){2}")
	//txt := message
	//
	//split := re.Split(txt, -1)
	//set := []string{}
	//
	//for i := range split {
	//	set = append(set, split[i])
	//}
	//
	//fmt.Printf("start:%v", arrayString)
}
