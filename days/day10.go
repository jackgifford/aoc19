package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"sort"
)

var grid [][]byte

func prettyPrint() {
	for _, v := range grid {
		fmt.Println(string(v))
	}
}

type astImp struct {
	x   int
	y   int
	ang float64
}

type asteroid struct {
	x int
	y int
}

var asteroids []asteroid

func findAsteroids() {
	for y, v := range grid {
		for x, k := range v {
			if rune(k) == '#' {
				newAst := asteroid{
					x: x,
					y: y,
				}
				asteroids = append(asteroids, newAst)
			}
		}
	}
}

func calcAngles(curr asteroid) map[float64]bool {
	maps := make(map[float64]bool, 0)
	for _, next := range asteroids {
		if curr.x == next.x && curr.y == next.y {
			continue
		}
		res := math.Atan2(float64(curr.x-next.x), float64(curr.y-next.y))
		maps[res] = true
	}
	return maps
}

func makeGrid() {
	f, err := os.Open("./day10.2.dat")
	if err != nil {
		panic("bad file")
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)

	for scanner.Scan() {
		grid = append(grid, scanner.Bytes())
	}
}

func getAngles(curr asteroid) []astImp {
	angles := make([]astImp, 0)
	for _, next := range asteroids {
		if curr.x == next.x && curr.y == next.y {
			continue
		}
		res := math.Atan2(float64(curr.x-next.x), float64(curr.y-next.y))
		//pi := 3
		deg := float64(res * (180.0 / math.Pi))
		if deg < 0 {
			deg = deg + 360.0
		}

		if deg == 0 {
			deg = 360
		}

		better := astImp{
			x:   next.x,
			y:   next.y,
			ang: float64(deg),
		}
		angles = append(angles, better)
	}
	return angles
}

// take this output sort numeric, and take the last elem
func partOne() {
	for _, ast := range asteroids {
		maps := calcAngles(ast)
		fmt.Printf("%v -> (%v,%v)\n", len(maps), ast.x, ast.y)
	}
}

func abs(x int) int {
	if x < 0 {
		return x * -1
	}

	return x
}

func manhattanDistance(x astImp, y astImp) float64 {
	return math.Sqrt(math.Pow(float64(x.x-y.x), 2) + math.Pow(float64(x.y-y.y), 2))
}

func partTwo() {
	station := asteroid{
		x: 20,
		y: 19,
	}

	stationImp := astImp{
		x:   20,
		y:   19,
		ang: 0,
	}

	maps := getAngles(station)



	groups := map[float64][]astImp{}

	for _, curr := range maps {
		groups[curr.ang] = append(groups[curr.ang], curr)
	}

	for _, list := range groups {
		sort.SliceStable(list, func(i, j int) bool {
			return manhattanDistance(stationImp, list[i]) < manhattanDistance(stationImp, list[j])
		})
	}

	i := 0
	keys := make([]float64, len(groups))

	for key := range groups {
		keys[i] = key
		i += 1
	}

	sort.SliceStable(keys, func(i, j int) bool {
		return keys[i] > keys[j]
	})

	i = 0
	for {

		for _, key := range keys {
			if len(groups[key]) == 0 {
				continue
			}
			curr := groups[key][0]
			fmt.Println(curr)
			groups[key] = groups[key][1:]
			if i == 199 {
				os.Exit(1)
			}
			i++
		}
	}

}

func main() {
	grid = make([][]byte, 0)
	makeGrid()
	findAsteroids()
	//partOne()
	partTwo()
}
