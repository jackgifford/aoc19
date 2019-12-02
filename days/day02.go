package days

import (
	"fmt"
)

var machineState [5]int = [5]int{1, 0, 0, 0, 99}
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
		systemState()
	}
}

func Run() {
	fmt.Println("Test One")
	loop()

	fmt.Println("Test two")
	machineState = [5]int{2,3,0,3,99}
	resetSystem()
	loop()

}
