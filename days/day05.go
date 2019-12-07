package main

import (
	"fmt"
)

var machineState []int
var i int = 0

func systemState() {
	fmt.Printf("Final state: ")
	fmt.Println(machineState)
}

func addOp(inputOne int, inputTwo int) {
	resReg := machineState[i+3]

	machineState[resReg] = inputOne + inputTwo
}

func multOp(inputOne int, inputTwo int) {
	resReg := machineState[i+3]

	machineState[resReg] = inputOne * inputTwo
}

func inputOp(location int) {
	fmt.Println("Machine wanted input gave 5")
	machineState[location] = 5
}

func outputOp(input int) {
	fmt.Println(input)
}

func jumpIfTrueOp(inputOne int, inputTwo int) bool {
	if inputOne != 0 {
		fmt.Println("jumping")
		i = inputTwo
	}

	return inputOne != 0
}

func jumpIfFalseOp(inputOne int, inputTwo int) bool {
	if inputOne == 0 {
		fmt.Println("jumping")
		i = inputTwo
	}

	return inputOne == 0
}

func lessThanOp(inputOne int, inputTwo int) {
	resReg := machineState[i+3]
	if inputOne < inputTwo {
		machineState[resReg] = 1
	} else {
		machineState[resReg] = 0
	}
}

func equalsOp(inputOne int, inputTwo int) {
	resReg := machineState[i+3]
	if inputOne == inputTwo {
		machineState[resReg] = 1
	} else {
		machineState[resReg] = 0
	}
}

func haltOp() {
	systemState()
	fmt.Println("Program finished :)")
}

func validate(position int) (int, [3]int) {
	instruction := machineState[position]
	opcode := instruction % 100
	instruction /= 100

	var arr [3]int
	var argCount int

	if opcode == 1 || opcode == 2 || opcode == 5 || opcode == 6 || opcode == 7 || opcode == 8 {
		argCount = 2
	} else if opcode == 3 || opcode == 4 {
		argCount = 1
	} else {
		argCount = 0
	}

	for curr := 0; curr < argCount; curr++ {
		modeType := instruction % 10
		instruction /= 10
		if modeType == 1 {
			arr[curr] = machineState[curr+position+1]
		} else {
			arr[curr] = machineState[machineState[curr+position+1]]
		}
	}

	return opcode, arr
}

func loop() {
	i = 0
	for {
		opCode, inputs := validate(i)

		if opCode == 1 {
			addOp(inputs[0], inputs[1])
			i += 4
		} else if opCode == 2 {
			multOp(inputs[0], inputs[1])
			i += 4
		} else if opCode == 3 {
			inputOp(machineState[i+1])
			i += 2
		} else if opCode == 4 {
			outputOp(inputs[0])
			i += 2
		} else if opCode == 5 {
			didJump := jumpIfTrueOp(inputs[0], inputs[1])
			if !didJump {
				i += 3
			}
		} else if opCode == 6 {
			didJump := jumpIfFalseOp(inputs[0], inputs[1])
			if !didJump {
				i += 3
			}

		} else if opCode == 7 {
			lessThanOp(inputs[0], inputs[1])
			i += 4
		} else if opCode == 8 {
			equalsOp(inputs[0], inputs[1])
			i += 4
		} else if opCode == 99 {
			haltOp()
			break
		} else {
			//fmt.Println(machineState)
			fmt.Printf("Unknown op: %d\n", opCode)
			panic("Machine broken!")
		}
	}
}

