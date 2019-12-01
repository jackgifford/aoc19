package days

import (
	"os"
	"bufio"
	"fmt"
	"strconv"
)

func checker(e error) {
	if e != nil {
			panic(e)
	}
}

func operation(data int) int {
	return int(data / 3) - 2
}

func part_one() string {
	file, err := os.Open("./samples/day01.part01.txt")
	checker(err)
	defer file.Close()
	scanner := bufio.NewScanner(file)

	total_fuel := 0

	for scanner.Scan() {
		data := scanner.Text()
		as_int, _ := strconv.Atoi(data) 
		total_fuel += operation(as_int)
	}

	return strconv.Itoa(total_fuel)
}

func part_two() string {
	file, err := os.Open("./samples/day01.part01.txt")
	checker(err)
	defer file.Close()
	scanner := bufio.NewScanner(file)

	total_fuel := 0

	for scanner.Scan() {
		data := scanner.Text()
		as_int, _ := strconv.Atoi(data) 
		
		for as_int = operation(as_int); as_int >= 0; as_int = operation(as_int) {

			total_fuel += as_int
		}

	}

	return strconv.Itoa(total_fuel)
}

type DayOne struct {}

// Entry point
func (empty DayOne) Run() {
	fmt.Printf("AOC: Day 1\n")
	fmt.Printf("Part One: %s\n", part_one())
	fmt.Printf("Part Two: %s\n", part_two())
}


