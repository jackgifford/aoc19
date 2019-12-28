package main

import (
	"fmt"
	"math"
	"os"
)

var currBoard *([5][]byte)
var state map[string]bool

type tile struct {
	level int
	num   int
}

func newTile(num int, level int) tile {
	return tile{
		level: level,
		num:   num,
	}
}

func getNum(x, y int) int {
	return (x * 5) + y + 1
}

func getTile(num, level int) byte {
	num = num - 1
	x := num / 5
	y := (num - 5*x)

	if _, ok := all[gen-1][level]; !ok {
		all[gen-1][level] = genNew()
	}

	return all[gen-1][level][x][y]
}

func getNeg(num int, level int) []tile {
	switch num {
	case 1:
		return []tile{
			newTile(2, level),
			newTile(6, level),
			newTile(12, level+1),
			newTile(8, level+1),
		}
	case 2:
		return []tile{
			newTile(1, level),
			newTile(7, level),
			newTile(3, level),
			newTile(8, level+1),
		}
	case 3:
		return []tile{
			newTile(2, level),
			newTile(8, level),
			newTile(4, level),
			newTile(8, level+1),
		}
	case 4:
		return []tile{
			newTile(3, level),
			newTile(9, level),
			newTile(5, level),
			newTile(8, level+1),
		}
	case 5:
		return []tile{
			newTile(4, level),
			newTile(10, level),
			newTile(14, level+1),
			newTile(8, level+1),
		}
	case 6:
		return []tile{
			newTile(1, level),
			newTile(7, level),
			newTile(11, level),
			newTile(12, level+1),
		}
	case 7:
		return []tile{
			newTile(2, level),
			newTile(8, level),
			newTile(12, level),
			newTile(6, level),
		}
	case 8:
		return []tile{
			newTile(3, level),
			newTile(7, level),
			newTile(9, level),
			newTile(1, level-1),
			newTile(2, level-1),
			newTile(3, level-1),
			newTile(4, level-1),
			newTile(5, level-1),
		}
	case 9:
		return []tile{
			newTile(4, level),
			newTile(10, level),
			newTile(14, level),
			newTile(8, level),
		}
	case 10:
		return []tile{
			newTile(5, level),
			newTile(9, level),
			newTile(15, level),
			newTile(14, level+1),
		}
	case 11:
		return []tile{
			newTile(6, level),
			newTile(12, level),
			newTile(16, level),
			newTile(12, level+1),
		}
	case 12:
		return []tile{
			newTile(7, level),
			newTile(11, level),
			newTile(17, level),
			newTile(1, level-1),
			newTile(6, level-1),
			newTile(11, level-1),
			newTile(16, level-1),
			newTile(21, level-1),
		}
	case 13:
		panic("Shouldn't be called")
	case 14:
		return []tile{
			newTile(9, level),
			newTile(15, level),
			newTile(19, level),
			newTile(5, level-1),
			newTile(10, level-1),
			newTile(15, level-1),
			newTile(20, level-1),
			newTile(25, level-1),
		}
	case 15:
		return []tile{
			newTile(10, level),
			newTile(14, level),
			newTile(20, level),
			newTile(14, level+1),
		}
	case 16:
		return []tile{
			newTile(11, level),
			newTile(17, level),
			newTile(21, level),
			newTile(12, level+1),
		}
	case 17:
		return []tile{
			newTile(12, level),
			newTile(18, level),
			newTile(22, level),
			newTile(16, level),
		}
	case 18:
		return []tile{
			newTile(19, level),
			newTile(23, level),
			newTile(17, level),
			newTile(21, level-1),
			newTile(22, level-1),
			newTile(23, level-1),
			newTile(24, level-1),
			newTile(25, level-1),
		}
	case 19:
		return []tile{
			newTile(14, level),
			newTile(20, level),
			newTile(24, level),
			newTile(18, level),
		}
	case 20:
		return []tile{
			newTile(15, level),
			newTile(19, level),
			newTile(25, level),
			newTile(14, level+1),
		}
	case 21:
		return []tile{
			newTile(16, level),
			newTile(22, level),
			newTile(18, level+1),
			newTile(12, level+1),
		}
	case 22:
		return []tile{
			newTile(21, level),
			newTile(17, level),
			newTile(23, level),
			newTile(18, level+1),
		}
	case 23:
		return []tile{
			newTile(22, level),
			newTile(18, level),
			newTile(24, level),
			newTile(18, level+1),
		}
	case 24:
		return []tile{
			newTile(23, level),
			newTile(19, level),
			newTile(25, level),
			newTile(18, level+1),
		}
	case 25:
		return []tile{
			newTile(24, level),
			newTile(20, level),
			newTile(14, level+1),
			newTile(18, level+1),
		}
	}
	panic("Shouldn't be called")
}

