package intcode

import (
	"fmt"
)

func (state *IntcodeState) systemState() {
	fmt.Printf("Final state: ")
	//fmt.Println(state.machineState)
}

func (state *IntcodeState) addOp(inputOne int, inputTwo int, resReg int) {
	state.machineState[resReg] = state.machineState[inputOne] + state.machineState[inputTwo]
}

func (state *IntcodeState) multOp(inputOne int, inputTwo int, resReg int) {
	state.machineState[resReg] = state.machineState[inputOne] * state.machineState[inputTwo]
}

func (state *IntcodeState) inputOp(location int) {
	input := state.input(state.GetName())
	fmt.Printf("Machine wanted input gave %d\n", input)
	state.machineState[location] = input
}

func (state *IntcodeState) outputOp(input int) {
	state.output(state.machineState[input], state.GetName())
}

func (state *IntcodeState) jumpIfTrueOp(inputOne int, inputTwo int) bool {
	if state.machineState[inputOne] != 0 {
		state.i = state.machineState[inputTwo]
	}

	return state.machineState[inputOne] != 0
}

func (state *IntcodeState) jumpIfFalseOp(inputOne int, inputTwo int) bool {
	if state.machineState[inputOne] == 0 {
		state.i = state.machineState[inputTwo]
	}

	return state.machineState[inputOne] == 0
}

func (state *IntcodeState) lessThanOp(inputOne int, inputTwo int, resReg int) {
	//resReg := state.machineState[resReg]
	if state.machineState[inputOne] < state.machineState[inputTwo] {
		state.machineState[resReg] = 1
	} else {
		state.machineState[resReg] = 0
	}
}

func (state *IntcodeState) equalsOp(inputOne int, inputTwo int, resReg int) {
	//resReg := state.machineState[resReg]
	if state.machineState[inputOne] == state.machineState[inputTwo] {
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
		argCount = 3
	} else if opcode == 3 || opcode == 4 || opcode == 9 {
		argCount = 1
	} else {
		argCount = 0
	}

	for curr := 0; curr < argCount; curr++ {
		modeType := instruction % 10
		instruction /= 10
		if modeType == 1 {
			arr[curr] = curr+position+1
		} else if modeType == 2 {
			arr[curr] = state.machineState[curr+position+1] +  state.base	
		} else {
			arr[curr] = state.machineState[curr+position+1]
		}
	}

	return opcode, arr
}

func (state *IntcodeState) Loop() {
	state.i = 0
	for {
		opCode, inputs := state.validate(state.i)

		if opCode == 1 {
			state.addOp(inputs[0], inputs[1], inputs[2])
			state.i += 4
		} else if opCode == 2 {
			state.multOp(inputs[0], inputs[1], inputs[2])
			state.i += 4
		} else if opCode == 3 {
			state.inputOp(inputs[0])
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
			state.lessThanOp(inputs[0], inputs[1], inputs[2])
			state.i += 4
		} else if opCode == 8 {
			state.equalsOp(inputs[0], inputs[1], inputs[2])
			state.i += 4
		} else if opCode == 9 {
			state.base += state.machineState[inputs[0]]
			state.i += 2
		} else if opCode == 99 {
			state.haltOp()
			break
		} else {
			fmt.Printf("Unknown op: %d\n", opCode)
			panic("Machine broken!")
		}
	}
}

func (state *IntcodeState) GetName() int {
	return state.name
}

type IntcodeState struct {
	i            int
	machineState []int
	output func(int, int) 
	input func(int) int
	name int
	base int
}

func NewProgram(initialState []int, name int, output func(int, int), input func(int) int) *IntcodeState {
	intcodeState := new(IntcodeState)
	intcodeState.i = 0
	intcodeState.output = output
	intcodeState.input = input
	intcodeState.machineState = make([]int, len(initialState))
	intcodeState.name = name
	copy(intcodeState.machineState, initialState)
	intcodeState.base = 0

	return intcodeState
}

