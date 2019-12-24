package main

import (
	"aoc19/intcode"
	"fmt"
	"os"
)

const size int = 100

var count int = 0

var isTrack bool = true
var isFin bool = false

var grid [size][size]int

var coordPair = []int{0, 0}
var currStage int = 0

func input(machineName int) int {
	if currStage == 0 {
		currStage = (currStage + 1) % 2
		return coordPair[0]
	}
	currStage = (currStage + 1) % 2
	return coordPair[1]
}

var gauntlent bool = true
var hasStarted bool = false

func output(inputVal int, machineName int) {
	//grid[coordPair[0]][coordPair[1]] = inputVal
	fmt.Printf("Machine out %v\n", inputVal)

	if inputVal == 1 {
		count += 1
		hasStarted = true
	}

	if inputVal == 0 && hasStarted {
		// game over
		hasStarted = false
		gauntlent = false
	}
	//fmt.Println(inputVal)
}

func runLine(x int, startY int) {
	count = 0
	y := startY
	gauntlent = true
	for gauntlent {
		coordPair[0] = x
		coordPair[1] = y

		code := []int{109, 424, 203, 1, 21101, 11, 0, 0, 1105, 1, 282, 21102, 1, 18, 0, 1105, 1, 259, 2102, 1, 1, 221, 203, 1, 21102, 1, 31, 0, 1106, 0, 282, 21101, 38, 0, 0, 1105, 1, 259, 20102, 1, 23, 2, 21202, 1, 1, 3, 21102, 1, 1, 1, 21102, 57, 1, 0, 1105, 1, 303, 2102, 1, 1, 222, 21002, 221, 1, 3, 21002, 221, 1, 2, 21102, 1, 259, 1, 21102, 1, 80, 0, 1105, 1, 225, 21102, 145, 1, 2, 21102, 91, 1, 0, 1106, 0, 303, 1201, 1, 0, 223, 20102, 1, 222, 4, 21102, 259, 1, 3, 21101, 225, 0, 2, 21102, 1, 225, 1, 21101, 0, 118, 0, 1105, 1, 225, 21001, 222, 0, 3, 21101, 91, 0, 2, 21102, 1, 133, 0, 1106, 0, 303, 21202, 1, -1, 1, 22001, 223, 1, 1, 21102, 148, 1, 0, 1105, 1, 259, 1201, 1, 0, 223, 21001, 221, 0, 4, 20101, 0, 222, 3, 21101, 20, 0, 2, 1001, 132, -2, 224, 1002, 224, 2, 224, 1001, 224, 3, 224, 1002, 132, -1, 132, 1, 224, 132, 224, 21001, 224, 1, 1, 21102, 195, 1, 0, 105, 1, 109, 20207, 1, 223, 2, 20101, 0, 23, 1, 21102, -1, 1, 3, 21101, 0, 214, 0, 1105, 1, 303, 22101, 1, 1, 1, 204, 1, 99, 0, 0, 0, 0, 109, 5, 1202, -4, 1, 249, 22101, 0, -3, 1, 22102, 1, -2, 2, 21202, -1, 1, 3, 21102, 1, 250, 0, 1106, 0, 225, 21201, 1, 0, -4, 109, -5, 2105, 1, 0, 109, 3, 22107, 0, -2, -1, 21202, -1, 2, -1, 21201, -1, -1, -1, 22202, -1, -2, -2, 109, -3, 2106, 0, 0, 109, 3, 21207, -2, 0, -1, 1206, -1, 294, 104, 0, 99, 21201, -2, 0, -2, 109, -3, 2105, 1, 0, 109, 5, 22207, -3, -4, -1, 1206, -1, 346, 22201, -4, -3, -4, 21202, -3, -1, -1, 22201, -4, -1, 2, 21202, 2, -1, -1, 22201, -4, -1, 1, 21201, -2, 0, 3, 21102, 343, 1, 0, 1106, 0, 303, 1105, 1, 415, 22207, -2, -3, -1, 1206, -1, 387, 22201, -3, -2, -3, 21202, -2, -1, -1, 22201, -3, -1, 3, 21202, 3, -1, -1, 22201, -3, -1, 2, 22101, 0, -4, 1, 21101, 384, 0, 0, 1105, 1, 303, 1105, 1, 415, 21202, -4, -1, -4, 22201, -4, -3, -4, 22202, -3, -2, -2, 22202, -2, -4, -4, 22202, -3, -2, -3, 21202, -4, -1, -2, 22201, -3, -2, 1, 21202, 1, 1, -4, 109, -5, 2105, 1, 0}
		emptyList := make([]int, 1000)
		code = append(code, emptyList...)
		machine := intcode.NewProgram(code, 0, output, input)
		machine.Loop()
		y++
	}
	fmt.Printf("Count: %v\n", count)
}

func main() {
	//runLine(9, 0)
	//runLine(10, 0)
	//runLine(11, 0)
	//runLine(12, 0)
	//runLine(13, 0)
	//runLine(14, 0)
	//runLine(15, 0)
	//runLine(16, 0)
	//runLine(17, 0)
	//runLine(18, 0)
	x := 10
	for x < 800 {
		runLine(x, 0)
		if count >= 100 {
			os.Exit(1)
		}
		x++
	}
	//runLine(80, 0)
	//runLine(90, 0)

	/*
		for x := 0; x < size; x++ {
			for y := 0; y < size; y++ {
				coordPair[0] = x
				coordPair[1] = y
				machine := intcode.NewProgram(	code, 0, output, input)
				machine.Loop()
			}
		} */

		/*
	for x := 0; x < size; x++ {
		for y := 0; y < size; y++ {
			val := "#"
			if grid[x][y] == 0 {
				val = "."
			}
			fmt.Printf("%v", val)
		}

		fmt.Printf("\n")
	}
	*/
}
