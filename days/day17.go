package main

import (
	"fmt"
)
	var tiles []string = []string{
		"..#..........",
		"..#..........",
		"#######...###",
		"#.#...#...#.#",
		"#############",
		"..#...#...#..",
		"..#####...^..",
	}

func intCheck(x, y int) bool {
	if x >= 0 && y + 1 >= 0 && x < len(tiles) && y + 1 < len(tiles[0]) && string(tiles[x][y+1]) != "#" {
		return false
	}
	if x >= 0 && y-1 >= 0 && x < len(tiles) && y-1 < len(tiles[0]) && string(tiles[x][y-1]) != "#" {
		return false
	}
	if x + 1 >= 0 && y >= 0 && x+1 < len(tiles) && y < len(tiles[0]) && string(tiles[x+1][y]) != "#" {
		return false
	}
	if x-1 >= 0 && y >= 0 && x-1 < len(tiles) && y < len(tiles[0]) && string(tiles[x-1][y]) != "#" {
		return false
	}

	return true
}

func main() {
	total := 0
	for i, x := range tiles {
		for j, y := range x {
			if string(y) == "#" {
				if intCheck(i, j) {
					total += (i * j)
				}
			}
		}
	}

	fmt.Println(total)
}
