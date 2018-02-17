package main

import (
	"fmt"
	"jasilven/aoc17/util"
	"log"
	"strings"
)

func getMaxKey(input map[int]int) (result int) {
	for k, _ := range input {
		if k > result {
			result = k
		}
	}
	return result
}

func getScanners(fname string) map[int]int {
	result := make(map[int]int)
	for _, line := range util.ReadLines(fname) {
		var layer, srange int
		line = strings.Replace(line, ":", "", 1)
		_, err := fmt.Sscanf(line, "%d%d", &layer, &srange)
		if err != nil {
			log.Fatal(err)
		}
		result[layer] = srange
	}
	return result
}

func solve1(scanners map[int]int) (result int) {
	for pico := 1; pico <= getMaxKey(scanners); pico++ {
		srange, exists := scanners[pico]
		if exists && ((pico % (2*srange - 2)) == 0) {
			result += pico * srange
		}
	}
	return
}

func solve2(scanners map[int]int) (result int) {
OUTER:
	for delay := 0; result == 0; delay++ {
		for pico := 0; pico <= getMaxKey(scanners); pico++ {
			srange, exists := scanners[pico]
			if exists && (((pico + delay) % (2*srange - 2)) == 0) {
				continue OUTER
			}
		}
		result = delay
	}
	return result
}

func main() {
	scanners := getScanners("day13_input.txt")
	part1, part2 := make(chan int, 1), make(chan int, 1)
	go func(sc map[int]int) { part1 <- solve1(sc) }(scanners)
	go func(sc map[int]int) { part2 <- solve2(sc) }(scanners)
	fmt.Printf("part 1: %v\npart 2: %v\n", <-part1, <-part2)
}
