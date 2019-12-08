package intcode

import (
	"fmt"
)

func (state *IntcodeState) systemState() {
	fmt.Printf("Final state: ")
	fmt.Println(state.machineState)
}

func (state *IntcodeState) addOp(inputOne int, inputTwo int) {
	resReg := state.machineState[state.i+3]

	state.machineState[resReg] = inputOne + inputTwo
}

func (state *IntcodeState) multOp(inputOne int, inputTwo int) {
	resReg := state.machineState[state.i+3]

	state.machineState[resReg] = inputOne * inputTwo
}

func (state *IntcodeState) inputOp(location int) {
	input := state.input()
	fmt.Printf("Machine wanted input gave %d\n", input)
	state.machineState[location] = input
}

func (state *IntcodeState) outputOp(input int) {
	state.output(input)
}

func (state *IntcodeState) jumpIfTrueOp(inputOne int, inputTwo int) bool {
	if inputOne != 0 {
		state.i = inputTwo
	}

	return inputOne != 0
}

func (state *IntcodeState) jumpIfFalseOp(inputOne int, inputTwo int) bool {
	if inputOne == 0 {
		state.i = inputTwo
	}

	return inputOne == 0
}

func (state *IntcodeState) lessThanOp(inputOne int, inputTwo int) {
	resReg := state.machineState[state.i+3]
	if inputOne < inputTwo {
		state.machineState[resReg] = 1
	} else {
		state.machineState[resReg] = 0
	}
}

func (state *IntcodeState) equalsOp(inputOne int, inputTwo int) {
	resReg := state.machineState[state.i+3]
	if inputOne == inputTwo {
		state.machineState[resReg] = 1
	} else {
		state.machineState[resReg] = 0
	}
}

func (state *IntcodeState) haltOp() {
	state.systemState()
	fmt.Println("Program finished :)")
}

func (state *IntcodeState) validate(position int) (int, [3]int) {
	instruction := state.machineState[position]
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
			arr[curr] = state.machineState[curr+position+1]
		} else {
			arr[curr] = state.machineState[state.machineState[curr+position+1]]
		}
	}

	return opcode, arr
}

func (state *IntcodeState) Loop() {
	state.i = 0
	for {
		opCode, inputs := state.validate(state.i)

		if opCode == 1 {
			state.addOp(inputs[0], inputs[1])
			state.i += 4
		} else if opCode == 2 {
			state.multOp(inputs[0], inputs[1])
			state.i += 4
		} else if opCode == 3 {
			state.inputOp(state.machineState[state.i+1])
			state.i += 2
		} else if opCode == 4 {
			state.outputOp(inputs[0])
			state.i += 2
		} else if opCode == 5 {
			didJump := state.jumpIfTrueOp(inputs[0], inputs[1])
			if !didJump {
				state.i += 3
			}
		} else if opCode == 6 {
			didJump := state.jumpIfFalseOp(inputs[0], inputs[1])
			if !didJump {
				state.i += 3
			}
		} else if opCode == 7 {
			state.lessThanOp(inputs[0], inputs[1])
			state.i += 4
		} else if opCode == 8 {
			state.equalsOp(inputs[0], inputs[1])
			state.i += 4
		} else if opCode == 99 {
			state.haltOp()
			break
		} else {
			fmt.Printf("Unknown op: %d\n", opCode)
			panic("Machine broken!")
		}
	}
}

type IntcodeState struct {
	i            int
	machineState []int
	output func(int) 
	input func() int
}

func NewProgram(initialState []int, output func(int), input func() int) *IntcodeState {
	intcodeState := new(IntcodeState)
	intcodeState.i = 0
	intcodeState.output = output
	intcodeState.input = input
	intcodeState.machineState = make([]int, len(initialState))
	copy(intcodeState.machineState, initialState)

	return intcodeState
}

