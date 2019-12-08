package main

import (
	"fmt"
	"aoc19/intcode"
	"sync"
)

type inputs struct {
	first int
	second int
	called bool
	channel chan int
	running bool
}

var amps map[int]inputs

var currAmp int = 0
//var program = []int{3,26,1001,26,-4,26,3,27,1002,27,2,27,1,27,26, 27,4,27,1001,28,-1,28,1005,28,6,99,0,0,5}

var program = []int  {3,8,1001,8,10,8,105,1,0,0,21,34,59,68,89,102,183,264,345,426,99999,3,9,102,5,9,9,1001,9,5,9,4,9,99,3,9,101,3,9,9,1002,9,5,9,101,5,9,9,1002,9,3,9,1001,9,5,9,4,9,99,3,9,101,5,9,9,4,9,99,3,9,102,4,9,9,101,3,9,9,102,5,9,9,101,4,9,9,4,9,99,3,9,1002,9,5,9,1001,9,2,9,4,9,99,3,9,1002,9,2,9,4,9,3,9,101,2,9,9,4,9,3,9,1001,9,2,9,4,9,3,9,101,1,9,9,4,9,3,9,102,2,9,9,4,9,3,9,1001,9,2,9,4,9,3,9,1001,9,2,9,4,9,3,9,1001,9,2,9,4,9,3,9,1001,9,2,9,4,9,3,9,102,2,9,9,4,9,99,3,9,1001,9,1,9,4,9,3,9,102,2,9,9,4,9,3,9,102,2,9,9,4,9,3,9,101,1,9,9,4,9,3,9,101,1,9,9,4,9,3,9,102,2,9,9,4,9,3,9,102,2,9,9,4,9,3,9,101,1,9,9,4,9,3,9,1001,9,2,9,4,9,3,9,1001,9,2,9,4,9,99,3,9,1002,9,2,9,4,9,3,9,102,2,9,9,4,9,3,9,102,2,9,9,4,9,3,9,1001,9,2,9,4,9,3,9,1001,9,2,9,4,9,3,9,102,2,9,9,4,9,3,9,101,1,9,9,4,9,3,9,1001,9,1,9,4,9,3,9,1002,9,2,9,4,9,3,9,102,2,9,9,4,9,99,3,9,101,1,9,9,4,9,3,9,1002,9,2,9,4,9,3,9,102,2,9,9,4,9,3,9,1002,9,2,9,4,9,3,9,102,2,9,9,4,9,3,9,1002,9,2,9,4,9,3,9,102,2,9,9,4,9,3,9,1002,9,2,9,4,9,3,9,101,1,9,9,4,9,3,9,101,2,9,9,4,9,99,3,9,1001,9,1,9,4,9,3,9,1001,9,2,9,4,9,3,9,101,1,9,9,4,9,3,9,102,2,9,9,4,9,3,9,1001,9,2,9,4,9,3,9,1002,9,2,9,4,9,3,9,101,1,9,9,4,9,3,9,1001,9,1,9,4,9,3,9,1001,9,2,9,4,9,3,9,1002,9,2,9,4,9,99}

func output(input int, name int) {
	index := (name + 1) % 5
	myIn := amps[index]
	
	if !myIn.running {
		fmt.Printf("output from %v\n", name)
		fmt.Printf("max:%v\n", input)
	} else {
		fmt.Printf("output from %v\n", name)
		fmt.Println(input)
	}


	if myIn.running {
		myIn.channel <- input
	}
	/*
	currAmp++
	myIn := amps[currAmp]
	myIn.second = input
	amps[currAmp] = myIn
	*/
}

func input(name int) int {
		fmt.Printf("input from %v\n", name)
	myIn := amps[name]

	if myIn.called {
		// get chan
		return <-myIn.channel
		//return myIn.second
		// return myIn.second
		// block on input
	}

	myIn.called = true
	amps[name] = myIn
	return myIn.first
}

func start(i int, wg *sync.WaitGroup) {
		newProg := intcode.NewProgram(program, i, output, input)
		newProg.Loop()
		inputDa := amps[i] 
		inputDa.running = false
		amps[i] = inputDa
		close(amps[i].channel)
		wg.Done()
}

func spinup(sets []int) {
	for i, v := range sets {
		amps[i] = inputs {
			first: v,
			second: 0,
			called: false,
			channel: make(chan int),
			running: true,
		}
	}

	wg := &sync.WaitGroup{}
	for i := 0; i < 5; i++ {
		wg.Add(1)
		go start(i, wg)
	}

	amps[0].channel <- 0
	currAmp = 0
	wg.Wait()
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
	temp := []int{5, 6, 7, 8, 9}
	heaps_algo(temp, 5)
	//temp := []int{9,8,7,6,5}
	//spinup(temp)
}
