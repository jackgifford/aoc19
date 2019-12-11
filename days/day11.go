package main

import (
	"aoc19/intcode"
	"fmt"
)

type robot struct {
	x     int
	y     int
	dir   int
	paint bool
}

var robotState = robot{
	x:     0,
	y:     0,
	dir:   0,
	paint: true,
}

func output(val int, machineName int) {
	paintTime := robotState.paint
	// reverse it
	robotState.paint = !robotState.paint
	if paintTime {
		fmt.Printf("Time to paint! %v %v\n", robotState.x, robotState.y)
		col := '.'
		if val == 1 {
			col = '#'
		}
		setColour(robotState.x, robotState.y, col)
	} else {

		if val == 1 {
			// turn right
			robotState.dir = ((robotState.dir + 90) % 360)
		} else {
			// turn left
			robotState.dir = ((robotState.dir - 90) % 360)
		}

		if robotState.dir < 0 {
			robotState.dir += 360
		}

		// move
		switch robotState.dir {
		case 0: // up
			robotState.y += 1
			break
		case 90: // right
			robotState.x += 1
			break
		case 180: // down
			robotState.y -= 1
			break
		case 270: //left
			robotState.x -= 1
			break
		default:
			panic("Invalid direction")
		}
		fmt.Println(robotState)

		fmt.Printf("Direction change!\n")
	}
}

func input(machineName int) int {
	currColour := getColour(robotState.x, robotState.y)

	// black
	if currColour == '.' {
		return 0
	}

	return 1
}

var grid = map[int]map[int]rune{}

func setColour(x int, y int, colour rune) {
	if grid[x] == nil {
		grid[x] = make(map[int]rune)
	}

	grid[x][y] = colour
}

func getColour(x int, y int) rune {
	if grid[x] == nil {
		grid[x] = make(map[int]rune)
		grid[x][y] = '.'
		return '.'
	}

	val, prs := grid[x][y]

	if !prs {
		grid[x][y] = '.'
		return '.'
	}

	return val
}

func outputDebug(outOne int, outTwo int) {
	output(outOne, 0)
	output(outTwo, 0)
}

