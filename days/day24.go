package main

import (
	"fmt"
	"math"
	"os"
)

var currBoard *([5][]byte)
var states map[string]bool

func calcBio(field []byte) int {
	currTotal := 0
	for i := 0; i < 25; i++ {
		if string(field[i]) == "#" {
			currTotal += int(math.Pow(2, float64(i)))
		}
	}

	return currTotal
}

func prettyPrint() {
	for x := 0; x < 5; x++ {
		fmt.Println(string(currBoard[x]))
	}
}

func getHash(x int) {
	base := make([]byte, 0)
	base = append(base, currBoard[0]...)
	base = append(base, currBoard[1]...)
	base = append(base, currBoard[2]...)
	base = append(base, currBoard[3]...)
	base = append(base, currBoard[4]...)

	if states[string(base)] {
		//fmt.Printf("%v\n", x)
		fmt.Printf("%v\n", calcBio(base))
		os.Exit(0)
	}
	states[string(base)] = true
	//fmt.Println(string(base))
}

func setup() {
	test := [5][]byte{}
	states = make(map[string]bool)

	test[0] = []byte("#####")
	test[1] = []byte("...##")
	test[2] = []byte("#..#.")
	test[3] = []byte("#....")
	test[4] = []byte("#...#")

	currBoard = &test
}

func shouldDie(x, y int) bool {
	adjCount := 0
	// above
	if x-1 >= 0 && string(currBoard[x-1][y]) == "#" {
		adjCount += 1
	}

	if x+1 < 5 && string(currBoard[x+1][y]) == "#" {
		adjCount += 1
	}

	if y-1 >= 0 && string(currBoard[x][y-1]) == "#" {
		adjCount += 1
	}

	if y+1 < 5 && string(currBoard[x][y+1]) == "#" {
		adjCount += 1
	}

	return adjCount != 1
}

func shouldInfest(x, y int) bool {
	adjCount := 0
	// above
	if x-1 >= 0 && string(currBoard[x-1][y]) == "#" {
		adjCount += 1
	}

	if x+1 < 5 && string(currBoard[x+1][y]) == "#" {
		adjCount += 1
	}

	if y-1 >= 0 && string(currBoard[x][y-1]) == "#" {
		adjCount += 1
	}

	if y+1 < 5 && string(currBoard[x][y+1]) == "#" {
		adjCount += 1
	}

	return adjCount == 1 || adjCount == 2
}

func nextGen() {
	test := [5][]byte{}

	for x := 0; x < 5; x++ {
		test[x] = make([]byte, 5)
		for y := 0; y < 5; y++ {
			if string(currBoard[x][y]) == "#" {
				if shouldDie(x, y) {
					test[x][y] = byte('.')
				} else {
					test[x][y] = byte('#')
				}
			} else {
				if shouldInfest(x, y) {
					test[x][y] = byte('#')
				} else {
					test[x][y] = byte('.')
				}
			}
		}
	}

	currBoard = &test
}

func main() {
	setup()
	x := 0
	for {
		x++
		nextGen()
		getHash(x)
	}
}

