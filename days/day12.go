package main

import (
	"fmt"
	"bytes"
)

type moon struct {
	xPos int
	yPos int
	zPos int

	xVel int
	yVel int
	zVel int

	xDel int
	yDel int
	zDel int
}

var moons = [](*moon){}
var states = map[string]bool{}

// sweats in discrete
func gcd(a, b int) int {
	for b != 0 {
		t := b
		b = a % b
		a = t
	}
	return a
}

func lcm(a, b int, extra ...int) int {
	result := a * b / gcd(a,b)

	for i := 0; i < len(extra); i++ {
		result = lcm(result, extra[i])
	}

	return result
}

func abs(x int) int {
	if x < 0 {
		return x * -1
	}

	return x
}

func calcEnergy(x int, y int, z int) int {
	return abs(x) + abs(y) + abs(z)
}

func apply(curr *moon, next *moon) {
	if curr.xPos > next.xPos {
		curr.xDel += -1
	} else if curr.xPos < next.xPos {
		curr.xDel += 1
	}

	if curr.yPos > next.yPos {
		curr.yDel += -1
	} else if curr.yPos < next.yPos {
		curr.yDel += 1
	}

	if curr.zPos > next.zPos {
		curr.zDel += -1
	} else if curr.zPos < next.zPos {
		curr.zDel += 1
	}

}

func calcGrav() {
	apply(moons[0], moons[1])
	apply(moons[0], moons[2])
	apply(moons[0], moons[3])

	apply(moons[1], moons[0])
	apply(moons[1], moons[2])
	apply(moons[1], moons[3])

	apply(moons[2], moons[1])
	apply(moons[2], moons[0])
	apply(moons[2], moons[3])

	apply(moons[3], moons[1])
	apply(moons[3], moons[2])
	apply(moons[3], moons[0])

}

// merged grav + velo into one step since it saves a loop
func applyGrav() {
	for _, curr := range moons {
		curr.xVel += curr.xDel
		curr.yVel += curr.yDel
		curr.zVel += curr.zDel

		curr.xPos += curr.xVel
		curr.yPos += curr.yVel
		curr.zPos += curr.zVel

		// Might be overkill
		curr.xDel = 0
		curr.yDel = 0
		curr.zDel = 0
	}
}

func addMoon(xPos int, yPos int, zPos int) *moon {

	newMoon := new(moon)
	newMoon.xPos = xPos
	newMoon.yPos = yPos
	newMoon.zPos = zPos

	return newMoon
}

func partTwo() {
	phaseMap := make(map[string]bool)

	// handle x first.
	xPhase := 0

	for {

		var b bytes.Buffer

		for _, curr := range moons {
			fmt.Fprintf(&b, "pos %v vel %v", curr.xPos, curr.xVel)
		}

		// build the string
		if phaseMap[b.String()] {
			break
		}

		phaseMap[b.String()] = true

		calcGrav()
		applyGrav()

		xPhase += 1
	}

	fmt.Printf("x repeats after %v steps!\n", xPhase)

	phaseMap = make(map[string]bool)
	setGame()

	// handle x first.
	yPhase := 0

	for {

		var b bytes.Buffer

		for _, curr := range moons {
			fmt.Fprintf(&b, "ypos %v yvel %v", curr.yPos, curr.yVel)
		}

		// build the string
		if phaseMap[b.String()] {
			break
		}

		phaseMap[b.String()] = true

		calcGrav()
		applyGrav()

		yPhase += 1
	}
	fmt.Printf("y repeats after %v steps!\n", yPhase)

	phaseMap = make(map[string]bool)
	setGame()

	// handle x first.
	zPhase := 0

	for {

		var b bytes.Buffer

		for _, curr := range moons {
			fmt.Fprintf(&b, "zpos %v zvel %v", curr.zPos, curr.zVel)
		}

		// build the string
		if phaseMap[b.String()] {
			break
		}

		phaseMap[b.String()] = true

		calcGrav()
		applyGrav()

		zPhase += 1
	}
	fmt.Printf("z repeats after %v steps!\n", zPhase)

	fmt.Printf("first phase rep %v \n", lcm(xPhase, yPhase, zPhase))

}

// It's naieve but we can get away with it
func partOne() {
	for i := 0; i < 1000; i++ {
		calcGrav()
		applyGrav()
	}

	totalEnergy := 0

	for _, curr := range moons {
		totalEnergy += (calcEnergy(curr.xPos, curr.yPos, curr.zPos) * calcEnergy(curr.xVel, curr.yVel, curr.zVel))
	}

	fmt.Printf("Part one: %v\n", totalEnergy)
}

func setGame() {
	//moons = make([](*moon), 4)
	moons = nil
	/*
	moons = append(moons, addMoon(-1, 0, 2))
	moons = append(moons, addMoon(2, -10, -7))
	moons = append(moons, addMoon(4, -8, 8))
	moons = append(moons, addMoon(3, 5, -1))
	*/

	moons = append(moons, addMoon(3, 15, 8))
	moons = append(moons, addMoon(5, -1, -2))
	moons = append(moons, addMoon(-10, 8, 2))
	moons = append(moons, addMoon(8, 4, -5))
}

func main() {
	setGame()

	/*
	moons = append(moons, addMoon(3, 15, 8))
	moons = append(moons, addMoon(5, -1, -2))
	moons = append(moons, addMoon(-10, 8, 2))
	moons = append(moons, addMoon(8, 4, -5))
	*/

	//partOne()
	partTwo()
}
