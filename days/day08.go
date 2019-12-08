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

func partOne() {
	length := 3
	height := 2
	pixels := length * height

	start := 0
	end := pixels
	data := "123456789012"

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

func partTwo() {
	length := 25
	height := 6

	pixels := length * height
	data = "0000"

	// need immutability
	baseLayer := []byte(data[0:pixels])

	start := pixels
	end := pixels * 2

	for end <= len(data) {
		layer := data[start:end]
		start += pixels
		end += pixels
		
		for i, v := range layer {
			if rune(baseLayer[i]) == '2' {
				baseLayer[i] = byte(v)
			}
		}
	}
	
	final := string(baseLayer)
	start = 0
	end = length

	for end <= len(final) {
			fmt.Println(final[start:end])
			start += length
			end += length
	}

}

func main() {
	//partOne()
	partTwo()
}