func main() {
	//machineState = []int{3,0,4,0,99}
	//machineState= []int{3,12,6,12,15,1,13,14,13,4,13,99,-1,0,1,9}
	//machineState=[]int{3,12,6,12,15,1,13,14,13,4,13,99,-1,0,1,9}
	machineState = []int{3, 225, 1, 225, 6, 6, 1100, 1, 238, 225, 104, 0, 1001, 152, 55, 224, 1001, 224, -68, 224, 4, 224, 1002, 223, 8, 223, 1001, 224, 4, 224, 1, 224, 223, 223, 1101, 62, 41, 225, 1101, 83, 71, 225, 102, 59, 147, 224, 101, -944, 224, 224, 4, 224, 1002, 223, 8, 223, 101, 3, 224, 224, 1, 224, 223, 223, 2, 40, 139, 224, 1001, 224, -3905, 224, 4, 224, 1002, 223, 8, 223, 101, 7, 224, 224, 1, 223, 224, 223, 1101, 6, 94, 224, 101, -100, 224, 224, 4, 224, 1002, 223, 8, 223, 101, 6, 224, 224, 1, 224, 223, 223, 1102, 75, 30, 225, 1102, 70, 44, 224, 101, -3080, 224, 224, 4, 224, 1002, 223, 8, 223, 1001, 224, 4, 224, 1, 223, 224, 223, 1101, 55, 20, 225, 1102, 55, 16, 225, 1102, 13, 94, 225, 1102, 16, 55, 225, 1102, 13, 13, 225, 1, 109, 143, 224, 101, -88, 224, 224, 4, 224, 1002, 223, 8, 223, 1001, 224, 2, 224, 1, 223, 224, 223, 1002, 136, 57, 224, 101, -1140, 224, 224, 4, 224, 1002, 223, 8, 223, 101, 6, 224, 224, 1, 223, 224, 223, 101, 76, 35, 224, 1001, 224, -138, 224, 4, 224, 1002, 223, 8, 223, 101, 5, 224, 224, 1, 223, 224, 223, 4, 223, 99, 0, 0, 0, 677, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1105, 0, 99999, 1105, 227, 247, 1105, 1, 99999, 1005, 227, 99999, 1005, 0, 256, 1105, 1, 99999, 1106, 227, 99999, 1106, 0, 265, 1105, 1, 99999, 1006, 0, 99999, 1006, 227, 274, 1105, 1, 99999, 1105, 1, 280, 1105, 1, 99999, 1, 225, 225, 225, 1101, 294, 0, 0, 105, 1, 0, 1105, 1, 99999, 1106, 0, 300, 1105, 1, 99999, 1, 225, 225, 225, 1101, 314, 0, 0, 106, 0, 0, 1105, 1, 99999, 1008, 677, 677, 224, 1002, 223, 2, 223, 1006, 224, 329, 1001, 223, 1, 223, 8, 677, 226, 224, 102, 2, 223, 223, 1006, 224, 344, 101, 1, 223, 223, 1107, 226, 226, 224, 1002, 223, 2, 223, 1006, 224, 359, 1001, 223, 1, 223, 1108, 677, 226, 224, 102, 2, 223, 223, 1005, 224, 374, 1001, 223, 1, 223, 1007, 226, 226, 224, 102, 2, 223, 223, 1006, 224, 389, 1001, 223, 1, 223, 108, 677, 677, 224, 1002, 223, 2, 223, 1005, 224, 404, 1001, 223, 1, 223, 1007, 677, 677, 224, 102, 2, 223, 223, 1005, 224, 419, 1001, 223, 1, 223, 8, 226, 677, 224, 102, 2, 223, 223, 1005, 224, 434, 101, 1, 223, 223, 1008, 677, 226, 224, 102, 2, 223, 223, 1006, 224, 449, 1001, 223, 1, 223, 7, 677, 677, 224, 102, 2, 223, 223, 1006, 224, 464, 1001, 223, 1, 223, 8, 226, 226, 224, 1002, 223, 2, 223, 1005, 224, 479, 1001, 223, 1, 223, 7, 226, 677, 224, 102, 2, 223, 223, 1006, 224, 494, 1001, 223, 1, 223, 7, 677, 226, 224, 1002, 223, 2, 223, 1005, 224, 509, 101, 1, 223, 223, 107, 677, 677, 224, 102, 2, 223, 223, 1006, 224, 524, 101, 1, 223, 223, 1007, 677, 226, 224, 102, 2, 223, 223, 1006, 224, 539, 101, 1, 223, 223, 107, 226, 226, 224, 1002, 223, 2, 223, 1006, 224, 554, 101, 1, 223, 223, 1008, 226, 226, 224, 102, 2, 223, 223, 1006, 224, 569, 1001, 223, 1, 223, 1107, 677, 226, 224, 1002, 223, 2, 223, 1005, 224, 584, 101, 1, 223, 223, 1107, 226, 677, 224, 102, 2, 223, 223, 1005, 224, 599, 101, 1, 223, 223, 1108, 226, 677, 224, 102, 2, 223, 223, 1005, 224, 614, 101, 1, 223, 223, 108, 677, 226, 224, 102, 2, 223, 223, 1005, 224, 629, 101, 1, 223, 223, 107, 226, 677, 224, 102, 2, 223, 223, 1006, 224, 644, 1001, 223, 1, 223, 1108, 226, 226, 224, 1002, 223, 2, 223, 1006, 224, 659, 101, 1, 223, 223, 108, 226, 226, 224, 102, 2, 223, 223, 1005, 224, 674, 101, 1, 223, 223, 4, 223, 99, 226}
	loop()
}
