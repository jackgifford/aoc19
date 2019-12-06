package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Planet struct {
	name    string
	planets [](*Planet)
	parent *Planet
}

var plans map[string]*Planet

func getPlanet(name string) *Planet {
	planet, ok := plans[name]
	if !ok {
		planet = new(Planet)
		planet.name = name

		plans[name] = planet
	}

	return planet
}

var visited map[string]bool

func setPlans() {
	for k, _ := range plans {
		visited[k] = false
	}
}

func dfsRecur(curr *Planet) {
	visited[curr.name] = true
	for _, w := range curr.planets {
		if !visited[w.name] {
			dfsRecur(w)
		}
	}
}

func bfs(start *Planet) {
	queue := []*Planet{ start }
	for len(queue)!= 0 {
		curr := queue[0]
		queue = queue[1:]

		if curr.name == "SAN" {
			fmt.Println("found santa")
			return
		}

		for _, child := range curr.planets {
			if !visited[child.name] {
				visited[child.name] = true
				child.parent = curr
				queue = append(queue, child)
			}
		}
	}

}

func visitCount() int {
	total := 0

	for _, v := range visited {
		if v {
			total += 1
		}
	}

	return total
}

func main() {
	plans = make(map[string]*Planet, 0)
	visited = make(map[string]bool, 0)

	f, err := os.Open("./day5.input")

	if err != nil {
		panic("Dead")
	}

	defer f.Close()

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		data := scanner.Text()

		splitter := strings.Split(data, ")")

		parent := getPlanet(splitter[0])
		child := getPlanet(splitter[1])

		parent.planets = append(parent.planets, child)
		child.planets = append(child.planets, parent)
	}

	/*
	tot := 0
	for v, k := range plans {
			fmt.Println(v)
		setPlans()
		dfsRecur(k)
		tot += visitCount() - 1 // You don't orbit yourself ;)
	}

	fmt.Println(tot)
	*/

	setPlans()
	bfs(plans["YOU"])

	curr := plans["SAN"]
	i := 0
	for curr.name != "YOU" {
		curr = curr.parent
		i++
	}

	fmt.Println(i)
}