func calcBio(field []byte) int {
	currTotal := 0
	for i := 0; i < 25; i++ {
		if string(field[i]) == "#" {
			currTotal += int(math.Pow(2, float64(i)))
		}
	}

	return currTotal
}

func prettyPrint(gen, level int) {
	for x := 0; x < 5; x++ {
		item := all[gen][level][x]
		fmt.Println(string(item))
		//fmt.Println(string(all[gen][level][x]))
	}
}

func getHash(x int) {
	base := make([]byte, 0)
	base = append(base, currBoard[0]...)
	base = append(base, currBoard[1]...)
	base = append(base, currBoard[2]...)
	base = append(base, currBoard[3]...)
	base = append(base, currBoard[4]...)

	if state[string(base)] {
		//fmt.Printf("%v\n", x)
		fmt.Printf("%v\n", calcBio(base))
		os.Exit(0)
	}
	state[string(base)] = true
	//fmt.Println(string(base))
}

func setup() {
	test := [5][]byte{}
	state = make(map[string]bool)

	test[0] = []byte("#####")
	test[1] = []byte("...##")
	test[2] = []byte("#..#.")
	test[3] = []byte("#....")
	test[4] = []byte("#...#")
	/*

		test[0] = []byte("....#")
		test[1] = []byte("#..#.")
		test[2] = []byte("#.?##")
		test[3] = []byte("..#..")
		test[4] = []byte("#....")

	*/
	all[0][0] = test

	//currBoard = &test
}

func shouldDie(negs []tile) bool {
	adjCount := 0
	for _, neg := range negs {
		if string(getTile(neg.num, neg.level)) == "#" {
			adjCount += 1
		}
	}

	return adjCount != 1

}

func shouldInfest(negs []tile) bool {
	adjCount := 0
	for _, neg := range negs {
		if string(getTile(neg.num, neg.level)) == "#" {
			adjCount += 1
		}
	}

	return adjCount == 1 || adjCount == 2
}

func nextGen(level int) {
	//x := 0
	newBase := genNew()
	for curr := 1; curr <= 25; curr++ {
		if curr == 13 {
			continue
		}

		neg := getNeg(curr, level)
		num := curr - 1
		x := num / 5
		y := (num - 5*x)

		if string(getTile(curr, level)) == "#" {
			if shouldDie(neg) {
				newBase[x][y] = byte('.')
			} else {
				newBase[x][y] = byte('#')
			}
		} else {
			if shouldInfest(neg) {
				newBase[x][y] = byte('#')
			} else {
				newBase[x][y] = byte('.')
			}

		}

		//fmt.Println(shouldDie(neg))
		//fmt.Println(neg)
	}

	all[gen][level] = newBase

	//fmt.Println(all)
}

var gen int
var all []map[int][5][]byte

func genNew() [5][]byte {
	test := [5][]byte{}
	test[0] = []byte(".....")
	test[1] = []byte(".....")
	test[2] = []byte("..?..")
	test[3] = []byte(".....")
	test[4] = []byte(".....")

	return test
}

func countTotal(level int) int {
	total := 0

	for x := 0; x < 5; x++ {
		for y := 0; y < 5; y++ {
			if string(all[200][level][x][y]) == "#" {
				total += 1
			}
		}
	}

	return total
}

func main() {
	gen = 0
	all = make([]map[int][5][]byte, 221)
	all[0] = make(map[int][5][]byte)
	setup()
	gen = 1

	x := 1
	for x < 221 {
		all[x] = make(map[int][5][]byte)

		curr := -400

		for curr < 400 {
			nextGen(curr)
			curr++
		}

		x++
		gen++
	}

	total := 0
	for k := -400; k < 400; k++ {
		total += countTotal(k)
	}
	fmt.Println(total)
}
