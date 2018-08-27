package main

import (
	"fmt"
	"jasilven/aoc-2017-go/util"
	"log"
	"strconv"
)

const input_file = "day5_input.txt"

func parseInstructions(fname string) (result []int) {
	for _, val := range util.ReadLines(fname) {
		num, err := strconv.Atoi(val)
		if err != nil {
			log.Fatal(err)
		}
		result = append(result, num)
	}
	return
}

func solve1(fname string) (steps int) {
	jumps, index := parseInstructions(fname), 0
	for index >= 0 && index < (len(jumps)) {
		index, jumps[index] = jumps[index]+index, jumps[index]+1
		steps++
	}
	return
}

func solve2(fname string) (steps int) {
	jumps, index := parseInstructions(fname), 0
	for index >= 0 && index < (len(jumps)) {
		offset := 1
		if jumps[index] >= 3 {
			offset = -1
		}
		index, jumps[index] = jumps[index]+index, jumps[index]+offset
		steps++
	}
	return
}

func main() {
	part1, part2 := make(chan int, 1), make(chan int, 1)
	go func(f string) { part1 <- solve1(f) }(input_file)
	go func(f string) { part2 <- solve2(f) }(input_file)
	fmt.Printf("part 1: %v\npart 2: %v\n", <-part1, <-part2)
}
