package main

import (
	"fmt"
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

var moons = [](*moon) {}

func abs(x int) int {
	if x < 0 {
		return x * -1
	}

	return x
}

func calcEnergy(x int, y int, z int) int{
	return abs(x) + abs(y) + abs(z)
}

func calcGrav() {
	for _, curr := range moons {
		for _, next := range moons {
			if curr == next {
					continue
			}

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
	}
}

func applyGrav() {
	for _, curr := range moons {
		curr.xVel += curr.xDel
		curr.yVel += curr.yDel
		curr.zVel += curr.zDel
		
		// Might be overkill
		curr.xDel = 0
		curr.yDel = 0
		curr.zDel = 0
	}
}

func applyVelo() {
	for _, curr := range moons {
		curr.xPos += curr.xVel
		curr.yPos += curr.yVel
		curr.zPos += curr.zVel
	}
}

func addMoon(xPos int, yPos int, zPos int) *moon {

	newMoon := new(moon)
	newMoon.xPos = xPos
	newMoon.yPos = yPos
	newMoon.zPos = zPos

	return newMoon
}

func main() {
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

	for i := 0; i < 1000; i++ {

	calcGrav()
	applyGrav()
	applyVelo()
	}
	totalEnergy := 0

	for _, curr := range moons {
		totalEnergy += (calcEnergy(curr.xPos, curr.yPos, curr.zPos) * calcEnergy(curr.xVel, curr.yVel, curr.zVel))
	}

	fmt.Println(totalEnergy)
	for _, curr := range moons {
		fmt.Printf("pos=<x= %v, y= %v, z=%v> vel=<x=%v, y= %v, z= %v>\n", curr.xPos, curr.yPos, curr.zPos, curr.xVel, curr.yVel, curr.zVel)
	}
}
