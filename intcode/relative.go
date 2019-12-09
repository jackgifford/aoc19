package intcode

import (
	"fmt"
)

// Both of these functions could do with some refactoring
func (state *RelativeIntcodeState) validate(position int) (int, [3]int) {
	instruction := state.baseState.machineState[position]
	opcode := instruction % 100
	instruction /= 100

	var arr [3]int
	var argCount int

	if opcode == 1 || opcode == 2 || opcode == 5 || opcode == 6 || opcode == 7 || opcode == 8 {
		argCount = 2
	} else if opcode == 3 || opcode == 4 || opcode == 9 {
		argCount = 1
	} else {
		argCount = 0
	}

	for curr := 0; curr < argCount; curr++ {
		modeType := instruction % 10
		instruction /= 10
		if modeType == 1 {
			arr[curr] = state.baseState.machineState[curr+position+1]
		} else if modeType == 2 {
			// rel input
			arr[curr] = state.baseState.machineState[state.relBase + curr + position + 1]
		} else {
			arr[curr] = state.baseState.machineState[state.baseState.machineState[curr+position+1]]
		}
	}

	return opcode, arr
}

func (state *RelativeIntcodeState) relOp(seek int) {
	state.relBase += seek
}

func (state *RelativeIntcodeState) Loop() {
	state.baseState.i = 0
	for {
		opCode, inputs := state.validate(state.baseState.i)

		if opCode == 1 {
			state.baseState.addOp(inputs[0], inputs[1])
			state.baseState.i += 4
		} else if opCode == 2 {
			state.baseState.multOp(inputs[0], inputs[1])
			state.baseState.i += 4
		} else if opCode == 3 {
			state.baseState.inputOp(inputs[0])
			//state.baseState.inputOp(state.baseState.machineState[state.baseState.i+1])
			state.baseState.i += 2
		} else if opCode == 4 {
			state.baseState.outputOp(inputs[0])
			state.baseState.i += 2
		} else if opCode == 5 {
			didJump := state.baseState.jumpIfTrueOp(inputs[0], inputs[1])
			if !didJump {
				state.baseState.i += 3
			}
		} else if opCode == 6 {
			didJump := state.baseState.jumpIfFalseOp(inputs[0], inputs[1])
			if !didJump {
				state.baseState.i += 3
			}
		} else if opCode == 7 {
			state.baseState.lessThanOp(inputs[0], inputs[1])
			state.baseState.i += 4
		} else if opCode == 8 {
			state.baseState.equalsOp(inputs[0], inputs[1])
			state.baseState.i += 4
		} else if opCode == 9 {
			state.relOp(inputs[0])
			//state.relOp(state.baseState.machineState[state.baseState.i + 1])
			state.baseState.i += 2
		} else if opCode == 99 {
			state.baseState.haltOp()
			break
		} else {
			fmt.Printf("Unknown op: %d\n", opCode)
			panic("Machine broken!")
		}
	}
}

type RelativeIntcodeState struct {
	baseState IntcodeState
	relBase int
}

func NewRelProg(initialState []int, name int, output func(int, int), input func(int) int) *RelativeIntcodeState {
	intcodeState := NewProgram(initialState, name, output, input)
	relState := new (RelativeIntcodeState)
	relState.baseState = *intcodeState
	relState.relBase = 0

	return relState
}