func main() {

	robotCode := []int{3, 8, 1005, 8, 324, 1106, 0, 11, 0, 0, 0, 104, 1, 104, 0, 3, 8, 1002, 8, -1, 10, 1001, 10, 1, 10, 4, 10, 1008, 8, 1, 10, 4, 10, 1001, 8, 0, 29, 1, 1107, 14, 10, 1006, 0, 63, 1006, 0, 71, 3, 8, 1002, 8, -1, 10, 101, 1, 10, 10, 4, 10, 1008, 8, 1, 10, 4, 10, 1002, 8, 1, 61, 1, 103, 18, 10, 1006, 0, 14, 1, 105, 7, 10, 3, 8, 1002, 8, -1, 10, 101, 1, 10, 10, 4, 10, 1008, 8, 1, 10, 4, 10, 101, 0, 8, 94, 1006, 0, 37, 1006, 0, 55, 2, 1101, 15, 10, 3, 8, 1002, 8, -1, 10, 1001, 10, 1, 10, 4, 10, 1008, 8, 0, 10, 4, 10, 101, 0, 8, 126, 2, 1006, 12, 10, 3, 8, 102, -1, 8, 10, 101, 1, 10, 10, 4, 10, 1008, 8, 1, 10, 4, 10, 1001, 8, 0, 152, 3, 8, 102, -1, 8, 10, 1001, 10, 1, 10, 4, 10, 108, 0, 8, 10, 4, 10, 101, 0, 8, 173, 1006, 0, 51, 1006, 0, 26, 3, 8, 102, -1, 8, 10, 101, 1, 10, 10, 4, 10, 1008, 8, 0, 10, 4, 10, 1001, 8, 0, 202, 2, 8, 18, 10, 1, 103, 19, 10, 1, 1102, 1, 10, 1006, 0, 85, 3, 8, 102, -1, 8, 10, 1001, 10, 1, 10, 4, 10, 108, 0, 8, 10, 4, 10, 1001, 8, 0, 238, 2, 1002, 8, 10, 1006, 0, 41, 3, 8, 102, -1, 8, 10, 1001, 10, 1, 10, 4, 10, 108, 0, 8, 10, 4, 10, 101, 0, 8, 267, 2, 1108, 17, 10, 2, 105, 11, 10, 1006, 0, 59, 1006, 0, 90, 3, 8, 1002, 8, -1, 10, 1001, 10, 1, 10, 4, 10, 1008, 8, 1, 10, 4, 10, 1001, 8, 0, 304, 101, 1, 9, 9, 1007, 9, 993, 10, 1005, 10, 15, 99, 109, 646, 104, 0, 104, 1, 21102, 936735777688, 1, 1, 21101, 341, 0, 0, 1105, 1, 445, 21101, 0, 937264173716, 1, 21101, 352, 0, 0, 1106, 0, 445, 3, 10, 104, 0, 104, 1, 3, 10, 104, 0, 104, 0, 3, 10, 104, 0, 104, 1, 3, 10, 104, 0, 104, 1, 3, 10, 104, 0, 104, 0, 3, 10, 104, 0, 104, 1, 21101, 3245513819, 0, 1, 21102, 1, 399, 0, 1105, 1, 445, 21102, 1, 29086470235, 1, 21102, 410, 1, 0, 1105, 1, 445, 3, 10, 104, 0, 104, 0, 3, 10, 104, 0, 104, 0, 21101, 825544712960, 0, 1, 21102, 1, 433, 0, 1106, 0, 445, 21102, 825460826472, 1, 1, 21101, 0, 444, 0, 1106, 0, 445, 99, 109, 2, 22102, 1, -1, 1, 21101, 0, 40, 2, 21101, 0, 476, 3, 21102, 466, 1, 0, 1105, 1, 509, 109, -2, 2105, 1, 0, 0, 1, 0, 0, 1, 109, 2, 3, 10, 204, -1, 1001, 471, 472, 487, 4, 0, 1001, 471, 1, 471, 108, 4, 471, 10, 1006, 10, 503, 1101, 0, 0, 471, 109, -2, 2106, 0, 0, 0, 109, 4, 2101, 0, -1, 508, 1207, -3, 0, 10, 1006, 10, 526, 21101, 0, 0, -3, 21202, -3, 1, 1, 21201, -2, 0, 2, 21101, 0, 1, 3, 21101, 0, 545, 0, 1105, 1, 550, 109, -4, 2105, 1, 0, 109, 5, 1207, -3, 1, 10, 1006, 10, 573, 2207, -4, -2, 10, 1006, 10, 573, 21202, -4, 1, -4, 1106, 0, 641, 21202, -4, 1, 1, 21201, -3, -1, 2, 21202, -2, 2, 3, 21101, 0, 592, 0, 1105, 1, 550, 22101, 0, 1, -4, 21101, 1, 0, -1, 2207, -4, -2, 10, 1006, 10, 611, 21102, 1, 0, -1, 22202, -2, -1, -2, 2107, 0, -3, 10, 1006, 10, 633, 22101, 0, -1, 1, 21102, 633, 1, 0, 105, 1, 508, 21202, -2, -1, -2, 22201, -4, -2, -4, 109, -5, 2105, 1, 0}

	emptyList := make([]int, 1000)
	robotCode = append(robotCode, emptyList...)
	setColour(0, 0, '#')
	robotProg := intcode.NewProgram(robotCode, 0, output, input)
	robotProg.Loop()

	/*
		outputDebug(1,0)
		outputDebug(0,0)
		outputDebug(1,0)
		outputDebug(1,0)
		outputDebug(0,1)
		outputDebug(1,0)
		outputDebug(1,0)
	*/

	fmt.Println(grid)
	totalCells := 0
	for _, v := range grid {
		totalCells += len(v)
	}

	for x := range grid {
		for k := range grid[x] {
			fmt.Println(k)
		}
	}

	for y := 0; y > -6; y-- {
		for x := 0; x < 43; x++ {

			fmt.Printf("%v", string(getColour(x, y)))
		}
		fmt.Printf("\n")
	}

	// generate a bound rect
	// transfer it across and print out

	fmt.Println(totalCells)

}
