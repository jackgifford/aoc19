package main

import (
	"fmt"
		)

func getZeros(layer string, val rune) int {
	count := 0
	for _, v := range layer {
		if v == val {
			count++
		}
	}

	return count
}

func main() {
	length := 25
	height := 6
	pixels := length * height

	start := 0
	end := pixels

	var maxLayer string
	maxCount := 1 <<31 -1


	for end <= len(data) {
		layer := data[start:end]
		start += pixels
		end += pixels

		zeros := getZeros(layer, '0')
		if zeros < maxCount {
			maxCount = zeros
			maxLayer = layer
		}

	}

	fmt.Printf("Least zeroes %v\n", maxLayer)
	fmt.Printf("%v\n", getZeros(maxLayer, '1') * getZeros(maxLayer, '2'))
}
