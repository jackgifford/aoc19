package main

import (
	"fmt"
	"strings"
	"strconv"
)

var reactions map[string]data
var inventory map[string]int

type data struct {
	reqs []string
	makes int
}

func define(name string, makes int, ingred []string) {
	reactions[name] = data{reqs: ingred, makes: makes}
}


func setup() {
		reactions = make(map[string]data)
		inventory = make(map[string]int)
		/*
		reactions["A"] = data{ reqs: []string{"10 ORE"}, makes: 10 }
		reactions["B"] = data{ reqs: []string{"1 ORE"}, makes: 1 }
		reactions["C"] = data{ reqs: []string{"7 A", "1 B"}, makes: 1}
		reactions["D"] = data{reqs: []string{"7 A", "1 C"}, makes: 1}
		reactions["E"] = data{reqs: []string{"7 A", "1 D"}, makes: 1}
		reactions["FUEL"] = data{reqs: []string{"7 A", "1 E"}, makes: 1}

		define("A", 2, []string{"9 ORE"})
		define("B", 3, []string{"8 ORE"})
		define("C", 5, []string{"7 ORE"})
		define("AB", 1, []string{"3 A", "4 B"})
		define("BC", 1, []string{"5 B", "7 C"})
		define("CA", 1, []string{"4 C", "1 A"})
		define("FUEL", 1, []string{"2 AB", "3 BC", "4 CA"})
		*/

		define("NZVS", 5, []string{"157 ORE"})
		define("DCFZ", 6, []string{"165 ORE"})
		define("FUEL", 1, []string{"44 XJWVT", "5 KHKGT", "1 QDVJ", "29 NZVS", "9 GPVTF", "48 HKGWZ" })
		define("QDVJ", 9, []string{"12 HKGWZ", "1 GPVTF", "8 PSHF"})
		define("PSHF", 7, []string{"179 ORE"})
		define("HKGWZ", 5, []string{"177 ORE"})
		define("XJWVT", 2, []string{"7 DCFZ", "7 PSHF"})
		define("GPVTF", 2, []string{"165 ORE"})
		define("KHKGT", 8, []string{"3 DCFZ", "7 NZVS", "5 HKGWZ", "10 PSHF" })


}

var oreCount int = 0

func getReqs(input string, amount int) {
		out := reactions[input]

		min := out.makes

		if inventory[input] > amount {
			inventory[input] -= amount
			return
		}

		if input == "ORE" {
			//fmt.Println(amount)
			oreCount += amount
			return
		}

		for _, item := range out.reqs {
			split := strings.Split(item, " ")
			//fmt.Println(item)
			valReq, _ := strconv.Atoi(split[0])
			// before making check if we have it already... 
			getReqs(split[1], valReq) }

		if amount > min {
			getReqs(input, amount - min)
		}

		if amount < min {
			inventory[input] += min - amount
		}
}

func main() {
	setup()
	getReqs("FUEL", 1)
	fmt.Println(inventory)
	fmt.Println(oreCount)
	//fmt.Println(reactions["1 FUEL"])

}

