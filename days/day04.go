package main

import (
	"fmt"
)

func hasOne(arr []int) bool {
	for _, elem := range arr {
		if elem == 1 {
			return true
		}
	}

	return false
}

func adjCheckSimple(arr [6]int) bool {
	for i := 1; i < 6; i++ {
		if arr[i] == arr[i-1] {
			return true
		}
	}

	return false
}

func adjCheck(arr [6]int) bool {
	var adjLen []int
	currLen := 0

	for i := 1; i < 6; i++ {
		if arr[i] == arr[i-1] {
			currLen += 1
		} else {
			adjLen = append(adjLen, currLen)
			currLen = 0
		}
	}

	adjLen = append(adjLen, currLen)
	return hasOne(adjLen)
}

func leftToRightCheck(arr [6]int) bool {
	for i := 1; i < 6; i++ {
		if arr[i] < arr[i-1] {
			return false
		}
	}

	return true
}

func intToArr(intVal int) [6]int {
	var intArr [6]int
	for i := 5; i >= 0; i-- {
		intArr[i] = intVal % 10
		intVal /= 10
	}

	return intArr
}

var totalValid int = 0

func testRunner(intVal int, adjCheckFunc func([6]int) bool) {
	testArr := intToArr(intVal)
	if adjCheckFunc(testArr) && leftToRightCheck(testArr) {
		totalValid += 1
	}
}

func main() {
	for currval := 234208; currval <= 765869; currval++ {
		testRunner(currval, adjCheckSimple)
	}

	fmt.Printf("Part One: %d\n", totalValid)
	totalValid = 0

	for currval := 234208; currval <= 765869; currval++ {
		testRunner(currval, adjCheck)
	}

	fmt.Printf("Part Two: %d\n", totalValid)
}
