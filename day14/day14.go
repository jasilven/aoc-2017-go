package main

import (
	"fmt"
	"jasilven/aoc17/hash"
	"log"
	"sort"
	"strconv"
)

const gSize = 128

func genGrid(input string) [gSize][gSize]int {
	var result [gSize][gSize]int
	for i := 0; i < gSize; i++ {
		h, row := hash.KnotHash(fmt.Sprintf("%s-%d", input, i)), ""
		for _, ch := range h {
			num, err := strconv.ParseInt(fmt.Sprintf("%c", ch), 16, 64)
			if err != nil {
				log.Fatal(err)
			}
			row = row + fmt.Sprintf("%0.4b", num)
		}
		if len(row) != gSize {
			log.Fatal("wrong row size in grid")
		}
		for j, ch := range row {
			if ch == '1' {
				result[i][j] = '#'
			} else {
				result[i][j] = '.'
			}
		}
	}
	return result
}

func solve1(grid *[gSize][gSize]int) (result int) {
	for row := 0; row < gSize; row++ {
		for col := 0; col < gSize; col++ {
			if grid[row][col] == '#' {
				result++
			}
		}
	}
	return
}

func initGridGroups(grid *[gSize][gSize]int) {
	squareCnt := 1
	for row := 0; row < gSize; row++ {
		for col := 0; col < gSize; col++ {
			if grid[row][col] == '#' {
				grid[row][col] = squareCnt
				squareCnt++
			} else {
				grid[row][col] = 0
			}
		}
	}
}

func getMinNeighbour(grid *[gSize][gSize]int, row, col int) int {
	neigbours := []int{}
	if row > 0 && grid[row-1][col] > 0 {
		neigbours = append(neigbours, grid[row-1][col])
	}
	if row < (gSize-1) && grid[row+1][col] > 0 {
		neigbours = append(neigbours, grid[row+1][col])
	}
	if col > 0 && grid[row][col-1] > 0 {
		neigbours = append(neigbours, grid[row][col-1])
	}
	if col < (gSize-1) && grid[row][col+1] > 0 {
		neigbours = append(neigbours, grid[row][col+1])
	}
	if len(neigbours) > 0 {
		sort.Ints(neigbours)
		return neigbours[0]
	}
	return 0
}

func solve2(grid *[gSize][gSize]int) int {
	initGridGroups(grid)
	cont := true
	for cont {
		cont = false
		for row := 0; row < gSize; row++ {
			for col := 0; col < gSize; col++ {
				min := getMinNeighbour(grid, row, col)
				if (min != 0) && (min < grid[row][col]) {
					grid[row][col] = min
					cont = true
				}
			}
		}
	}
	groups := make(map[int]bool)
	for row := 0; row < gSize; row++ {
		for col := 0; col < gSize; col++ {
			groups[grid[row][col]] = true
		}
	}
	delete(groups, 0)
	return len(groups)
}

func main() {
	grid := genGrid("uugsqrei")
	grid2 := grid
	part1, part2 := make(chan int, 1), make(chan int, 1)
	go func(g *[gSize][gSize]int) { part1 <- solve1(g) }(&grid)
	go func(g *[gSize][gSize]int) { part2 <- solve2(g) }(&grid2)
	fmt.Printf("part 1: %v\npart 2: %v\n", <-part1, <-part2)
}
