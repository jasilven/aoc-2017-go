package main

import (
	"fmt"
	"log"
	"math"
)

const xyfmt = "%d.%d"

func solve1(input int) int {
	level := 0
	for i, j := 0, 3; ; i, j = i+2, j+2 {
		if input > i*i && input <= j*j {
			level = j
			break
		}
	}
	count := input - ((level - 2) * (level - 2))
	return (level/2 + int(math.Abs(float64(count%(level-1)-(level/2)))))
}

type gwalker struct {
	grid      map[string]int
	x, y, dir int
}

func (gw *gwalker) move() {
	switch gw.dir {
	case 0:
		gw.y++
	case 1:
		gw.x--
	case 2:
		gw.y--
	case 3:
		gw.x++
	}
}

func (gw *gwalker) turnIfcorner() {
	x, y := gw.x, gw.y
	switch gw.dir {
	case 0:
		x--
	case 1:
		y--
	case 2:
		x++
	case 3:
		y++
	}
	_, ok := gw.grid[fmt.Sprintf(xyfmt, x, y)]
	if !ok {
		gw.dir = (gw.dir + 1) % 4
	}
}

func (gw *gwalker) getValue() int {
	result, ok := gw.grid[fmt.Sprintf(xyfmt, gw.x, gw.y)]
	if !ok {
		log.Fatalf("no value at point: %s", fmt.Sprint(xyfmt, gw.x, gw.y))
	}
	return result
}

func (gw *gwalker) setValue(value, x, y int) {
	gw.grid[fmt.Sprintf(xyfmt, x, y)] = value
}

func (gw *gwalker) updateGrid() {
	dx := []int{0, -1, -1, -1, 0, 1, 1, 1}
	dy := []int{1, 1, 0, -1, -1, -1, 0, 1}
	sum := 0
	for i := 0; i < len(dx); i++ {
		n, ok := gw.grid[fmt.Sprintf(xyfmt, gw.x+dx[i], gw.y+dy[i])]
		if ok {
			sum += n
		}
	}
	gw.setValue(sum, gw.x, gw.y)
}

func solve2(input int) (result int) {
	gw := gwalker{grid: make(map[string]int), x: 1, y: 0, dir: 3}
	gw.setValue(1, 0, 0)
	for {
		gw.updateGrid()
		result = gw.getValue()
		if result > input {
			break
		}
		gw.turnIfcorner()
		gw.move()
	}
	return
}

func main() {
	fmt.Println("Part 1:", solve1(289326))
	fmt.Println("Part 2:", solve2(289326))
}
