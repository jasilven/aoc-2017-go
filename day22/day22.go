package main

import (
	"fmt"
	"jasilven/aoc-2017-go/util"
	"log"
)

const (
	Clean    int = 0
	Weakened int = 1
	Infected int = 2
	Flagged  int = 3
)

type virus struct {
	x, y, dir int
	nodes     map[string]int
}

func (v *virus) execute() (result bool) {
	key := fmt.Sprintf("%d.%d", v.x, v.y)
	switch v.nodes[key] {
	case Infected:
		v.dir = (v.dir + 1) % 4
		v.nodes[key] = Clean
	case Clean:
		v.dir = (v.dir + 3) % 4
		v.nodes[key] = Infected
		result = true
	default:
		log.Fatal("unknown nodes state:", v.nodes[key])
	}
	return
}

func (v *virus) execute2() (result bool) {
	key := fmt.Sprintf("%d.%d", v.x, v.y)
	switch v.nodes[key] {
	case Flagged:
		v.dir = (v.dir + 2) % 4
		v.nodes[key] = Clean
	case Weakened:
		v.nodes[key] = Infected
		result = true
	case Infected:
		v.dir = (v.dir + 1) % 4
		v.nodes[key] = Flagged
	case Clean:
		v.dir = (v.dir + 3) % 4
		v.nodes[key] = Weakened
	default:
		log.Fatal("unknown nodes state:", v.nodes[key])
	}
	return
}

func (v *virus) move() {
	switch v.dir {
	case 0:
		v.x--
	case 1:
		v.y--
	case 2:
		v.x++
	case 3:
		v.y++
	}
}

func loadNodes(fname string) (map[string]int, int, int) {
	result := make(map[string]int)
	lines := util.ReadLines(fname)
	for i := 0; i < len(lines); i++ {
		line := []byte(lines[i])
		for j := 0; j < len(line); j++ {
			key := fmt.Sprintf("%d.%d", -i, j)
			if line[j] == '#' {
				result[key] = Infected
			} else {
				result[key] = Clean
			}
		}
	}
	return result, len(lines), len(lines[0])
}

func solve1(nodes map[string]int, x, y int) (result int) {
	v := virus{x: -x / 2, y: y / 2, nodes: nodes, dir: 2}
	for i := 0; i < 10000; i++ {
		if v.execute() {
			result++
		}
		v.move()
	}
	return
}

func solve2(nodes map[string]int, x, y int) (result int) {
	v := virus{x: -x / 2, y: y / 2, nodes: nodes, dir: 2}
	for i := 0; i < 10000000; i++ {
		if v.execute2() {
			result++
		}
		v.move()
	}
	return
}

func main() {
	nodes, x, y := loadNodes("day22_input.txt")
	nodes2 := make(map[string]int)
	for k, v := range nodes {
		nodes2[k] = v
	}
	part1 := solve1(nodes, x, y)
	part2 := solve2(nodes2, x, y)
	fmt.Printf("Part 1: %v\nPart 2: %v\n", part1, part2)
}
