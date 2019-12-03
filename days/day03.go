package days

import (
		"fmt"
		"strings"
		"strconv"
)

var game_board = map[int]map[int]int { }

func mathAbs(x int) int {
		if x <= 0 {
			return x * -1
		}
		return x
}

func recordPath(x int, y int) {
		if game_board[x] == nil {
				game_board[x] = make(map[int]int)
		}

		game_board[x][y] = 1
}

func checkCollision(x int, y int) {
		if x == 0 && y == 0 {
				return
		}
		if game_board[x] != nil && game_board[x][y] == 1 {
				fmt.Printf("%d\n", mathAbs(x) + mathAbs(y))	
		}
}

func wire_two(wirePath []string, myFunc func(int, int)) {
		x := 0
		y := 0

		for _, pair := range wirePath {
				direction := rune(pair[0])
				steps, _ := strconv.Atoi(pair[1:])
				for i := 0; i < steps; i++ {

						myFunc(x, y)

						switch direction {
								case 'R': x += 1 
								case 'U': y += 1 
								case 'L': x -= 1 
								case 'D': y -= 1 
						}
				}
		}
}

func Run() {
		wire_two(strings.Split("R8,U5,L5,D3", ","), recordPath)
		wire_two(strings.Split("U7,R6,D4,L4", ","), checkCollision)
}

