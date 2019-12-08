package main

import (
	"fmt"
	"aoc19/intcode"
)

type inputs struct {
	first int
	second int
	called bool
}

var amps map[int]inputs

var currAmp int = 0

//var program = []int{3, 15, 3, 16, 1002, 16, 10, 16, 1, 16, 15, 15, 4, 15, 99, 0, 0}
//var program = []int{3,23,3,24,1002,24,10,24,1002,23,-1,23,101,5,23,23,1,24,23,23,4,23,99,0,0}
//var program = []int{3,31,3,32,1002,32,10,32,1001,31,-2,31,1007,31,0,33,1002,33,7,33,1,33,31,31,1,32,31,31,4,31,99,0,0,0}
var program = []int  {3,8,1001,8,10,8,105,1,0,0,21,34,59,68,89,102,183,264,345,426,99999,3,9,102,5,9,9,1001,9,5,9,4,9,99,3,9,101,3,9,9,1002,9,5,9,101,5,9,9,1002,9,3,9,1001,9,5,9,4,9,99,3,9,101,5,9,9,4,9,99,3,9,102,4,9,9,101,3,9,9,102,5,9,9,101,4,9,9,4,9,99,3,9,1002,9,5,9,1001,9,2,9,4,9,99,3,9,1002,9,2,9,4,9,3,9,101,2,9,9,4,9,3,9,1001,9,2,9,4,9,3,9,101,1,9,9,4,9,3,9,102,2,9,9,4,9,3,9,1001,9,2,9,4,9,3,9,1001,9,2,9,4,9,3,9,1001,9,2,9,4,9,3,9,1001,9,2,9,4,9,3,9,102,2,9,9,4,9,99,3,9,1001,9,1,9,4,9,3,9,102,2,9,9,4,9,3,9,102,2,9,9,4,9,3,9,101,1,9,9,4,9,3,9,101,1,9,9,4,9,3,9,102,2,9,9,4,9,3,9,102,2,9,9,4,9,3,9,101,1,9,9,4,9,3,9,1001,9,2,9,4,9,3,9,1001,9,2,9,4,9,99,3,9,1002,9,2,9,4,9,3,9,102,2,9,9,4,9,3,9,102,2,9,9,4,9,3,9,1001,9,2,9,4,9,3,9,1001,9,2,9,4,9,3,9,102,2,9,9,4,9,3,9,101,1,9,9,4,9,3,9,1001,9,1,9,4,9,3,9,1002,9,2,9,4,9,3,9,102,2,9,9,4,9,99,3,9,101,1,9,9,4,9,3,9,1002,9,2,9,4,9,3,9,102,2,9,9,4,9,3,9,1002,9,2,9,4,9,3,9,102,2,9,9,4,9,3,9,1002,9,2,9,4,9,3,9,102,2,9,9,4,9,3,9,1002,9,2,9,4,9,3,9,101,1,9,9,4,9,3,9,101,2,9,9,4,9,99,3,9,1001,9,1,9,4,9,3,9,1001,9,2,9,4,9,3,9,101,1,9,9,4,9,3,9,102,2,9,9,4,9,3,9,1001,9,2,9,4,9,3,9,1002,9,2,9,4,9,3,9,101,1,9,9,4,9,3,9,1001,9,1,9,4,9,3,9,1001,9,2,9,4,9,3,9,1002,9,2,9,4,9,99}

func output(input int) {
	fmt.Println(input)
	currAmp++
	myIn := amps[currAmp]
	myIn.second = input
	amps[currAmp] = myIn
}

func input() int {
	myIn := amps[currAmp]

	if myIn.called {
		return myIn.second
	}

	myIn.called = true
	amps[currAmp] = myIn
	return myIn.first
}

func spinup(sets []int) {
	for i, v := range sets {
		amps[i] = inputs {
			first: v,
			second: 0,
			called: false,
		}
	}

	for i := 0; i < 5; i++ {
		newProg := intcode.NewProgram(program, "A", output, input)
		newProg.Loop()
	}
	currAmp = 0
	fmt.Printf("Final: %v\n", amps[5].second)
}

func heaps_algo(sets []int, length int) {
	if length == 1 {
		fmt.Println(sets)
		spinup(sets)
		return
	}

	for i := 0; i < length-1; i++ {
		heaps_algo(sets, length-1)

		if length%2 == 0 {
			temp := sets[length-1]
			sets[length-1] = sets[i]
			sets[i] = temp
		} else {
			temp := sets[length-1]
			sets[length-1] = sets[0]
			sets[0] = temp
		}
	}

	heaps_algo(sets, length-1)
}

func main() {
	amps = make(map[int]inputs, 0)
	temp := []int{1, 2, 3, 4, 0}
	heaps_algo(temp, 5)
	//temp := []int{4,3,2,1,0}
	//spinup(temp)
}
