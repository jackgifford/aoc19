package days

import (
	"fmt"
)

var machineState []int
var i int = 0

func systemState() {
	fmt.Println(machineState)
}

func addOp() {
	inputOne := machineState[machineState[i+1]]
	inputTwo := machineState[machineState[i+2]]
	resReg := machineState[i+3]

	machineState[resReg] = inputOne + inputTwo
}

func multOp() {
	inputOne := machineState[machineState[i+1]]
	inputTwo := machineState[machineState[i+2]]
	resReg := machineState[i+3]

	machineState[resReg] = inputOne * inputTwo
}

func haltOp() {
	fmt.Println("Program finished :)")
	systemState()
}

func loop() {
	i = 0
	for {
		opCode := machineState[i]

		if opCode == 1 {
			addOp()
		} else if opCode == 2 {
			multOp()
		} else if opCode == 99 {
			haltOp()
			break
		} else {
			panic("Machine broken!")
		}

		i += 4
		//systemState()
	}
}

func partTwo() {
	noun := 0
	for {
		for verb := 0; verb < 100; verb++ {

			machineState = []int{1, noun, verb, 3, 1, 1, 2, 3, 1, 3, 4, 3, 1, 5, 0, 3, 2, 10, 1, 19, 1, 6, 19, 23, 2, 23, 6, 27, 1, 5, 27, 31, 1, 31, 9, 35, 2, 10, 35, 39, 1, 5, 39, 43, 2, 43, 10, 47, 1, 47, 6, 51, 2, 51, 6, 55, 2, 55, 13, 59, 2, 6, 59, 63, 1, 63, 5, 67, 1, 6, 67, 71, 2, 71, 9, 75, 1, 6, 75, 79, 2, 13, 79, 83, 1, 9, 83, 87, 1, 87, 13, 91, 2, 91, 10, 95, 1, 6, 95, 99, 1, 99, 13, 103, 1, 13, 103, 107, 2, 107, 10, 111, 1, 9, 111, 115, 1, 115, 10, 119, 1, 5, 119, 123, 1, 6, 123, 127, 1, 10, 127, 131, 1, 2, 131, 135, 1, 135, 10, 0, 99, 2, 14, 0, 0}

			loop()

			if machineState[0] == 19690720 {
				panic("we found it")
			}
		}
		noun += 1
	}
}

func Run() {
	machineState = []int{1, 0, 0, 0, 99}
	fmt.Println("Test One")
	loop()

	fmt.Println("Test two")
	machineState = []int{2, 3, 0, 3, 99}
	loop()

	fmt.Println("test three")
	machineState = []int{2, 4, 4, 5, 99, 0}
	loop()

	fmt.Println("test four")
	machineState = []int{1, 1, 1, 4, 99, 5, 6, 0, 99}
	loop()

	fmt.Println("Broken program (pre crash)")
	machineState = []int{1, 0, 0, 3, 1, 1, 2, 3, 1, 3, 4, 3, 1, 5, 0, 3, 2, 10, 1, 19, 1, 6, 19, 23, 2, 23, 6, 27, 1, 5, 27, 31, 1, 31, 9, 35, 2, 10, 35, 39, 1, 5, 39, 43, 2, 43, 10, 47, 1, 47, 6, 51, 2, 51, 6, 55, 2, 55, 13, 59, 2, 6, 59, 63, 1, 63, 5, 67, 1, 6, 67, 71, 2, 71, 9, 75, 1, 6, 75, 79, 2, 13, 79, 83, 1, 9, 83, 87, 1, 87, 13, 91, 2, 91, 10, 95, 1, 6, 95, 99, 1, 99, 13, 103, 1, 13, 103, 107, 2, 107, 10, 111, 1, 9, 111, 115, 1, 115, 10, 119, 1, 5, 119, 123, 1, 6, 123, 127, 1, 10, 127, 131, 1, 2, 131, 135, 1, 135, 10, 0, 99, 2, 14, 0, 0}
	loop()

	fmt.Println("Broken program (pre crash)")

	machineState = []int{1, 12, 2, 3, 1, 1, 2, 3, 1, 3, 4, 3, 1, 5, 0, 3, 2, 10, 1, 19, 1, 6, 19, 23, 2, 23, 6, 27, 1, 5, 27, 31, 1, 31, 9, 35, 2, 10, 35, 39, 1, 5, 39, 43, 2, 43, 10, 47, 1, 47, 6, 51, 2, 51, 6, 55, 2, 55, 13, 59, 2, 6, 59, 63, 1, 63, 5, 67, 1, 6, 67, 71, 2, 71, 9, 75, 1, 6, 75, 79, 2, 13, 79, 83, 1, 9, 83, 87, 1, 87, 13, 91, 2, 91, 10, 95, 1, 6, 95, 99, 1, 99, 13, 103, 1, 13, 103, 107, 2, 107, 10, 111, 1, 9, 111, 115, 1, 115, 10, 119, 1, 5, 119, 123, 1, 6, 123, 127, 1, 10, 127, 131, 1, 2, 131, 135, 1, 135, 10, 0, 99, 2, 14, 0, 0}

	loop()

	partTwo()

}
